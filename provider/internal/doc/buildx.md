{{% examples %}}
## Example Usage
{{% example %}}
### Push to AWS ECR with caching

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as docker from "@pulumi/docker";

const ecrRepository = new aws.ecr.Repository("ecr-repository", {});
const authToken = aws.ecr.getAuthorizationTokenOutput({
    registryId: ecrRepository.registryId,
});
const myImage = new docker.buildx.Image("my-image", {
    cacheFrom: [{
        registry: {
            ref: pulumi.interpolate`${ecrRepository.repositoryUrl}:cache`,
        },
    }],
    cacheTo: [{
        registry: {
            imageManifest: true,
            ociMediaTypes: true,
            ref: pulumi.interpolate`${ecrRepository.repositoryUrl}:cache`,
        },
    }],
    context: {
        location: "./app",
    },
    dockerfile: {
        location: "./Dockerfile",
    },
    push: true,
    registries: [{
        address: ecrRepository.repositoryUrl,
        password: authToken.apply(authToken => authToken.password),
        username: authToken.apply(authToken => authToken.userName),
    }],
    tags: [pulumi.interpolate`${ecrRepository.repositoryUrl}:latest`],
});
```
```python
import pulumi
import pulumi_aws as aws
import pulumi_docker as docker

ecr_repository = aws.ecr.Repository("ecr-repository")
auth_token = aws.ecr.get_authorization_token_output(registry_id=ecr_repository.registry_id)
my_image = docker.buildx.Image("my-image",
    cache_from=[docker.buildx.CacheFromEntryArgs(
        registry=docker.buildx.CacheFromRegistryArgs(
            ref=ecr_repository.repository_url.apply(lambda repository_url: f"{repository_url}:cache"),
        ),
    )],
    cache_to=[docker.buildx.CacheToEntryArgs(
        registry=docker.buildx.CacheToRegistryArgs(
            image_manifest=True,
            oci_media_types=True,
            ref=ecr_repository.repository_url.apply(lambda repository_url: f"{repository_url}:cache"),
        ),
    )],
    context=docker.buildx.BuildContextArgs(
        location="./app",
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        location="./Dockerfile",
    ),
    push=True,
    registries=[docker.buildx.RegistryAuthArgs(
        address=ecr_repository.repository_url,
        password=auth_token.password,
        username=auth_token.user_name,
    )],
    tags=[ecr_repository.repository_url.apply(lambda repository_url: f"{repository_url}:latest")])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var ecrRepository = new Aws.Ecr.Repository("ecr-repository");

    var authToken = Aws.Ecr.GetAuthorizationToken.Invoke(new()
    {
        RegistryId = ecrRepository.RegistryId,
    });

    var myImage = new Docker.Buildx.Image("my-image", new()
    {
        CacheFrom = new[]
        {
            new Docker.Buildx.Inputs.CacheFromEntryArgs
            {
                Registry = new Docker.Buildx.Inputs.CacheFromRegistryArgs
                {
                    Ref = ecrRepository.RepositoryUrl.Apply(repositoryUrl => $"{repositoryUrl}:cache"),
                },
            },
        },
        CacheTo = new[]
        {
            new Docker.Buildx.Inputs.CacheToEntryArgs
            {
                Registry = new Docker.Buildx.Inputs.CacheToRegistryArgs
                {
                    ImageManifest = true,
                    OciMediaTypes = true,
                    Ref = ecrRepository.RepositoryUrl.Apply(repositoryUrl => $"{repositoryUrl}:cache"),
                },
            },
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "./app",
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "./Dockerfile",
        },
        Push = true,
        Registries = new[]
        {
            new Docker.Buildx.Inputs.RegistryAuthArgs
            {
                Address = ecrRepository.RepositoryUrl,
                Password = authToken.Apply(getAuthorizationTokenResult => getAuthorizationTokenResult.Password),
                Username = authToken.Apply(getAuthorizationTokenResult => getAuthorizationTokenResult.UserName),
            },
        },
        Tags = new[]
        {
            ecrRepository.RepositoryUrl.Apply(repositoryUrl => $"{repositoryUrl}:latest"),
        },
    });

});

