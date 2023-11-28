export { ContainerArgs, ContainerState } from "./container";
export type Container = import("./container").Container;
export declare const Container: typeof import("./container").Container;
export { GetLogsArgs, GetLogsResult, GetLogsOutputArgs } from "./getLogs";
export declare const getLogs: typeof import("./getLogs").getLogs;
export declare const getLogsOutput: typeof import("./getLogs").getLogsOutput;
export { GetNetworkArgs, GetNetworkResult, GetNetworkOutputArgs } from "./getNetwork";
export declare const getNetwork: typeof import("./getNetwork").getNetwork;
export declare const getNetworkOutput: typeof import("./getNetwork").getNetworkOutput;
export { GetPluginArgs, GetPluginResult, GetPluginOutputArgs } from "./getPlugin";
export declare const getPlugin: typeof import("./getPlugin").getPlugin;
export declare const getPluginOutput: typeof import("./getPlugin").getPluginOutput;
export { GetRegistryImageArgs, GetRegistryImageResult, GetRegistryImageOutputArgs } from "./getRegistryImage";
export declare const getRegistryImage: typeof import("./getRegistryImage").getRegistryImage;
export declare const getRegistryImageOutput: typeof import("./getRegistryImage").getRegistryImageOutput;
export { GetRemoteImageArgs, GetRemoteImageResult, GetRemoteImageOutputArgs } from "./getRemoteImage";
export declare const getRemoteImage: typeof import("./getRemoteImage").getRemoteImage;
export declare const getRemoteImageOutput: typeof import("./getRemoteImage").getRemoteImageOutput;
export { ImageArgs } from "./image";
export type Image = import("./image").Image;
export declare const Image: typeof import("./image").Image;
export { NetworkArgs, NetworkState } from "./network";
export type Network = import("./network").Network;
export declare const Network: typeof import("./network").Network;
export { PluginArgs, PluginState } from "./plugin";
export type Plugin = import("./plugin").Plugin;
export declare const Plugin: typeof import("./plugin").Plugin;
export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export declare const Provider: typeof import("./provider").Provider;
export { RegistryImageArgs, RegistryImageState } from "./registryImage";
export type RegistryImage = import("./registryImage").RegistryImage;
export declare const RegistryImage: typeof import("./registryImage").RegistryImage;
export { RemoteImageArgs, RemoteImageState } from "./remoteImage";
export type RemoteImage = import("./remoteImage").RemoteImage;
export declare const RemoteImage: typeof import("./remoteImage").RemoteImage;
export { SecretArgs, SecretState } from "./secret";
export type Secret = import("./secret").Secret;
export declare const Secret: typeof import("./secret").Secret;
export { ServiceArgs, ServiceState } from "./service";
export type Service = import("./service").Service;
export declare const Service: typeof import("./service").Service;
export { ServiceConfigArgs, ServiceConfigState } from "./serviceConfig";
export type ServiceConfig = import("./serviceConfig").ServiceConfig;
export declare const ServiceConfig: typeof import("./serviceConfig").ServiceConfig;
export { TagArgs, TagState } from "./tag";
export type Tag = import("./tag").Tag;
export declare const Tag: typeof import("./tag").Tag;
export { VolumeArgs, VolumeState } from "./volume";
export type Volume = import("./volume").Volume;
export declare const Volume: typeof import("./volume").Volume;
export * from "./types/enums";
import * as config from "./config";
import * as types from "./types";
export { config, types, };
