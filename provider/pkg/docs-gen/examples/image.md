{{% examples %}}
## Example Usage
{{% example %}}
### A Docker image build

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const demoImage = new docker.Image("demo-image", {
    build: {
        context: ".",
        dockerfile: "Dockerfile",
        platform: "linux/amd64",
    },
    imageName: "username/image:tag1",
    skipPush: true,
});
export const imageName = demoImage.imageName;
```
```python
import pulumi
import pulumi_docker as docker

demo_image = docker.Image("demo-image",
    build={
        "context": ".",
        "dockerfile": "Dockerfile",
        "platform": "linux/amd64",
    },
    image_name="username/image:tag1",
    skip_push=True)
pulumi.export("imageName", demo_image.image_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var demoImage = new Docker.Image("demo-image", new()
    {
        Build = new Docker.Inputs.DockerBuildArgs
        {
            Context = ".",
            Dockerfile = "Dockerfile",
            Platform = "linux/amd64",
        },
        ImageName = "username/image:tag1",
        SkipPush = true,
    });

    return new Dictionary<string, object?>
    {
        ["imageName"] = demoImage.ImageName,
    };
});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		demoImage, err := docker.NewImage(ctx, "demo-image", &docker.ImageArgs{
			Build: &docker.DockerBuildArgs{
				Context:    pulumi.String("."),
				Dockerfile: pulumi.String("Dockerfile"),
				Platform:   pulumi.String("linux/amd64"),
			},
			ImageName: pulumi.String("username/image:tag1"),
			SkipPush:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		ctx.Export("imageName", demoImage.ImageName)
		return nil
	})
}
```
```yaml
config: {}
description: A Docker image build
name: image-yaml
outputs:
    imageName: ${demo-image.imageName}
resources:
    demo-image:
        options:
            version: v4.4.0
        properties:
            build:
                context: .
                dockerfile: Dockerfile
                platform: linux/amd64
            imageName: username/image:tag1
            skipPush: true
        type: docker:Image
runtime: yaml
variables: {}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.Image;
import com.pulumi.docker.ImageArgs;
import com.pulumi.docker.inputs.DockerBuildArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var demoImage = new Image("demoImage", ImageArgs.builder()
            .build(DockerBuildArgs.builder()
                .context(".")
                .dockerfile("Dockerfile")
                .platform("linux/amd64")
                .build())
            .imageName("username/image:tag1")
            .skipPush(true)
            .build());

        ctx.export("imageName", demoImage.imageName());
    }
}
```
{{% /example %}}
{{% example %}}
### A Docker image build and push

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const demoPushImage = new docker.Image("demo-push-image", {
    build: {
        context: ".",
        dockerfile: "Dockerfile",
    },
    imageName: "docker.io/username/push-image:tag1",
});
export const imageName = demoPushImage.imageName;
export const repoDigest = demoPushImage.repoDigest;
```
```python
import pulumi
import pulumi_docker as docker

demo_push_image = docker.Image("demo-push-image",
    build={
        "context": ".",
        "dockerfile": "Dockerfile",
    },
    image_name="docker.io/username/push-image:tag1")
pulumi.export("imageName", demo_push_image.image_name)
pulumi.export("repoDigest", demo_push_image.repo_digest)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var demoPushImage = new Docker.Image("demo-push-image", new()
    {
        Build = new Docker.Inputs.DockerBuildArgs
        {
            Context = ".",
            Dockerfile = "Dockerfile",
        },
        ImageName = "docker.io/username/push-image:tag1",
    });

    return new Dictionary<string, object?>
    {
        ["imageName"] = demoPushImage.ImageName,
        ["repoDigest"] = demoPushImage.RepoDigest,
    };
});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		demoPushImage, err := docker.NewImage(ctx, "demo-push-image", &docker.ImageArgs{
			Build: &docker.DockerBuildArgs{
				Context:    pulumi.String("."),
				Dockerfile: pulumi.String("Dockerfile"),
			},
			ImageName: pulumi.String("docker.io/username/push-image:tag1"),
		})
		if err != nil {
			return err
		}
		ctx.Export("imageName", demoPushImage.ImageName)
		ctx.Export("repoDigest", demoPushImage.RepoDigest)
		return nil
	})
}
```
```yaml
config: {}
description: A Docker image build and push
name: image-push-yaml
outputs:
    imageName: ${demo-push-image.imageName}
    repoDigest: ${demo-push-image.repoDigest}
resources:
    demo-push-image:
        options:
            version: v4.4.0
        properties:
            build:
                context: .
                dockerfile: Dockerfile
            imageName: docker.io/username/push-image:tag1
        type: docker:Image
runtime: yaml
variables: {}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.Image;
import com.pulumi.docker.ImageArgs;
import com.pulumi.docker.inputs.DockerBuildArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var demoPushImage = new Image("demoPushImage", ImageArgs.builder()
            .build(DockerBuildArgs.builder()
                .context(".")
                .dockerfile("Dockerfile")
                .build())
            .imageName("docker.io/username/push-image:tag1")
            .build());

        ctx.export("imageName", demoPushImage.imageName());
        ctx.export("repoDigest", demoPushImage.repoDigest());
    }
}
```
{{% /example %}}
{{% example %}}
### Docker image build using caching with AWS Elastic Container Registry

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as docker from "@pulumi/docker";

