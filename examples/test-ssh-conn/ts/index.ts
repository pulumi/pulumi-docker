import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";
import * as digitalocean from "@pulumi/digitalocean";

const remoteHost = new digitalocean.Droplet("docker-host-test", {
    image: "docker-20-04",
    region: "fra1",
    size: "c-8",
    sshKeys: ["cb:dd:70:4d:49:2f:86:eb:fd:bb:e4:8b:04:fc:b0:cb"],
})

function sleep(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
async function addSleep(addr: string, ms: number) {
    console.log("Sleeping...")
    await sleep(ms)
    console.log("Done Sleeping")
    return addr
}

let ip = remoteHost.ipv4Address.apply(ipv4Address => addSleep(ipv4Address, 20000))

export const ipOutput = ip

const provider = new docker.Provider("docker-provider", {
    host: pulumi.interpolate`ssh://root@${ip}`,
    sshOpts: [
        "-i", "/home/runner/.ssh",
        "-o", "StrictHostKeyChecking=no",
        "-o", "UserKnownHostsFile=/dev/null"
    ],
});

const remoteImage = new docker.RemoteImage("image", {
    name: "nginx"
}, {
    provider: provider
});

const container = new docker.Container("container", {
    image: remoteImage.imageId
}, {
    provider: provider
});
