package main

import (
	"encoding/base64"
	"encoding/json"

	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Fetch the Docker Hub auth info from config.
		cfg := config.New(ctx, "")
		username := cfg.Require("dockerUsername")
		password := cfg.RequireSecret("dockerPassword").(pulumi.StringOutput)

		// Build and publish the app image.
		image, err := docker.NewImage(ctx, "my-image", &docker.ImageArgs{
			Build:     &docker.DockerBuildArgs{Context: pulumi.String("app")},
			ImageName: pulumi.String(username + "/myapp"),
			Registry: docker.ImageRegistryArgs{
				Server:   pulumi.String("docker.io"),
				Username: pulumi.String(username),
				Password: password,
			},
		})
		if err != nil {
			return err
		}

		// Ensure we can pull from the Docker Hub.
		pullSecret, err := corev1.NewSecret(ctx, "my-regcred", &corev1.SecretArgs{
			Type: pulumi.String("kubernetes.io/dockerconfigjson"),
			StringData: pulumi.StringMap{
				".dockerconfigjson": password.ApplyString(func(pwd string) string {
					config := map[string]interface{}{
						"auths": map[string]interface{}{
							"https://index.docker.io/v1/": map[string]interface{}{
								"username": username,
								"password": pwd,
								"auth":     base64.StdEncoding.EncodeToString([]byte(username + ":" + pwd)),
							},
						},
					}
					ret, _ := json.Marshal(config)
					return string(ret)
				}),
			},
		})
		if err != nil {
			return err
		}

		// Deploy a load-balanced service that uses this image.
		labels := pulumi.StringMap{"app": pulumi.String("my-app")}
		_, err = appsv1.NewDeployment(ctx, "my-app-dep", &appsv1.DeploymentArgs{
			Spec: appsv1.DeploymentSpecArgs{
				Selector: metav1.LabelSelectorArgs{
					MatchLabels: labels,
				},
				Replicas: pulumi.Int(1),
				Template: corev1.PodTemplateSpecArgs{
					Metadata: metav1.ObjectMetaArgs{
						Labels: labels,
					},
					Spec: corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							corev1.ContainerArgs{
								Name:  labels["app"],
								Image: image.ImageName,
							},
						},
						ImagePullSecrets: corev1.LocalObjectReferenceArray{
							corev1.LocalObjectReferenceArgs{
								Name: pullSecret.Metadata.Name(),
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		svc, err := corev1.NewService(ctx, "my-app-svc", &corev1.ServiceArgs{
			Spec: corev1.ServiceSpecArgs{
				Selector: labels,
				Type:     pulumi.String("LoadBalancer"),
				Ports: corev1.ServicePortArray{
					corev1.ServicePortArgs{Port: pulumi.Int(80)},
				},
			},
		})
		if err != nil {
			return err
		}

		// Export the resulting image name and tag.
		ctx.Export("imageName", image.ImageName)
		// Export the Kubernetes ingress IP to access the service.
		ctx.Export("serviceIp", svc.Status.ApplyString(func(status *corev1.ServiceStatus) string {
			return *status.LoadBalancer.Ingress[0].Ip
		}))
		return nil
	})
}