```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ecrRepository, err := ecr.NewRepository(ctx, "ecr-repository", nil)
		if err != nil {
			return err
		}
		authToken := ecr.GetAuthorizationTokenOutput(ctx, ecr.GetAuthorizationTokenOutputArgs{
			RegistryId: ecrRepository.RegistryId,
		}, nil)
		_, err = buildx.NewImage(ctx, "my-image", &buildx.ImageArgs{
			CacheFrom: buildx.CacheFromEntryArray{
				&buildx.CacheFromEntryArgs{
					Registry: &buildx.CacheFromRegistryArgs{
						Ref: ecrRepository.RepositoryUrl.ApplyT(func(repositoryUrl string) (string, error) {
							return fmt.Sprintf("%v:cache", repositoryUrl), nil
						}).(pulumi.StringOutput),
					},
				},
			},
			CacheTo: buildx.CacheToEntryArray{
				&buildx.CacheToEntryArgs{
					Registry: &buildx.CacheToRegistryArgs{
						ImageManifest: pulumi.Bool(true),
						OciMediaTypes: pulumi.Bool(true),
						Ref: ecrRepository.RepositoryUrl.ApplyT(func(repositoryUrl string) (string, error) {
							return fmt.Sprintf("%v:cache", repositoryUrl), nil
						}).(pulumi.StringOutput),
					},
				},
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("./app"),
			},
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("./Dockerfile"),
			},
			Push: pulumi.Bool(true),
			Registries: buildx.RegistryAuthArray{
				&buildx.RegistryAuthArgs{
					Address: ecrRepository.RepositoryUrl,
					Password: authToken.ApplyT(func(authToken ecr.GetAuthorizationTokenResult) (*string, error) {
						return &authToken.Password, nil
					}).(pulumi.StringPtrOutput),
					Username: authToken.ApplyT(func(authToken ecr.GetAuthorizationTokenResult) (*string, error) {
						return &authToken.UserName, nil
					}).(pulumi.StringPtrOutput),
				},
			},
			Tags: pulumi.StringArray{
				ecrRepository.RepositoryUrl.ApplyT(func(repositoryUrl string) (string, error) {
					return fmt.Sprintf("%v:latest", repositoryUrl), nil
				}).(pulumi.StringOutput),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Push to AWS ECR with caching
name: ecr
resources:
    ecr-repository:
        type: aws:ecr:Repository
    my-image:
        properties:
            cacheFrom:
                - registry:
                    ref: ${ecr-repository.repositoryUrl}:cache
            cacheTo:
                - registry:
                    imageManifest: true
                    ociMediaTypes: true
                    ref: ${ecr-repository.repositoryUrl}:cache
            context:
                location: ./app
            dockerfile:
                location: ./Dockerfile
            push: true
            registries:
                - address: ${ecr-repository.repositoryUrl}
                  password: ${auth-token.password}
                  username: ${auth-token.userName}
            tags:
                - ${ecr-repository.repositoryUrl}:latest
        type: docker:buildx/image:Image
runtime: yaml
variables:
    auth-token:
        fn::aws:ecr:getAuthorizationToken:
            registryId: ${ecr-repository.registryId}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ecr.Repository;
import com.pulumi.aws.ecr.EcrFunctions;
import com.pulumi.aws.ecr.inputs.GetAuthorizationTokenArgs;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.CacheFromEntryArgs;
import com.pulumi.docker.buildx.inputs.CacheFromRegistryArgs;
import com.pulumi.docker.buildx.inputs.CacheToEntryArgs;
import com.pulumi.docker.buildx.inputs.CacheToRegistryArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
import com.pulumi.docker.buildx.inputs.RegistryAuthArgs;
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
        var ecrRepository = new Repository("ecrRepository");

        final var authToken = EcrFunctions.getAuthorizationToken(GetAuthorizationTokenArgs.builder()
            .registryId(ecrRepository.registryId())
            .build());

        var myImage = new Image("myImage", ImageArgs.builder()        
            .cacheFrom(CacheFromEntryArgs.builder()
                .registry(CacheFromRegistryArgs.builder()
                    .ref(ecrRepository.repositoryUrl().applyValue(repositoryUrl -> String.format("%s:cache", repositoryUrl)))
                    .build())
                .build())
            .cacheTo(CacheToEntryArgs.builder()
                .registry(CacheToRegistryArgs.builder()
                    .imageManifest(true)
                    .ociMediaTypes(true)
                    .ref(ecrRepository.repositoryUrl().applyValue(repositoryUrl -> String.format("%s:cache", repositoryUrl)))
                    .build())
                .build())
            .context(BuildContextArgs.builder()
                .location("./app")
                .build())
            .dockerfile(DockerfileArgs.builder()
                .location("./Dockerfile")
                .build())
            .push(true)
            .registries(RegistryAuthArgs.builder()
                .address(ecrRepository.repositoryUrl())
                .password(authToken.applyValue(getAuthorizationTokenResult -> getAuthorizationTokenResult).applyValue(authToken -> authToken.applyValue(getAuthorizationTokenResult -> getAuthorizationTokenResult.password())))
                .username(authToken.applyValue(getAuthorizationTokenResult -> getAuthorizationTokenResult).applyValue(authToken -> authToken.applyValue(getAuthorizationTokenResult -> getAuthorizationTokenResult.userName())))
                .build())
            .tags(ecrRepository.repositoryUrl().applyValue(repositoryUrl -> String.format("%s:latest", repositoryUrl)))
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Multi-platform image

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "app",
    },
    platforms: [
        docker.buildx.image.Platform.Plan9_amd64,
        docker.buildx.image.Platform.Plan9_386,
    ],
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    platforms=[
        docker.buildx/image.Platform.PLAN9_AMD64,
        docker.buildx/image.Platform.PLAN9_386,
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Platforms = new[]
        {
            Docker.Buildx.Image.Platform.Plan9_amd64,
            Docker.Buildx.Image.Platform.Plan9_386,
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Platforms: buildx.PlatformArray{
				buildx.Platform_Plan9_amd64,
				buildx.Platform_Plan9_386,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Multi-platform image
name: multi-platform
resources:
    image:
        properties:
            context:
                location: app
            platforms:
                - plan9/amd64
                - plan9/386
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .platforms(            
                "plan9/amd64",
                "plan9/386")
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Registry export

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "app",
    },
    push: true,
    registries: [{
        address: "docker.io",
        password: dockerHubPassword,
        username: "pulumibot",
    }],
    tags: ["docker.io/pulumi/pulumi:3.107.0"],
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    push=True,
    registries=[docker.buildx.RegistryAuthArgs(
        address="docker.io",
        password=docker_hub_password,
        username="pulumibot",
    )],
    tags=["docker.io/pulumi/pulumi:3.107.0"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Push = true,
        Registries = new[]
        {
            new Docker.Buildx.Inputs.RegistryAuthArgs
            {
                Address = "docker.io",
                Password = dockerHubPassword,
                Username = "pulumibot",
            },
        },
        Tags = new[]
        {
            "docker.io/pulumi/pulumi:3.107.0",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Push: pulumi.Bool(true),
			Registries: buildx.RegistryAuthArray{
				&buildx.RegistryAuthArgs{
					Address:  pulumi.String("docker.io"),
					Password: pulumi.Any(dockerHubPassword),
					Username: pulumi.String("pulumibot"),
				},
			},
			Tags: pulumi.StringArray{
				pulumi.String("docker.io/pulumi/pulumi:3.107.0"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Registry export
name: registry
resources:
    image:
        properties:
            context:
                location: app
            push: true
            registries:
                - address: docker.io
                  password: ${dockerHubPassword}
                  username: pulumibot
            tags:
                - docker.io/pulumi/pulumi:3.107.0
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.RegistryAuthArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .push(true)
            .registries(RegistryAuthArgs.builder()
                .address("docker.io")
                .password(dockerHubPassword)
                .username("pulumibot")
                .build())
            .tags("docker.io/pulumi/pulumi:3.107.0")
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Caching

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    cacheFrom: [{
        local: {
            src: "tmp/cache",
        },
    }],
    cacheTo: [{
        local: {
            dest: "tmp/cache",
            mode: docker.buildx.image.CacheMode.Max,
        },
    }],
    context: {
        location: "app",
    },
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    cache_from=[docker.buildx.CacheFromEntryArgs(
        local=docker.buildx.CacheFromLocalArgs(
            src="tmp/cache",
        ),
    )],
    cache_to=[docker.buildx.CacheToEntryArgs(
        local=docker.buildx.CacheToLocalArgs(
            dest="tmp/cache",
            mode=docker.buildx/image.CacheMode.MAX,
        ),
    )],
    context=docker.buildx.BuildContextArgs(
        location="app",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        CacheFrom = new[]
        {
            new Docker.Buildx.Inputs.CacheFromEntryArgs
            {
                Local = new Docker.Buildx.Inputs.CacheFromLocalArgs
                {
                    Src = "tmp/cache",
                },
            },
        },
        CacheTo = new[]
        {
            new Docker.Buildx.Inputs.CacheToEntryArgs
            {
                Local = new Docker.Buildx.Inputs.CacheToLocalArgs
                {
                    Dest = "tmp/cache",
                    Mode = Docker.Buildx.Image.CacheMode.Max,
                },
            },
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			CacheFrom: buildx.CacheFromEntryArray{
				&buildx.CacheFromEntryArgs{
					Local: &buildx.CacheFromLocalArgs{
						Src: pulumi.String("tmp/cache"),
					},
				},
			},
			CacheTo: buildx.CacheToEntryArray{
				&buildx.CacheToEntryArgs{
					Local: &buildx.CacheToLocalArgs{
						Dest: pulumi.String("tmp/cache"),
						Mode: buildx.CacheModeMax,
					},
				},
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Caching
name: caching
resources:
    image:
        properties:
            cacheFrom:
                - local:
                    src: tmp/cache
            cacheTo:
                - local:
                    dest: tmp/cache
                    mode: max
            context:
                location: app
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.CacheFromEntryArgs;
import com.pulumi.docker.buildx.inputs.CacheFromLocalArgs;
import com.pulumi.docker.buildx.inputs.CacheToEntryArgs;
import com.pulumi.docker.buildx.inputs.CacheToLocalArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .cacheFrom(CacheFromEntryArgs.builder()
                .local(CacheFromLocalArgs.builder()
                    .src("tmp/cache")
                    .build())
                .build())
            .cacheTo(CacheToEntryArgs.builder()
                .local(CacheToLocalArgs.builder()
                    .dest("tmp/cache")
                    .mode("max")
                    .build())
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Build arguments

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    buildArgs: {
        SET_ME_TO_TRUE: "true",
    },
    context: {
        location: "app",
    },
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    build_args={
        "SET_ME_TO_TRUE": "true",
    },
    context=docker.buildx.BuildContextArgs(
        location="app",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        BuildArgs = 
        {
            { "SET_ME_TO_TRUE", "true" },
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			BuildArgs: pulumi.StringMap{
				"SET_ME_TO_TRUE": pulumi.String("true"),
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Build arguments
name: build-args
resources:
    image:
        properties:
            buildArgs:
                SET_ME_TO_TRUE: "true"
            context:
                location: app
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .buildArgs(Map.of("SET_ME_TO_TRUE", "true"))
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Build targets

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "app",
    },
    targets: [
        "build-me",
        "also-build-me",
    ],
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    targets=[
        "build-me",
        "also-build-me",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Targets = new[]
        {
            "build-me",
            "also-build-me",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Targets: pulumi.StringArray{
				pulumi.String("build-me"),
				pulumi.String("also-build-me"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Build targets
name: build-targets
resources:
    image:
        properties:
            context:
                location: app
            targets:
                - build-me
                - also-build-me
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .targets(            
                "build-me",
                "also-build-me")
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Named contexts

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {context: {
    location: "app",
    named: {
        "golang:latest": {
            location: "docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984",
        },
    },
}});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image", context=docker.buildx.BuildContextArgs(
    location="app",
    named={
        "golang:latest": docker.buildx.ContextArgs(
            location="docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984",
        ),
    },
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
            Named = 
            {
                { "golang:latest", new Docker.Buildx.Inputs.ContextArgs
                {
                    Location = "docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984",
                } },
            },
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
				Named: buildx.ContextMap{
					"golang:latest": &buildx.ContextArgs{
						Location: pulumi.String("docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Named contexts
name: named-contexts
resources:
    image:
        properties:
            context:
                location: app
                named:
                    golang:latest:
                        location: docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .named(Map.of("golang:latest", Map.of("location", "docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984")))
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Remote context

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {context: {
    location: "https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile",
}});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image", context=docker.buildx.BuildContextArgs(
    location="https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile",
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Remote context
name: remote-context
resources:
    image:
        properties:
            context:
                location: https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile")
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Inline Dockerfile

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "app",
    },
    dockerfile: {
        inline: `FROM busybox
