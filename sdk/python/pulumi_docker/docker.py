# Copyright 2016-2020, Pulumi Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License")
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
import asyncio
import json
import math
import os
import re
from random import random
from typing import Optional, Union, List, Mapping, Sequence
from distutils.version import LooseVersion

import pulumi
from .utils import get_image_name_and_tag


class Registry:
    registry: pulumi.Input[str]
    """
    Registry server url
    """
    username: pulumi.Input[str]
    """
    Username for the registry
    """
    password: pulumi.Input[str]
    """
    Password for the registry
    """

    def __init__(self, registry: pulumi.Input[str], username: pulumi.Input[str], password: pulumi.Input[str]):
        """
        Registry is the information required to login to a Docker registry.


        :param pulumi.Input[str] registry: Registry server url
        :param pulumi.Input[str] username: Username for the registry
        :param pulumi.Input[str] password: Password for the registry
        """
        self.registry = registry
        self.username = username
        self.password = password


class CacheFrom:
    """
    CacheFrom may be used to specify build stages to use for the Docker build cache. The final image
    is always implicitly included.
    """

    stages: Optional[Sequence[pulumi.Input[pulumi.Input[str]]]]
    """
    An optional list of build stages to use for caching. Each build stage in this list will be
    built explicitly and pushed to the target repository. A given stage's image will be tagged as
    "[stage-name]".
    """

    def __init__(self, stages: Optional[Sequence[pulumi.Input[pulumi.Input[str]]]] = None):
        self.stages = stages


class DockerBuild:
    context: Optional[pulumi.Input[str]]
    """
    context is a path to a directory to use for the Docker build context, usually the directory
    in which the Dockerfile resides (although dockerfile may be used to choose a custom location
    independent of this choice). If not specified, the context defaults to the current working
    directory if a relative path is used, it is relative to the current working directory that
    Pulumi is evaluating.
    """

    dockerfile: Optional[pulumi.Input[str]]
    """
    dockerfile may be used to override the default Dockerfile name and/or location.  By default,
    it is assumed to be a file named Dockerfile in the root of the build context.
    """

    args: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]
    """
    An optional map of named build-time argument variables to set during the Docker build.  This
    flag allows you to pass built-time variables that can be accessed like environment variables
    inside the `RUN` instruction.
    """

    cache_from: Optional[pulumi.Input[Union[bool, CacheFrom]]]
    """
    An optional CacheFrom object with information about the build stages to use for the Docker
    build cache. This parameter maps to the --cache-from argument to the Docker CLI. If this
    parameter is `true`, only the final image will be pulled and passed to --cache-from if it is
    a CacheFrom object, the stages named therein will also be pulled and passed to --cache-from.
    """

    extra_options: Optional[Sequence[pulumi.Input[pulumi.Input[str]]]]
    """
    An optional catch-all list of arguments to provide extra CLI options to the docker build command.  For
    example `['--network', 'host']`.
    """

    env: Optional[Mapping[str, str]]
    """
    Environment variables to set on the invocation of `docker build`, for example to support
    `DOCKER_BUILDKIT=1 docker build`.
    """

    target: Optional[pulumi.Input[str]]
    """
    The target of the dockerfile to build
    """

    def __init__(self, context=None, dockerfile=None, args=None, cache_from=None, extra_options=None, env=None,
                 target=None):
        """
        DockerBuild may be used to specify detailed instructions about how to build a container.


        :param Optional[pulumi.Input[str]] context: context is a path to a directory to use for the Docker build
            context, usually the directory in which the Dockerfile resides (although dockerfile may be used to choose
            a custom location independent of this choice). If not specified, the context defaults to the current working
            directory if a relative path is used, it is relative to the current working directory that
            Pulumi is evaluating.
        :param Optional[pulumi.Input[str]] dockerfile: dockerfile may be used to override the default Dockerfile name
            and/or location.  By default, it is assumed to be a file named Dockerfile in the root of the build context.
        :param Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] args: An optional map of named build-time
            argument variables to set during the Docker build.  This flag allows you to pass built-time variables that
            can be accessed like environment variables inside the `RUN` instruction.
        :param Optional[pulumi.Input[Union[bool, CacheFrom]]] cache_from: An optional CacheFrom object with information
            about the build stages to use for the Docker build cache. This parameter maps to the --cache-from argument
            to the Docker CLI. If this parameter is `true`, only the final image will be pulled and passed to
            --cache-from if it is a CacheFrom object, the stages named therein will also be pulled and passed to
            --cache-from.
        :param Optional[Sequence[pulumi.Input[pulumi.Input[str]]]] extra_options: An optional catch-all list of arguments
            to provide extra CLI options to the docker build command.  For example `['--network', 'host']`.
        :param Optional[Mapping[str, str]] env: Environment variables to set on the invocation of `docker build`, for
         example to support `DOCKER_BUILDKIT=1 docker build`.
        :param Optional[pulumi.Input[str]] target: The target of the dockerfile to build
        """
        self.context = context
        self.dockerfile = dockerfile
        self.args = args
        self.cache_from = cache_from
        self.extra_options = extra_options
        self.env = env
        self.target = target


