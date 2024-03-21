{{% examples %}}
## Example Usage
{{% example %}}
### Multi-platform registry caching

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const amd64 = new docker.buildx.Image("amd64", {
    cacheFrom: [{
        registry: {
            ref: "docker.io/pulumi/pulumi:cache-amd64",
        },
    }],
    cacheTo: [{
        registry: {
            mode: docker.buildx.image.CacheMode.Max,
            ref: "docker.io/pulumi/pulumi:cache-amd64",
        },
    }],
    context: {
        location: "app",
    },
    platforms: [docker.buildx.image.Platform.Linux_amd64],
    tags: ["docker.io/pulumi/pulumi:3.107.0-amd64"],
});
const arm64 = new docker.buildx.Image("arm64", {
    cacheFrom: [{
        registry: {
            ref: "docker.io/pulumi/pulumi:cache-arm64",
        },
    }],
    cacheTo: [{
        registry: {
            mode: docker.buildx.image.CacheMode.Max,
            ref: "docker.io/pulumi/pulumi:cache-arm64",
        },
    }],
    context: {
        location: "app",
    },
    platforms: [docker.buildx.image.Platform.Linux_arm64],
    tags: ["docker.io/pulumi/pulumi:3.107.0-arm64"],
});
const index = new docker.buildx.Index("index", {
    sources: [
        amd64.ref,
        arm64.ref,
    ],
    tag: "docker.io/pulumi/pulumi:3.107.0",
});
export const ref = index.ref;
```
```python
import pulumi
import pulumi_docker as docker

amd64 = docker.buildx.Image("amd64",
    cache_from=[docker.buildx.CacheFromArgs(
        registry=docker.buildx.CacheFromRegistryArgs(
            ref="docker.io/pulumi/pulumi:cache-amd64",
        ),
    )],
    cache_to=[docker.buildx.CacheToArgs(
        registry=docker.buildx.CacheToRegistryArgs(
            mode=docker.buildx/image.CacheMode.MAX,
            ref="docker.io/pulumi/pulumi:cache-amd64",
        ),
    )],
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    platforms=[docker.buildx/image.Platform.LINUX_AMD64],
    tags=["docker.io/pulumi/pulumi:3.107.0-amd64"])
arm64 = docker.buildx.Image("arm64",
    cache_from=[docker.buildx.CacheFromArgs(
        registry=docker.buildx.CacheFromRegistryArgs(
            ref="docker.io/pulumi/pulumi:cache-arm64",
        ),
    )],
    cache_to=[docker.buildx.CacheToArgs(
        registry=docker.buildx.CacheToRegistryArgs(
            mode=docker.buildx/image.CacheMode.MAX,
            ref="docker.io/pulumi/pulumi:cache-arm64",
        ),
    )],
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    platforms=[docker.buildx/image.Platform.LINUX_ARM64],
    tags=["docker.io/pulumi/pulumi:3.107.0-arm64"])
index = docker.buildx.Index("index",
    sources=[
        amd64.ref,
        arm64.ref,
    ],
    tag="docker.io/pulumi/pulumi:3.107.0")