const ecrRepository = new aws.ecr.Repository("ecr-repository", {name: "docker-repository"});
const authToken = aws.ecr.getAuthorizationTokenOutput({
    registryId: ecrRepository.registryId,
});
const myAppImage = new docker.Image("my-app-image", {
    build: {
        args: {
            BUILDKIT_INLINE_CACHE: "1",
        },
        cacheFrom: {
            images: [pulumi.interpolate`${ecrRepository.repositoryUrl}:latest`],
        },
        context: "app/",
        dockerfile: "app/Dockerfile",
    },
    imageName: pulumi.interpolate`${ecrRepository.repositoryUrl}:latest`,
    registry: {
        password: pulumi.secret(authToken.apply(authToken => authToken.password)),
        server: ecrRepository.repositoryUrl,
        username: authToken.apply(authToken => authToken.userName),
    },
});
export const imageName = myAppImage.imageName;
```
```python
import pulumi
import pulumi_aws as aws
import pulumi_docker as docker

ecr_repository = aws.ecr.Repository("ecr-repository", name="docker-repository")
auth_token = aws.ecr.get_authorization_token_output(registry_id=ecr_repository.registry_id)
my_app_image = docker.Image("my-app-image",
    build={
        "args": {
            "BUILDKIT_INLINE_CACHE": "1",
        },
        "cache_from": {
            "images": [ecr_repository.repository_url.apply(lambda repository_url: f"{repository_url}:latest")],
        },
        "context": "app/",
        "dockerfile": "app/Dockerfile",
    },
    image_name=ecr_repository.repository_url.apply(lambda repository_url: f"{repository_url}:latest"),
    registry={
        "password": pulumi.Output.secret(auth_token.password),
        "server": ecr_repository.repository_url,
        "username": auth_token.user_name,
    })
pulumi.export("imageName", my_app_image.image_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var ecrRepository = new Aws.Ecr.Repository("ecr-repository", new()
    {
        Name = "docker-repository",
    });

    var authToken = Aws.Ecr.GetAuthorizationToken.Invoke(new()
    {
        RegistryId = ecrRepository.RegistryId,
    });

    var myAppImage = new Docker.Image("my-app-image", new()
    {
        Build = new Docker.Inputs.DockerBuildArgs
        {
            Args = 
            {
                { "BUILDKIT_INLINE_CACHE", "1" },
            },
            CacheFrom = new Docker.Inputs.CacheFromArgs
            {
                Images = new[]
                {
                    ecrRepository.RepositoryUrl.Apply(repositoryUrl => $"{repositoryUrl}:latest"),
                },
            },
            Context = "app/",
            Dockerfile = "app/Dockerfile",
        },
        ImageName = ecrRepository.RepositoryUrl.Apply(repositoryUrl => $"{repositoryUrl}:latest"),
        Registry = new Docker.Inputs.RegistryArgs
        {
            Password = Output.CreateSecret(authToken.Apply(getAuthorizationTokenResult => getAuthorizationTokenResult.Password)),
            Server = ecrRepository.RepositoryUrl,
            Username = authToken.Apply(getAuthorizationTokenResult => getAuthorizationTokenResult.UserName),
        },
    });

    return new Dictionary<string, object?>
    {
        ["imageName"] = myAppImage.ImageName,
    };
});