class Error(Exception):
    pass


class ResourceError(Error):
    def __init__(self, message: str, resource: Optional[pulumi.Resource], hide_stack: Optional[bool] = False):
        self.resource = resource
        self.hide_stack = hide_stack
        super().__init__(message)


async def use_docker_password_stdin(log_resource: pulumi.Resource):
    # Verify that 'docker' is on the PATH and get the client/server versions
    try:
        docker_version_str = await run_command_that_must_succeed(
            "docker", ["version", "-f", "{{json .}}"], log_resource)
        # IDEA: In the future we could warn here on out-of-date versions of Docker which may not support key
        # features we want to use.

        pulumi.log.debug(f'\'docker version\' => {docker_version_str}', log_resource)
    except Exception:
        raise ResourceError("No 'docker' command available on PATH: Please install to use container 'build' mode.",
                            log_resource)

    # Decide whether to use --password or --password-stdin based on the client version.
    try:
        version_data: any = json.loads(docker_version_str)
        client_version: str = version_data['Client']['Version']
        return LooseVersion(client_version) >= LooseVersion("17.07.0")
    except Exception as err:
        pulumi.log.info(f'Could not process Docker version ({err})', log_resource)
    return False


async def build_and_push_image(
    base_image_name: str,
    path_or_build: pulumi.Input[Union[str, DockerBuild]],
    repository_url: pulumi.Input[str],
    log_resource: pulumi.Resource,
    registry: Optional[Registry],
    skip_push: bool = False
) -> str:
    """
    build_and_push_image will build and push the Dockerfile and context from [pathOrBuild] into the
    requested docker repo [repository_url].  It returns the unique target image name for the image in
    the docker repository.  During preview this will build the image, and return the target image
    name, without pushing. During a normal update, it will do the same, as well as tag and push the
    image.
    """

    # Give an initial message indicating what we're about to do.  That way, if anything
    # takes a while, the user has an idea about what's going on.
    log_debug("Starting docker build and push...", log_resource)

    check_repository_url(repository_url)

    _, tag = get_image_name_and_tag(base_image_name)

    # login immediately if we're going to have to actually communicate with a remote registry.
    #
    # We know we have to login if:
    #
    #  1. We're doing an update.  In that case, we'll always want to login so we can push our
    #     images to the remote registry.
    #
    # 2. We're in preview or update and the build information contains 'cache from' information. In
    #    that case, we'll want want to pull from the registry and will need to login for that.

    pull_from_cache = not isinstance(path_or_build,
                                     str) and path_or_build and path_or_build.cache_from and repository_url is not None

    # If no `registry` info was passed in we simply assume docker is already
    # logged-in to the correct registry (or uses auto-login via credential helpers).
    if registry:
        if not pulumi.runtime.is_dry_run() or pull_from_cache:
            log_debug("Logging in to registry...", log_resource)
            await login_to_registry(registry, log_resource)

    # If the container specified a cache_from parameter, first set up the cached stages.
    cache_from = None
    if pull_from_cache:
        _cache_from_param = CacheFrom() if isinstance(path_or_build.cache_from, bool) else path_or_build.cache_from
        cache_from_param = _cache_from_param if _cache_from_param else CacheFrom()
        cache_from = pull_cache(base_image_name, cache_from_param, repository_url, log_resource)

    # Next, build the image.
    log_ephemeral(f"Building image '{ path_or_build if isinstance(path_or_build, str) else path_or_build.context or '.'}'...", log_resource)
    build_result = await build_image(base_image_name, path_or_build, log_resource, cache_from)
    image_id, stages = build_result.image_id, build_result.stages
    log_ephemeral("Image build succeeded.", log_resource)

    if image_id is None:
        raise Error("Internal error: docker build did not produce an imageId.")

    # Generate a name that uniquely will identify this built image.  This is similar in purpose to
    # the name@digest form that can be normally be retrieved from a docker repository.  However,
    # this tag doesn't require actually pushing the image, nor does it require communicating with
    # some external system, making it suitable for unique identification, even during preview.
    # This also means that if docker produces a new imageId, we'll get a new name here, ensuring that
    # resources (like docker.Image and cloud.Service) will be appropriately replaced.
    unique_tagged_image_name = create_tagged_image_name(repository_url, tag, image_id)

    # Use those to push the image.  Then just return the unique target name. as the final result
    # for our caller to use. Only push the image during an update, do not push during a preview.
    if not pulumi.runtime.is_dry_run() and not skip_push:
        # Push the final image first, then push the stage images to use for caching.
        log_ephemeral(f"Pushing image '{base_image_name}'...", log_resource)

        # First, push with both the optionally-requested-tag *and* imageId (which is guaranteed to
        # be defined).  By using the imageId we give the image a fully unique location that we can
        # successfully pull regardless of whatever else has happened at this repository_url.

        # Next, push only with the optionally-requested-tag.  Users of this API still want to get a
        # nice and simple url that they can reach this image at, without having the explicit imageId
        # hash added to it.  Note: this location is not guaranteed to be idempotent.  For example,
        # pushes on other machines might overwrite that location.
        await tag_and_push_image(base_image_name, repository_url, tag, image_id, log_resource=log_resource)
        await tag_and_push_image(base_image_name, repository_url, tag, image_id=None, log_resource=log_resource)

        for stage in stages:
            await tag_and_push_image(
                local_stage_image_name(base_image_name, stage),
                repository_url,
                stage,
                image_id=None,
                log_resource=log_resource
            )

        log_ephemeral("Image push succeeded.", log_resource)

    # If we got here, then building/pushing didn't throw any errors.  Update the status bar
    # indicating that things worked properly.  That way, the info bar isn't stuck showing the very
    # last thing printed by some subcommand we launched.
    log_debug("Successfully pushed to docker", log_resource)

    return unique_tagged_image_name


