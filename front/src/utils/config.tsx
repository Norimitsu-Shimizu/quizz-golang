import getNextConfig from "next/config";
import { EnvConfig } from "../../env/types/env";

export const getPublicConfig = (): EnvConfig => {
  const { publicRuntimeConfig } = getNextConfig();
  return publicRuntimeConfig;
};

export const getServerConfig = () => {
  const { serverRuntimeConfig } = getNextConfig();
  return serverRuntimeConfig;
};
