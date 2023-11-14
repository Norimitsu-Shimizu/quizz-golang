import axios from "axios";
// import { getPublicConfig } from "@/utils/config";

// const config = getPublicConfig();

export const client = axios.create({
  baseURL: "http://localhost:8090",
});

const FetchClient = () => {
  client.interceptors.request.use(function (config) {
    const accessToken = localStorage.getItem("_access_token");
    const client = localStorage.getItem("_client");

    if (accessToken && client) {
      config.headers["access-token"] = accessToken;
      config.headers["client"] = client;
    }
    return config;
  });

  return client;
};

export default FetchClient;
