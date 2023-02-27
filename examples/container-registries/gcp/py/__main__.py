import pulumi
import pulumi_gcp as gcp
import pulumi_docker as docker

# Create a private GCR repository.
registry = gcp.container.Registry('my-registry')
registry_url = registry.id.apply(lambda _: gcp.container.get_registry_repository().repository_url)

# Get registry info (creds and endpoint).
image_name = registry_url.apply(lambda url: f'{url}/myapp')
registry_info = None # use gcloud for auth.

# Build and publish the image.
image = docker.Image('my-image',
    build=DockerBuildArgs(
        context='app',
    ),
    image_name=image_name,
    registry=registry_info,
)

# Export the resulting base name in addition to the specific version pushed.
pulumi.export('baseImageName', image.base_image_name)
pulumi.export('imageName', image.image_name)
