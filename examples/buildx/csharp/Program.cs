using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Docker = Pulumi.Docker;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var dockerHubPassword = config.Require("dockerHubPassword");
    var multiPlatform = new Docker.Buildx.Image("multiPlatform", new()
    {
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile.multiPlatform",
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Platforms = new[]
        {
            Docker.Buildx.Platform.Plan9_amd64,
            Docker.Buildx.Platform.Plan9_386,
        },
    });

    var registryPush = new Docker.Buildx.Image("registryPush", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Tags = new[]
        {
            "docker.io/pulumibot/buildkit-e2e:example",
        },
        Exports = new[]
        {
            new Docker.Buildx.Inputs.ExportEntryArgs
            {
                Registry = new Docker.Buildx.Inputs.ExportRegistryArgs
                {
                    OciMediaTypes = true,
                    Push = false,
                },
            },
        },
        Registries = new[]
        {
            new Docker.Buildx.Inputs.RegistryAuthArgs
            {
                Address = "docker.io",
                Username = "pulumibot",
                Password = dockerHubPassword,
            },
        },
    });

    var cached = new Docker.Buildx.Image("cached", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        CacheTo = new[]
        {
            new Docker.Buildx.Inputs.CacheToEntryArgs
            {
                Local = new Docker.Buildx.Inputs.CacheToLocalArgs
                {
                    Dest = "tmp/cache",
                    Mode = Docker.Buildx.CacheMode.Max,
                },
            },
        },
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
    });

    var buildArgs = new Docker.Buildx.Image("buildArgs", new()
    {
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile.buildArgs",
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        BuildArgs = 
        {
            { "SET_ME_TO_TRUE", "true" },
        },
    });

    var extraHosts = new Docker.Buildx.Image("extraHosts", new()
    {
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile.extraHosts",
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        AddHosts = new[]
        {
            "metadata.google.internal:169.254.169.254",
        },
    });

    var sshMount = new Docker.Buildx.Image("sshMount", new()
    {
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile.sshMount",
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Ssh = new[]
        {
            new Docker.Buildx.Inputs.SSHArgs
            {
                Id = "default",
            },
        },
    });

    var secrets = new Docker.Buildx.Image("secrets", new()
    {
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile.secrets",
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Secrets = 
        {
            { "password", "hunter2" },
        },
    });

    var labels = new Docker.Buildx.Image("labels", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
        Labels = 
        {
            { "description", "This image will get a descriptive label 👍" },
        },
    });

    var targets = new Docker.Buildx.Image("targets", new()
    {
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile.targets",
        },
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

    var namedContexts = new Docker.Buildx.Image("namedContexts", new()
    {
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Location = "app/Dockerfile.namedContexts",
        },
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

    var remoteContext = new Docker.Buildx.Image("remoteContext", new()
    {
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile",
        },
    });

    var remoteContextWithInline = new Docker.Buildx.Image("remoteContextWithInline", new()
    {
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Inline = @"FROM busybox
COPY hello.c ./
",
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "https://github.com/docker-library/hello-world.git",
        },
    });

    var inline = new Docker.Buildx.Image("inline", new()
    {
        Dockerfile = new Docker.Buildx.Inputs.DockerfileArgs
        {
            Inline = @"FROM alpine
RUN echo ""This uses an inline Dockerfile! 👍""
",
        },
        Context = new Docker.Buildx.Inputs.BuildContextArgs
        {
            Location = "app",
        },
    });

    var dockerLoad = new Docker.Buildx.Image("dockerLoad", new()
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

    return new Dictionary<string, object?>
    {
        ["platforms"] = multiPlatform.Platforms,
    };
});

