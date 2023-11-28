"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
exports.Container = void 0;
const pulumi = require("@pulumi/pulumi");
const utilities = require("./utilities");
/**
 * <!-- Bug: Type and Name are switched -->
 * Manages the lifecycle of a Docker container.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * // Find the latest Ubuntu precise image.
 * const ubuntuRemoteImage = new docker.RemoteImage("ubuntuRemoteImage", {name: "ubuntu:precise"});
 * // Start a container
 * const ubuntuContainer = new docker.Container("ubuntuContainer", {image: ubuntuRemoteImage.imageId});
 * ```
 *
 * ## Import
 *
 * ### Example Assuming you created a `container` as follows #!/bin/bash docker run --name foo -p8080:80 -d nginx
 *
 * prints the container ID
 *
 * 9a550c0f0163d39d77222d3efd58701b625d47676c25c686c95b5b92d1cba6fd you provide the definition for the resource as follows terraform resource "docker_container" "foo" {
 *
 *  name
 *
 * = "foo"
 *
 *  image = "nginx"
 *
 *  ports {
 *
 *  internal = "80"
 *
 *  external = "8080"
 *
 *  } } then the import command is as follows #!/bin/bash
 *
 * ```sh
 *  $ pulumi import docker:index/container:Container foo 9a550c0f0163d39d77222d3efd58701b625d47676c25c686c95b5b92d1cba6fd
 * ```
 */
