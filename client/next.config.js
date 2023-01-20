const path = require("path");

/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
};

const sassOptions = {
    includePaths: [path.join(__dirname, "styles")],
};

module.exports = {
    nextConfig,
    sassOptions,
};