def log_ephemeral(message: str, log_resource: pulumi.Resource):
    try:
        pulumi.log.info(message, log_resource, stream_id=None, ephemeral=True)
    except TypeError:
        # that pulumi version does not support ephemeral
        pulumi.log.info(message, log_resource, stream_id=None)

def log_debug(message: str, log_resource: pulumi.Resource):
    try:
        pulumi.log.debug(message, log_resource, stream_id=None, ephemeral=True)
    except TypeError:
        # that pulumi version does not support ephemeral
        pulumi.log.info(message, log_resource, stream_id=None)

def check_repository_url(repository_url: str):
    _, tag = get_image_name_and_tag(repository_url)

    # We want to report an advisory error to users so that they don't accidentally include a 'tag'
    # in the repo url they supply.  i.e. their repo url can be:
    #
    #      docker.mycompany.com/namespace/myimage
    #
    # but should not be:
    #
    #      docker.mycompany.com/namespace/myimage:latest
    #
    # We could consider removing this check entirely.  However, it is likely valuable to catch
    # clear mistakes where a tag was included in a repo url inappropriately.
    #
    # However, since we do have the check, we need to ensure that we do allow the user to specify
    # a *port* on their repository that the are communicating with.  i.e. it's fine to have:
    #
    #      docker.mycompany.com:5000 or
    #      docker.mycompany.com:5000/namespace/myimage
    #
    # So check if this actually does look like a port, and don't report an error in that case.
    #
    # From: https:#www.w3.org/Addressing/URL/url-spec.txt
    #
    #      port        digits
    #
    # Regex = any number of digits, optionally followed by / and any remainder.

    if tag and not re.match(r'^\d+(/.*)?', tag):
        raise Error(f'[repository_url] should not contain a tag: {tag}')


def local_stage_image_name(image_name: str, stage: str):
    return f'{image_name}-{stage}'


def create_tagged_image_name(repository_url: str, tag: Optional[str], image_id: Optional[str]) -> str:
    pieces: List = []
    if tag:
        pieces.append(tag)

    if image_id:
        pieces.append(image_id)

    # Note: we don't do any validation that the tag is well formed, as per:
    # https://docs.docker.com/engine/reference/commandline/tag
    #
    # If there are any issues with it, we'll just let docker report the problem.
    full_tag = "-".join(pieces)
    return f'{repository_url}:{full_tag}' if full_tag else repository_url