class Container extends pulumi.CustomResource {
    /**
     * Get an existing Container resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name, id, state, opts) {
        return new Container(name, state, Object.assign(Object.assign({}, opts), { id: id }));
    }
    /**
     * Returns true if the given object is an instance of Container.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj) {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Container.__pulumiType;
    }
    constructor(name, argsOrState, opts) {
        let resourceInputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState;
            resourceInputs["attach"] = state ? state.attach : undefined;
            resourceInputs["bridge"] = state ? state.bridge : undefined;
            resourceInputs["capabilities"] = state ? state.capabilities : undefined;
            resourceInputs["cgroupnsMode"] = state ? state.cgroupnsMode : undefined;
            resourceInputs["command"] = state ? state.command : undefined;
            resourceInputs["containerLogs"] = state ? state.containerLogs : undefined;
            resourceInputs["containerReadRefreshTimeoutMilliseconds"] = state ? state.containerReadRefreshTimeoutMilliseconds : undefined;
            resourceInputs["cpuSet"] = state ? state.cpuSet : undefined;
            resourceInputs["cpuShares"] = state ? state.cpuShares : undefined;
            resourceInputs["destroyGraceSeconds"] = state ? state.destroyGraceSeconds : undefined;
            resourceInputs["devices"] = state ? state.devices : undefined;
            resourceInputs["dns"] = state ? state.dns : undefined;
            resourceInputs["dnsOpts"] = state ? state.dnsOpts : undefined;
            resourceInputs["dnsSearches"] = state ? state.dnsSearches : undefined;
            resourceInputs["domainname"] = state ? state.domainname : undefined;
            resourceInputs["entrypoints"] = state ? state.entrypoints : undefined;
            resourceInputs["envs"] = state ? state.envs : undefined;
            resourceInputs["exitCode"] = state ? state.exitCode : undefined;
            resourceInputs["gpus"] = state ? state.gpus : undefined;
            resourceInputs["groupAdds"] = state ? state.groupAdds : undefined;
            resourceInputs["healthcheck"] = state ? state.healthcheck : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["hosts"] = state ? state.hosts : undefined;
            resourceInputs["image"] = state ? state.image : undefined;
            resourceInputs["init"] = state ? state.init : undefined;
            resourceInputs["ipcMode"] = state ? state.ipcMode : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["logDriver"] = state ? state.logDriver : undefined;
            resourceInputs["logOpts"] = state ? state.logOpts : undefined;
            resourceInputs["logs"] = state ? state.logs : undefined;
            resourceInputs["maxRetryCount"] = state ? state.maxRetryCount : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["memorySwap"] = state ? state.memorySwap : undefined;
            resourceInputs["mounts"] = state ? state.mounts : undefined;
            resourceInputs["mustRun"] = state ? state.mustRun : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkDatas"] = state ? state.networkDatas : undefined;
            resourceInputs["networkMode"] = state ? state.networkMode : undefined;
            resourceInputs["networksAdvanced"] = state ? state.networksAdvanced : undefined;
            resourceInputs["pidMode"] = state ? state.pidMode : undefined;
            resourceInputs["ports"] = state ? state.ports : undefined;
            resourceInputs["privileged"] = state ? state.privileged : undefined;
            resourceInputs["publishAllPorts"] = state ? state.publishAllPorts : undefined;
            resourceInputs["readOnly"] = state ? state.readOnly : undefined;
            resourceInputs["removeVolumes"] = state ? state.removeVolumes : undefined;
            resourceInputs["restart"] = state ? state.restart : undefined;
            resourceInputs["rm"] = state ? state.rm : undefined;
            resourceInputs["runtime"] = state ? state.runtime : undefined;
            resourceInputs["securityOpts"] = state ? state.securityOpts : undefined;
            resourceInputs["shmSize"] = state ? state.shmSize : undefined;
            resourceInputs["start"] = state ? state.start : undefined;
            resourceInputs["stdinOpen"] = state ? state.stdinOpen : undefined;
            resourceInputs["stopSignal"] = state ? state.stopSignal : undefined;
            resourceInputs["stopTimeout"] = state ? state.stopTimeout : undefined;
            resourceInputs["storageOpts"] = state ? state.storageOpts : undefined;
            resourceInputs["sysctls"] = state ? state.sysctls : undefined;
            resourceInputs["tmpfs"] = state ? state.tmpfs : undefined;
            resourceInputs["tty"] = state ? state.tty : undefined;
            resourceInputs["ulimits"] = state ? state.ulimits : undefined;
            resourceInputs["uploads"] = state ? state.uploads : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["usernsMode"] = state ? state.usernsMode : undefined;
            resourceInputs["volumes"] = state ? state.volumes : undefined;
            resourceInputs["wait"] = state ? state.wait : undefined;
            resourceInputs["waitTimeout"] = state ? state.waitTimeout : undefined;
            resourceInputs["workingDir"] = state ? state.workingDir : undefined;
        }
        else {
            const args = argsOrState;
            if ((!args || args.image === undefined) && !opts.urn) {
                throw new Error("Missing required property 'image'");
            }
            resourceInputs["attach"] = args ? args.attach : undefined;
            resourceInputs["capabilities"] = args ? args.capabilities : undefined;
            resourceInputs["cgroupnsMode"] = args ? args.cgroupnsMode : undefined;
            resourceInputs["command"] = args ? args.command : undefined;
            resourceInputs["containerReadRefreshTimeoutMilliseconds"] = args ? args.containerReadRefreshTimeoutMilliseconds : undefined;
            resourceInputs["cpuSet"] = args ? args.cpuSet : undefined;
            resourceInputs["cpuShares"] = args ? args.cpuShares : undefined;
            resourceInputs["destroyGraceSeconds"] = args ? args.destroyGraceSeconds : undefined;
            resourceInputs["devices"] = args ? args.devices : undefined;
            resourceInputs["dns"] = args ? args.dns : undefined;
            resourceInputs["dnsOpts"] = args ? args.dnsOpts : undefined;
            resourceInputs["dnsSearches"] = args ? args.dnsSearches : undefined;
            resourceInputs["domainname"] = args ? args.domainname : undefined;
            resourceInputs["entrypoints"] = args ? args.entrypoints : undefined;
            resourceInputs["envs"] = args ? args.envs : undefined;
            resourceInputs["gpus"] = args ? args.gpus : undefined;
            resourceInputs["groupAdds"] = args ? args.groupAdds : undefined;
            resourceInputs["healthcheck"] = args ? args.healthcheck : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["hosts"] = args ? args.hosts : undefined;
            resourceInputs["image"] = args ? args.image : undefined;
            resourceInputs["init"] = args ? args.init : undefined;
            resourceInputs["ipcMode"] = args ? args.ipcMode : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["logDriver"] = args ? args.logDriver : undefined;
            resourceInputs["logOpts"] = args ? args.logOpts : undefined;
            resourceInputs["logs"] = args ? args.logs : undefined;
            resourceInputs["maxRetryCount"] = args ? args.maxRetryCount : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["memorySwap"] = args ? args.memorySwap : undefined;
            resourceInputs["mounts"] = args ? args.mounts : undefined;
            resourceInputs["mustRun"] = args ? args.mustRun : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkMode"] = args ? args.networkMode : undefined;
            resourceInputs["networksAdvanced"] = args ? args.networksAdvanced : undefined;
            resourceInputs["pidMode"] = args ? args.pidMode : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
            resourceInputs["privileged"] = args ? args.privileged : undefined;
            resourceInputs["publishAllPorts"] = args ? args.publishAllPorts : undefined;
            resourceInputs["readOnly"] = args ? args.readOnly : undefined;
            resourceInputs["removeVolumes"] = args ? args.removeVolumes : undefined;
            resourceInputs["restart"] = args ? args.restart : undefined;
            resourceInputs["rm"] = args ? args.rm : undefined;
            resourceInputs["runtime"] = args ? args.runtime : undefined;
            resourceInputs["securityOpts"] = args ? args.securityOpts : undefined;
            resourceInputs["shmSize"] = args ? args.shmSize : undefined;
            resourceInputs["start"] = args ? args.start : undefined;
            resourceInputs["stdinOpen"] = args ? args.stdinOpen : undefined;
            resourceInputs["stopSignal"] = args ? args.stopSignal : undefined;
            resourceInputs["stopTimeout"] = args ? args.stopTimeout : undefined;
            resourceInputs["storageOpts"] = args ? args.storageOpts : undefined;
            resourceInputs["sysctls"] = args ? args.sysctls : undefined;
            resourceInputs["tmpfs"] = args ? args.tmpfs : undefined;
            resourceInputs["tty"] = args ? args.tty : undefined;
            resourceInputs["ulimits"] = args ? args.ulimits : undefined;
            resourceInputs["uploads"] = args ? args.uploads : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["usernsMode"] = args ? args.usernsMode : undefined;
            resourceInputs["volumes"] = args ? args.volumes : undefined;
            resourceInputs["wait"] = args ? args.wait : undefined;
            resourceInputs["waitTimeout"] = args ? args.waitTimeout : undefined;
            resourceInputs["workingDir"] = args ? args.workingDir : undefined;
            resourceInputs["bridge"] = undefined /*out*/;
            resourceInputs["containerLogs"] = undefined /*out*/;
            resourceInputs["exitCode"] = undefined /*out*/;
            resourceInputs["networkDatas"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Container.__pulumiType, name, resourceInputs, opts);
    }
}
exports.Container = Container;
/** @internal */
Container.__pulumiType = 'docker:index/container:Container';
//# sourceMappingURL=container.js.map