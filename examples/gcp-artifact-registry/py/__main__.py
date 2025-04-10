import pulumi
import pulumi_gcp as gcp
import pulumi_docker as docker

# Create a private GCR repository.
registry = gcp.container.Registry('my-registry')
registry_url = registry.id.apply(lambda _: gcp.container.get_registry_repository().repository_url)

# Get image name
image_name = registry_url.apply(lambda url: f'{url}/myapp')

# Build and publish the image.
image = docker.Image('my-image',
    build=docker.DockerBuildArgs(
        context='app',
    ),
    image_name=image_name,
)

# Export the resulting image name
pulumi.export('imageName', image.image_name)
pulumi.export('repoDigest', image.repo_digest)
