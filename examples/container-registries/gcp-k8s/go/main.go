package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
	"github.com/pulumi/pulumi-gcp/sdk/v2/go/gcp/container"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a private GCR registry.
		registry, err := container.NewRegistry(ctx, "my-registry", nil)
		if err != nil {
			return err
		}
		registryUrl := registry.ID().ApplyString(func(_ string) (string, error) {
			rep, err := container.GetRegistryRepository(ctx, nil)
			if err != nil {
				return "", err
			}
			return rep.RepositoryUrl, nil
		})

		// Get registry info (creds and endpoint).
		imageName := pulumi.Sprintf("%s/myapp", registryUrl)
		registryInfo := docker.ImageRegistryArgs{} // use gcloud

		// Build and publish the app image.
		image, err := docker.NewImage(ctx, "my-image", &docker.ImageArgs{
			Build:     &docker.DockerBuildArgs{Context: pulumi.String("app")},
			ImageName: imageName,
			Registry:  registryInfo,
		})

		// Export the resulting base name in addition to the specific version pushed.
		ctx.Export("baseImageName", image.BaseImageName)
		ctx.Export("fullImageName", image.ImageName)

		// Create a load balanced Kubernetes service using this image, and export its IP.
		appLabels := pulumi.StringMap{"app": pulumi.String("myapp")}
		_, deperr := appsv1.NewDeployment(ctx, "app-dep", &appsv1.DeploymentArgs{
			Metadata: &metav1.ObjectMetaArgs{Labels: appLabels},
			Spec: appsv1.DeploymentSpecArgs{
				Selector: &metav1.LabelSelectorArgs{MatchLabels: appLabels},
				Replicas: pulumi.Int(3),
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{Labels: appLabels},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							corev1.ContainerArgs{
								Name:  pulumi.String("myapp"),
								Image: image.ImageName,
							},
						},
					},
				},
			},
		})
		if deperr != nil {
			return deperr
		}
		appSvc, svcerr := corev1.NewService(ctx, "app-svc", &corev1.ServiceArgs{
			Metadata: &metav1.ObjectMetaArgs{Labels: appLabels},
			Spec: &corev1.ServiceSpecArgs{
				Type: pulumi.String("LoadBalancer"),
				Ports: corev1.ServicePortArray{
					corev1.ServicePortArgs{Port: pulumi.Int(80)},
				},
				Selector: appLabels,
			},
		})
		if svcerr != nil {
			return svcerr
		}
		ctx.Export("appIp", appSvc.Status.ApplyT(func(status *corev1.ServiceStatus) *string {
			return status.LoadBalancer.Ingress[0].Ip
		}))

		return nil
	})
}