pulumi.export("ref", index.ref)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var amd64 = new Docker.Buildx.Image("amd64", new()
    {
        CacheFrom = new[]
        {
            new Docker.Buildx.Inputs.CacheFromArgs
            {
                Registry = new Docker.Buildx.Inputs.CacheFromRegistryArgs
                {
                    Ref = "docker.io/pulumi/pulumi:cache-amd64",
                },
            },
        },
        CacheTo = new[]
        {
            new Docker.Buildx.Inputs.CacheToArgs
            {
                Registry = new Docker.Buildx.Inputs.CacheToRegistryArgs
                {
                    Mode = Docker.Buildx.Image.CacheMode.Max,
                    Ref = "docker.io/pulumi/pulumi:cache-amd64",
                },
            },
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Platforms = new[]
        {
            Docker.Buildx.Image.Platform.Linux_amd64,
        },
        Tags = new[]
        {
            "docker.io/pulumi/pulumi:3.107.0-amd64",
        },
    });

    var arm64 = new Docker.Buildx.Image("arm64", new()
    {
        CacheFrom = new[]
        {
            new Docker.Buildx.Inputs.CacheFromArgs
            {
                Registry = new Docker.Buildx.Inputs.CacheFromRegistryArgs
                {
                    Ref = "docker.io/pulumi/pulumi:cache-arm64",
                },
            },
        },
        CacheTo = new[]
        {
            new Docker.Buildx.Inputs.CacheToArgs
            {
                Registry = new Docker.Buildx.Inputs.CacheToRegistryArgs
                {
                    Mode = Docker.Buildx.Image.CacheMode.Max,
                    Ref = "docker.io/pulumi/pulumi:cache-arm64",
                },
            },
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Platforms = new[]
        {
            Docker.Buildx.Image.Platform.Linux_arm64,
        },
        Tags = new[]
        {
            "docker.io/pulumi/pulumi:3.107.0-arm64",
        },
    });

    var index = new Docker.Buildx.Index("index", new()
    {
        Sources = new[]
        {
            amd64.Ref,
            arm64.Ref,
        },
        Tag = "docker.io/pulumi/pulumi:3.107.0",
    });

    return new Dictionary<string, object?>
    {
        ["ref"] = index.Ref,
    };
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
		amd64, err := buildx.NewImage(ctx, "amd64", &buildx.ImageArgs{
			CacheFrom: buildx.CacheFromArray{
				&buildx.CacheFromArgs{
					Registry: &buildx.CacheFromRegistryArgs{
						Ref: pulumi.String("docker.io/pulumi/pulumi:cache-amd64"),
					},
				},
			},
			CacheTo: buildx.CacheToArray{
				&buildx.CacheToArgs{
					Registry: &buildx.CacheToRegistryArgs{
						Mode: buildx.CacheModeMax,
						Ref:  pulumi.String("docker.io/pulumi/pulumi:cache-amd64"),
					},
				},
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Platforms: buildx.PlatformArray{
				buildx.Platform_Linux_amd64,
			},
			Tags: pulumi.StringArray{
				pulumi.String("docker.io/pulumi/pulumi:3.107.0-amd64"),
			},
		})
		if err != nil {
			return err
		}
		arm64, err := buildx.NewImage(ctx, "arm64", &buildx.ImageArgs{
			CacheFrom: buildx.CacheFromArray{
				&buildx.CacheFromArgs{
					Registry: &buildx.CacheFromRegistryArgs{
						Ref: pulumi.String("docker.io/pulumi/pulumi:cache-arm64"),
					},
				},
			},
			CacheTo: buildx.CacheToArray{
				&buildx.CacheToArgs{
					Registry: &buildx.CacheToRegistryArgs{
						Mode: buildx.CacheModeMax,
						Ref:  pulumi.String("docker.io/pulumi/pulumi:cache-arm64"),
					},
				},
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Platforms: buildx.PlatformArray{
				buildx.Platform_Linux_arm64,
			},
			Tags: pulumi.StringArray{
				pulumi.String("docker.io/pulumi/pulumi:3.107.0-arm64"),
			},
		})
		if err != nil {
			return err
		}
		index, err := buildx.NewIndex(ctx, "index", &buildx.IndexArgs{
			Sources: pulumi.StringArray{
				amd64.Ref,
				arm64.Ref,
			},
			Tag: pulumi.String("docker.io/pulumi/pulumi:3.107.0"),
		})
		if err != nil {
			return err
		}
		ctx.Export("ref", index.Ref)
		return nil
	})
}
```
```yaml
description: Multi-platform registry caching
name: registry-caching
outputs:
    ref: ${index.ref}
resources:
    amd64:
        properties:
            cacheFrom:
                - registry:
                    ref: docker.io/pulumi/pulumi:cache-amd64
            cacheTo:
                - registry:
                    mode: max
                    ref: docker.io/pulumi/pulumi:cache-amd64
            context:
                location: app
            platforms:
                - linux/amd64
            tags:
                - docker.io/pulumi/pulumi:3.107.0-amd64
        type: docker:buildx/image:Image
    arm64:
        properties:
            cacheFrom:
                - registry:
                    ref: docker.io/pulumi/pulumi:cache-arm64
            cacheTo:
                - registry:
                    mode: max
                    ref: docker.io/pulumi/pulumi:cache-arm64
            context:
                location: app
            platforms:
                - linux/arm64
            tags:
                - docker.io/pulumi/pulumi:3.107.0-arm64
        type: docker:buildx/image:Image
    index:
        properties:
            sources:
                - ${amd64.ref}
                - ${arm64.ref}
            tag: docker.io/pulumi/pulumi:3.107.0
        type: docker:buildx/image:Index
runtime: yaml
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.CacheFromArgs;
import com.pulumi.docker.buildx.inputs.CacheFromRegistryArgs;
import com.pulumi.docker.buildx.inputs.CacheToArgs;
import com.pulumi.docker.buildx.inputs.CacheToRegistryArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.Index;
import com.pulumi.docker.buildx.IndexArgs;
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
        var amd64 = new Image("amd64", ImageArgs.builder()        
            .cacheFrom(CacheFromArgs.builder()
                .registry(CacheFromRegistryArgs.builder()
                    .ref("docker.io/pulumi/pulumi:cache-amd64")
                    .build())
                .build())
            .cacheTo(CacheToArgs.builder()
                .registry(CacheToRegistryArgs.builder()
                    .mode("max")
                    .ref("docker.io/pulumi/pulumi:cache-amd64")
                    .build())
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .platforms("linux/amd64")
            .tags("docker.io/pulumi/pulumi:3.107.0-amd64")
            .build());

        var arm64 = new Image("arm64", ImageArgs.builder()        
            .cacheFrom(CacheFromArgs.builder()
                .registry(CacheFromRegistryArgs.builder()
                    .ref("docker.io/pulumi/pulumi:cache-arm64")
                    .build())
                .build())
            .cacheTo(CacheToArgs.builder()
                .registry(CacheToRegistryArgs.builder()
                    .mode("max")
                    .ref("docker.io/pulumi/pulumi:cache-arm64")
                    .build())
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .platforms("linux/arm64")
            .tags("docker.io/pulumi/pulumi:3.107.0-arm64")
            .build());

        var index = new Index("index", IndexArgs.builder()        
            .sources(            
                amd64.ref(),
                arm64.ref())
            .tag("docker.io/pulumi/pulumi:3.107.0")
            .build());

        ctx.export("ref", index.ref());
    }
}
```
{{% /example %}}
{{% /examples %}}