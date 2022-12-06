{{% examples %}}
## Example Usage
{{% example %}}
### A minimal Pulumi YAML program

```typescript
import * as pulumi from "@pulumi/pulumi";

export const imageName = guinsImage.imageName;
```
```python
import pulumi

pulumi.export("imageName", guins_image["imageName"])
```
```csharp
using System.Collections.Generic;
using Pulumi;

return await Deployment.RunAsync(() => 
{
    return new Dictionary<string, object?>
    {
        ["imageName"] = guinsImage.ImageName,
    };
});

```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ctx.Export("imageName", guinsImage.ImageName)
		return nil
	})
}
```
```yaml
config: {}
description: A minimal Pulumi YAML program
name: image-yaml
outputs:
    imageName: ${guins-image.imageName}
resources:
    guins-image:
        properties:
            build:
                context: .
                dockerfile: Dockerfile
            imageName: gsaenger/test-yaml:tag1
            skipPush: true
        type: docker:Image
runtime: yaml
variables: {}
```
{{% /example %}}
{{% /examples %}}