async def pull_cache(
    image_name: str,
    cache_from,
    repo_url: str,
    log_resource: pulumi.Resource
) -> Optional[Sequence[str]]:
    # Ensure that we have a repository URL. If we don't, we won't be able to pull anything.
    if not repo_url:
        return None

    pulumi.log.debug(f'pulling cache for {image_name} from {repo_url}', log_resource)

    cache_from_images: List = []
    stages = (cache_from.stages if cache_from.stages else []) + [""]
    for stage in stages:
        tag = f':{stage}' if stage else ""
        image = f'{repo_url}{tag}'

        # Try to pull the existing image if it exists.  This may fail if the image does not exist.
        # That's fine, just move onto the next stage.  Also, pass along a flag saying that we
        # should print that error as a warning instead.  We don't want the update to succeed but
        # the user to then get a nasty "error:" message at the end.
        command_result = await run_command_that_can_fail(
            "docker", ["pull", image], log_resource,
            report_full_command_line=True, report_error_as_warning=True
        )
        if command_result.code:
            continue

        cache_from_images.append(image)

    return cache_from_images


class BuildResult:
    image_id: str
    stages: Sequence[str]

    def __init__(self, image_id, stages):
        self.image_id = image_id
        self.stages = stages


async def build_image(
    image_name: str,
    path_or_build: Union[str, DockerBuild],
    log_resource: pulumi.Resource,
    cache_from: Optional[Sequence[str]]
) -> BuildResult:
    if isinstance(path_or_build, str):
        build = DockerBuild(context=path_or_build)
    elif path_or_build:
        build = path_or_build
    else:
        raise ResourceError(f'Cannot build a container with an empty build specification', log_resource)

    # If the build context is missing, default it to the working directory.
    if not build.context:
        build.context = "."

    log_debug(
        f'Building container image \'{image_name}\': context={build.context}' +
        (f', dockerfile={build.dockerfile}' if build.dockerfile else "") +
        (f', args={json.dumps(build.args)}' if build.args else "") +
        (f', target={build.target}' if build.target else ""), log_resource)

    # If the container build specified build stages to cache, build each in turn.
    stages = []
    if build.cache_from and not isinstance(build.cache_from, bool) and build.cache_from.stages:
        for stage in build.cache_from.stages:
            await docker_build(
                local_stage_image_name(image_name, stage), build,
                cache_from=cache_from, log_resource=log_resource, target=stage)
            stages.append(stage)

    # Invoke Docker CLI commands to build.
    await docker_build(image_name, build, log_resource, cache_from)

    # Finally, inspect the image so we can return the SHA digest. Do not forward the output of this
    # command this to the CLI to show the user.
    inspect_result = await run_command_that_must_succeed(
        "docker", ["image", "inspect", "-f", "{{.Id}}", image_name], log_resource)
    if not inspect_result:
        raise ResourceError(
            f'No digest available for image {image_name}', log_resource)

    # From https:#docs.docker.com/registry/spec/api/#content-digests
    #
    # the image id will be a "algorithm:hex" pair.  We don't care about the algorithm part.  All we
    # want is the unique portion we can use elsewhere.  Since we are also going to place this in an
    # image tag, we also don't want the colon, as that's not legal there.  So simply grab the hex
    # portion after the colon and return that.

    image_id = inspect_result.strip()
    colon_index = image_id.rfind(":")
    image_id = image_id if colon_index < 0 else image_id[colon_index + 1:]

    return BuildResult(image_id, stages)


async def docker_build(
    image_name: str,
    build: DockerBuild,
    log_resource: pulumi.Resource,
    cache_from: Optional[Sequence[str]],
    target: Optional[str] = None
) -> str:
    # Prepare the build arguments.
    build_args: List[str] = ["build"]
    if build.dockerfile:
        build_args.extend(["-f", build.dockerfile])  # add a custom Dockerfile location.

    if build.args:
        for arg, build_arg in build.args.items():
            build_args.extend(["--build-arg", f'{arg}={build_arg}'])
    if build.target:
        build_args.extend(["--target", build.target])
    if cache_from:
        for image in cache_from:
            build_args.extend(["--cache-from", image])
    if build.extra_options:
        build_args.extend(build.extra_options)

    build_args.extend(["-t", image_name])  # tag the image with the chosen name.
    if target:
        build_args.extend(["--target", target])

    if build.context:
        build_args.append(build.context)  # push the docker build context onto the path.

    return await run_command_that_must_succeed("docker", build_args, log_resource, env=build.env)


