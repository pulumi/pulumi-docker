package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.docker.buildx.Image;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.inputs.DockerfileArgs;
import com.pulumi.docker.buildx.inputs.BuildContextArgs;
import com.pulumi.docker.buildx.inputs.ExportArgs;
import com.pulumi.docker.buildx.inputs.ExportRegistryArgs;
import com.pulumi.docker.buildx.inputs.RegistryAuthArgs;
import com.pulumi.docker.buildx.inputs.CacheToArgs;
import com.pulumi.docker.buildx.inputs.CacheToLocalArgs;
import com.pulumi.docker.buildx.inputs.CacheFromArgs;
import com.pulumi.docker.buildx.inputs.CacheFromLocalArgs;
import com.pulumi.docker.buildx.inputs.SSHArgs;
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
        final var config = ctx.config();
        final var dockerHubPassword = config.get("dockerHubPassword");
        var multiPlatform = new Image("multiPlatform", ImageArgs.builder()        
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile.multiPlatform")
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .platforms(            
                "plan9/amd64",
                "plan9/386")
            .build());

        var registryPush = new Image("registryPush", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .tags("docker.io/pulumibot/buildkit-e2e:example")
            .exports(ExportArgs.builder()
                .registry(ExportRegistryArgs.builder()
                    .ociMediaTypes(true)
                    .push(false)
                    .build())
                .build())
            .registries(RegistryAuthArgs.builder()
                .address("docker.io")
                .username("pulumibot")
                .password(dockerHubPassword)
                .build())
            .build());

        var cached = new Image("cached", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .cacheTo(CacheToArgs.builder()
                .local(CacheToLocalArgs.builder()
                    .dest("tmp/cache")
                    .mode("max")
                    .build())
                .build())
            .cacheFrom(CacheFromArgs.builder()
                .local(CacheFromLocalArgs.builder()
                    .src("tmp/cache")
                    .build())
                .build())
            .build());

        var buildArgs = new Image("buildArgs", ImageArgs.builder()        
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile.buildArgs")
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .buildArgs(Map.of("SET_ME_TO_TRUE", "true"))
            .build());

        var extraHosts = new Image("extraHosts", ImageArgs.builder()        
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile.extraHosts")
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .addHosts("metadata.google.internal:169.254.169.254")
            .build());

        var sshMount = new Image("sshMount", ImageArgs.builder()        
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile.sshMount")
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .ssh(SSHArgs.builder()
                .id("default")
                .build())
            .build());

        var secrets = new Image("secrets", ImageArgs.builder()        
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile.secrets")
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .secrets(Map.of("password", "hunter2"))
            .build());

        var labels = new Image("labels", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .labels(Map.of("description", "This image will get a descriptive label 👍"))
            .build());

        var target = new Image("target", ImageArgs.builder()        
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile.target")
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .target("build-me")
            .build());

        var namedContexts = new Image("namedContexts", ImageArgs.builder()        
            .dockerfile(DockerfileArgs.builder()
                .location("app/Dockerfile.namedContexts")
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .named(Map.of("golang:latest", Map.of("location", "docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984")))
                .build())
            .build());

        var remoteContext = new Image("remoteContext", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile")
                .build())
            .build());

        var remoteContextWithInline = new Image("remoteContextWithInline", ImageArgs.builder()        
            .dockerfile(DockerfileArgs.builder()
                .inline(
                    "FROM busybox\n" +
                    "COPY hello.c ./")
                .build())
            .context(BuildContextArgs.builder()
                .location("https://github.com/docker-library/hello-world.git")
                .build())
            .build());

        var inline = new Image("inline", ImageArgs.builder()        
            .dockerfile(DockerfileArgs.builder()
                .inline(
                    "FROM alpine\n" +
                    "RUN echo 'This uses an inline Dockerfile! 👍'")
                .build())
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .build());

        var dockerLoad = new Image("dockerLoad", ImageArgs.builder()        
            .context(BuildContextArgs.builder()
                .location("app")
                .build())
            .exports(ExportArgs.builder()
                .docker(ExportDockerArgs.builder()
                    .tar(true)
                    .build())
                .build())
            .build());

        ctx.export("platforms", multiPlatform.platforms());
    }
}
