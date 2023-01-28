const path = require("path");

/** @type {import('next').NextConfig} */

const sassOptions = {
    includePaths: [path.join(__dirname, "styles")],
};

module.exports = {
    sassOptions,
    reactStrictMode: false,
    output: "standalone",
};
