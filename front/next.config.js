/** @type {import('next').NextConfig} */
const envConfig = require(`./env/${process.env.APP_ENV || "development"}.ts`);

const nextConfig = {
  reactStrictMode: true,

  publicRuntimeConfig: {
    ...envConfig,
  },

  images: {
    remotePatterns: [
      {
        protocol: 'http',
        hostname: 'localhost',
      },
    ],
  },

}

module.exports = nextConfig
