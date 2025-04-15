import pulumi
import pulumi_gcp as gcp
import pulumi_random as random
import pulumi_docker as docker

# Create a random suffix for the repository name
random_suffix = random.RandomString("random-suffix",
    length=6,
    special=False,
    upper=False
)
repo_name = pulumi.Output.concat("docker-test-repo-", random_suffix.result)

# Create a private GCP artifact registry
registry = gcp.artifactregistry.Repository("my-registry",
    format="DOCKER",
    repository_id=repo_name,
    location="us-central1",  # change to your desired region
)

registry_url = pulumi.Output.concat(registry.location, "-docker.pkg.dev/", registry.project, "/", registry.repository_id)

# Get registry info (creds and endpoint)
image_name = pulumi.Output.concat(registry_url, "/myapp")

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
