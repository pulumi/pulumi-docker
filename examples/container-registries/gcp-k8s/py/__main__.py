import pulumi
import pulumi_gcp as gcp
import pulumi_docker as docker
import pulumi_kubernetes as k8s

# Create a private GCR repository.
registry = gcp.container.Registry('my-registry')
registry_url = registry.id.apply(lambda _: gcp.container.get_registry_repository().repository_url)

# Get registry info (creds and endpoint).
image_name = registry_url.apply(lambda url: f'{url}/myapp')
registry_info = None # use gcloud for auth.

# Build and publish the image.
image = docker.Image('my-image',
    build='app',
    image_name=image_name,
    registry=registry_info,
)

# Export the resulting base name in addition to the specific version pushed.
pulumi.export('baseImageName', image.base_image_name)
pulumi.export('imageName', image.image_name)

# Create a load balanced Kubernetes service using this image, and export its IP.
app_labels = { 'app': 'myapp' }
app_dep = k8s.apps.v1.Deployment('app-dep',
    spec={
        'selector': { 'matchLabels': app_labels },
        'replicas': 3,
        'template': {
            'metadata': { 'labels': app_labels },
            'spec': {
                'containers': [{
                    'name': 'myapp',
                    'image': image.image_name,
                }],
            },
        },
    },
)
app_svc = k8s.core.v1.Service('app-svc',
    metadata={ 'labels': app_labels },
    spec={
        'type': 'LoadBalancer',
        'ports': [{ 'port': 80, 'targetPort': 80, 'protocol': 'TCP' }],
        'selector': app_labels,
    }
)
pulumi.export('appIp', app_svc.status.apply(lambda s: s.loadbalancer.ingress[0].ip))