```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ecrRepository, err := ecr.NewRepository(ctx, "ecr-repository", &ecr.RepositoryArgs{
			Name: pulumi.String("docker-repository"),
		})
		if err != nil {
			return err
		}
		authToken := ecr.GetAuthorizationTokenOutput(ctx, ecr.GetAuthorizationTokenOutputArgs{
			RegistryId: ecrRepository.RegistryId,
		}, nil)
		myAppImage, err := docker.NewImage(ctx, "my-app-image", &docker.ImageArgs{
			Build: &docker.DockerBuildArgs{
				Args: pulumi.StringMap{
					"BUILDKIT_INLINE_CACHE": pulumi.String("1"),
				},
				CacheFrom: &docker.CacheFromArgs{
					Images: pulumi.StringArray{
						ecrRepository.RepositoryUrl.ApplyT(func(repositoryUrl string) (string, error) {
							return fmt.Sprintf("%v:latest", repositoryUrl), nil
						}).(pulumi.StringOutput),
					},
				},
				Context:    pulumi.String("app/"),
				Dockerfile: pulumi.String("app/Dockerfile"),
			},
			ImageName: ecrRepository.RepositoryUrl.ApplyT(func(repositoryUrl string) (string, error) {
				return fmt.Sprintf("%v:latest", repositoryUrl), nil
			}).(pulumi.StringOutput),
			Registry: &docker.RegistryArgs{
				Password: pulumi.ToSecret(authToken.ApplyT(func(authToken ecr.GetAuthorizationTokenResult) (*string, error) {
					return &authToken.Password, nil
				}).(pulumi.StringPtrOutput)).(pulumi.StringOutput),
				Server: ecrRepository.RepositoryUrl,
				Username: authToken.ApplyT(func(authToken ecr.GetAuthorizationTokenResult) (*string, error) {
					return &authToken.UserName, nil
				}).(pulumi.StringPtrOutput),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("imageName", myAppImage.ImageName)
		return nil
	})
}
```
```yaml
config: {}
description: Docker image build using caching with AWS Elastic Container Registry
name: image-caching-yaml
outputs:
    imageName: ${my-app-image.imageName}
resources:
    ecr-repository:
        properties:
            name: docker-repository
        type: aws:ecr:Repository
    my-app-image:
        options:
            version: v4.1.2
        properties:
            build:
                args:
                    BUILDKIT_INLINE_CACHE: "1"
                cacheFrom:
                    images:
                        - ${ecr-repository.repositoryUrl}:latest
                context: app/
                dockerfile: app/Dockerfile
            imageName: ${ecr-repository.repositoryUrl}:latest
            registry:
                password:
                    fn::secret: ${authToken.password}
                server: ${ecr-repository.repositoryUrl}
                username: ${authToken.userName}
        type: docker:Image
runtime: yaml
variables:
    authToken:
        fn::aws:ecr:getAuthorizationToken:
            registryId: ${ecr-repository.registryId}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ecr.Repository;
import com.pulumi.aws.ecr.RepositoryArgs;
import com.pulumi.aws.ecr.EcrFunctions;
import com.pulumi.aws.ecr.inputs.GetAuthorizationTokenArgs;
import com.pulumi.docker.Image;
import com.pulumi.docker.ImageArgs;
import com.pulumi.docker.inputs.DockerBuildArgs;
import com.pulumi.docker.inputs.CacheFromArgs;
import com.pulumi.docker.inputs.RegistryArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var ecrRepository = new Repository("ecrRepository", RepositoryArgs.builder()
            .name("docker-repository")
            .build());

        final var authToken = EcrFunctions.getAuthorizationToken(GetAuthorizationTokenArgs.builder()
            .registryId(ecrRepository.registryId())
            .build());

        var myAppImage = new Image("myAppImage", ImageArgs.builder()
            .build(DockerBuildArgs.builder()
                .args(Map.of("BUILDKIT_INLINE_CACHE", "1"))
                .cacheFrom(CacheFromArgs.builder()
                    .images(ecrRepository.repositoryUrl().applyValue(_repositoryUrl -> String.format("%s:latest", _repositoryUrl)))
                    .build())
                .context("app/")
                .dockerfile("app/Dockerfile")
                .build())
            .imageName(ecrRepository.repositoryUrl().applyValue(_repositoryUrl -> String.format("%s:latest", _repositoryUrl)))
            .registry(RegistryArgs.builder()
                .password(Output.ofSecret(authToken.applyValue(_authToken -> _authToken.password())))
                .server(ecrRepository.repositoryUrl())
                .username(authToken.applyValue(_authToken -> _authToken.userName()))
                .build())
            .build());

        ctx.export("imageName", myAppImage.imageName());
    }
}
```
{{% /example %}}
{{% /examples %}}