class LoginResult:
    registry_name: str
    username: str
    done: asyncio.Task

    def __init__(self, registry_name: str, username: str, done: asyncio.Task):
        self.registry_name = registry_name
        self.username = username
        self.done = done


# Keep track of registries and users that have been logged in.  If we've already logged into that
# registry with that user, there's no need to do it again.
login_results: List[LoginResult] = []


async def login_to_registry(registry: Registry, log_resource: pulumi.Resource):
    registry_name = registry.registry
    username = registry.username
    password = registry.password

    # See if we've issued an outstanding requests to login into this registry.  If so, just
    # await the results of that login request.  Otherwise, create a new request and keep it
    # around so that future login requests will see it.
    result: LoginResult = None
    for existing in login_results:
        if existing.registry_name == registry_name and existing.username == username:
            log_debug(f'Reusing existing login for {username}@{registry_name}', log_resource)
            result = existing
            break

    if not result:
        # An existing login wasn't found; WARNING: we must not await anything between this check
        # for existing logins, and appending our new login attempt, otherwise an async interleaving
        # could sneak in, log into the same server, and yield a redundant login (which will error out).
        docker_password_stdin = use_docker_password_stdin(log_resource)

        # pass 'report_full_command_line: false' here so that if we fail to login we don't emit the
        # username/password in our logs.  Instead, we'll just say "'docker login' failed with code ..."
        coro: asyncio.Coroutine = None
        if docker_password_stdin:
            coro = run_command_that_must_succeed("docker", ["login", "-u", username, "--password-stdin", registry_name],
                                                 log_resource, report_full_command_line=False, stdin=password)
        else:
            coro = run_command_that_must_succeed("docker", ["login", "-u", username, "-p", password, registry_name],
                                                 log_resource, report_full_command_line=False)

        done = asyncio.create_task(coro)
        result = LoginResult(registry_name, username, done)
        login_results.append(result)

    await result.done


async def tag_and_push_image(
    image_name: str, repository_url: str,
    tag: Optional[str], image_id: Optional[str],
    log_resource: pulumi.Resource
):
    async def do_tag_and_push(target_name: str):
        await run_command_that_must_succeed("docker", ["tag", image_name, target_name], log_resource)
        await run_command_that_must_succeed("docker", ["push", target_name], log_resource)

    # Ensure we have a unique target name for this image, and tag and push to that unique target.
    await do_tag_and_push(create_tagged_image_name(repository_url, tag, image_id))

    # If the user provided a tag themselves (like "x/y:dev") then also tag and push directly to
    # that 'dev' tag.  This is not going to be a unique location, and future pushes will overwrite
    # this location.  However, that's ok as there's still the unique target we generated above.
    #
    # Note: don't need to do this if imageId was 'undefined' as the above line will have already
    # taken care of things for us.
    if tag is not None and image_id is not None:
        await do_tag_and_push(create_tagged_image_name(repository_url, tag, image_id=None))

    return


class CommandResult:
    code: int
    stdout: str

    def __init__(self, code, stdout):
        self.code = code
        self.stdout = stdout


def get_command_line_message(
    cmd: str, args: Sequence[str], report_full_command_line: bool, env: Optional[Mapping[str, str]] = None
):
    elements = []
    if env:
        elements.append(" ".join(map(lambda k: f'{k}={env[k]}', env.keys())))
    elements.append(cmd)

    argstr = " ".join(args) if report_full_command_line else args[0]
    elements.append(argstr)
    return f"'{' '.join(elements)}'"


def get_failure_message(
    cmd: str, args: Sequence[str], report_full_command_line: bool, code: int, env: Optional[Mapping[str, str]] = None
):
    return f'{get_command_line_message(cmd, args, report_full_command_line, env)} failed with exit code {code}'


# [report_full_command_line] is used to determine if the full command line should be reported
# when an error happens.  In general reporting the full command line is fine.  But it should be set
# to false if it might contain sensitive information (like a username/password)
async def run_command_that_must_succeed(
    cmd: str,
    args: Sequence[str],
    log_resource: pulumi.Resource,
    report_full_command_line: bool = True,
    stdin: Optional[str] = None,
    env: Optional[Mapping] = None
) -> str:
    command_result = await run_command_that_can_fail(
        cmd, args, log_resource, report_full_command_line, False, stdin, env)
    code, stdout = command_result.code, command_result.stdout

    if code != 0:
        # Fail the entire build and push.  This includes the full output of the command so that at
        # the end the user can review the full docker message about what the problem was.
        #
        # Note: a message about the command failing will have already been ephemerally reported to
        # the status column.
        raise ResourceError(
            f'{get_failure_message(cmd, args, report_full_command_line, code)}\n{stdout}', log_resource)

    return stdout