COPY hello.c ./
`,
    },
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        inline="""FROM busybox
COPY hello.c ./
""",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Inline = @"FROM busybox
COPY hello.c ./
",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Dockerfile: &buildx.DockerfileArgs{
				Inline: pulumi.String("FROM busybox\nCOPY hello.c ./\n"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Inline Dockerfile
name: inline
resources:
    image:
        properties:
            context:
                location: app
            dockerfile:
                inline: |
                    FROM busybox
                    COPY hello.c ./
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .dockerfile(DockerfileArgs.builder()
                .inline("""
FROM busybox
COPY hello.c ./
                """)
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Remote context

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "https://github.com/docker-library/hello-world.git",
    },
    dockerfile: {
        location: "app/Dockerfile",
    },
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="https://github.com/docker-library/hello-world.git",
    ),
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "https://github.com/docker-library/hello-world.git",
        },
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile",
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("https://github.com/docker-library/hello-world.git"),
			},
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Remote context
name: remote-context
resources:
    image:
        properties:
            context:
                location: https://github.com/docker-library/hello-world.git
            dockerfile:
                location: app/Dockerfile
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("https://github.com/docker-library/hello-world.git")
                .build())
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile")
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Local export

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const image = new docker.buildx.Image("image", {
    context: {
        location: "app",
    },
    exports: [{
        docker: {
            tar: true,
        },
    }],
});
```
```python
import pulumi
import pulumi_docker as docker

image = docker.buildx.Image("image",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    exports=[docker.buildx.ExportEntryArgs(
        docker=docker.buildx.ExportDockerArgs(
            tar=True,
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var image = new Docker.Buildx.Image("image", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Exports = new[]
        {
            new Docker.Buildx.Inputs.ExportEntryArgs
            {
                Docker = new Docker.Buildx.Inputs.ExportDockerArgs
                {
                    Tar = true,
                },
            },
        },
    });

});

```
```go
package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := buildx.NewImage(ctx, "image", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Exports: buildx.ExportEntryArray{
				&buildx.ExportEntryArgs{
					Docker: &buildx.ExportDockerArgs{
						Tar: pulumi.Bool(true),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```yaml
description: Local export
name: docker-load
resources:
    image:
        properties:
            context:
                location: app
            exports:
                - docker:
                    tar: true
        type: docker:buildx/image:Image
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.ExportEntryArgs;
import com.pulumi.docker.buildx.inputs.ExportDockerArgs;
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
        var image = new Image("image", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .exports(ExportEntryArgs.builder()
                .docker(ExportDockerArgs.builder()
                    .tar(true)
                    .build())
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}