async def run_command_that_can_fail(
    cmd_name: str,
    args: Sequence[str],
    log_resource: pulumi.Resource,
    report_full_command_line: bool,
    report_error_as_warning: bool,
    stdin: Optional[str] = None,
    env: Optional[Mapping[str, str]] = None
) -> CommandResult:
    """
    Runs a CLI command in a child process, returning a future for the process's exit. Both stdout
    and stderr are redirected to process.stdout and process.stderr by default.

    If the [stdin] argument is defined, its contents are piped into stdin for the child process.

    [log_resource] is used to specify the resource to associate command output with. Stderr messages
    are always sent (since they may contain important information about something that's gone wrong).
    Stdout messages will be logged ephemerally to this resource.  This lets the user know there is
    progress, without having that dumped on them at the end.  If an error occurs though, the stdout
    content will be printed.
    """

    # Let the user ephemerally know the command we're going to execute.
    log_debug(f"Executing {get_command_line_message(cmd_name, args, report_full_command_line, env)}", log_resource)

    # Generate a unique stream-ID that we'll associate all the docker output with. This will allow
    # each spawned CLI command's output to associated with 'resource' and also streamed to the UI
    # in pieces so that it can be displayed live.  The stream-ID is so that the UI knows these
    # messages are all related and should be considered as one large message (just one that was
    # sent over in chunks).
    #
    # We use Math.random here in case our package is loaded multiple times in memory (i.e. because
    # different downstream dependencies depend on different versions of us).  By being random we
    # effectively make it completely unlikely that any two cli outputs could map to the same stream
    # id.
    #
    # Pick a reasonably distributed number between 0 and 2^30.  This will fit as an int32
    # which the grpc layer needs.
    stream_id = math.floor(random() * (1 << 30))

    if env is not None:
        env = os.environ.copy().update(env)

    process = await asyncio.create_subprocess_exec(
        cmd_name, *args, env=env,
        stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE, stdin=asyncio.subprocess.PIPE)

    # If stdin input is present, we need to ensure it's encoded into bytes.
    if isinstance(stdin, str):
        stdin = stdin.encode('utf-8')

    # We store the results from stdout in memory and will return them as a str.
    stdout_chunks: List[str] = []
    stderr_chunks: List[str] = []

    # A None value for process.returncode indicates that the process hasn't terminated yet
    while process.returncode is None:
        outs, errs = await process.communicate(input=stdin)
        if outs:
            # Report all stdout messages as ephemeral messages.  That way they show up in the
            # info bar as they're happening.  But they do not overwhelm the user as the end
            # of the run.
            for line in outs.splitlines():
                log_ephemeral(line, log_resource)

            stdout_chunks.append(outs.rstrip().decode('utf-8'))
        if errs:
            # We can't stream these stderr messages as we receive them because we don't knows at
            # this point because Docker uses stderr for both errors and warnings.  So, instead, we
            # just collect the messages, and wait for the process to end to decide how to report
            # them.
            stderr_chunks.append(errs.rstrip().decode('utf-8'))

    code = process.returncode

    # Collapse our stored stdout/stderr messages into single strings.
    stderr = ''.join(stderr_chunks)
    stdout = ''.join(stdout_chunks)

    # If we got any stderr messages, report them as an error/warning depending on the
    # result of the operation.
    if stderr:
        if code and not report_error_as_warning:
            # Command returned non-zero code.  Treat these stderr messages as an error.
            pulumi.log.error(stderr, log_resource, stream_id)
        else:
            # command succeeded.  These were just warning.
            pulumi.log.warn(stderr, log_resource, stream_id)

    # If the command failed report an ephemeral message indicating which command it was.
    # That way the user can immediately see something went wrong in the info bar.  The
    # caller (normally run_command_that_can_succeed) can choose to also report this
    # non-ephemerally.
    if code:
        log_ephemeral(get_failure_message(cmd_name, args, report_full_command_line, code), log_resource)

    return CommandResult(code, stdout)
