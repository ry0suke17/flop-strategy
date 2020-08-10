const withSourceMaps = require('@zeit/next-source-maps');
const dotenv = require('dotenv');
const path = require('path');

dotenv.config();

module.exports = withSourceMaps({
  target: 'serverless',
  env: {
    BASE_PATH: process.env.BASE_PATH || 'http://localhost:8080',
  },
  webpack: (config, options) => {
    config.resolve.alias['@fls-lib'] = path.join(__dirname, 'lib');
    config.resolve.alias['@fls-api-clinet'] = path.join(
      __dirname,
      'fls-api-clinet',
    );
    config.resolve.alias['@fls-pages'] = path.join(__dirname, 'pages');
    config.resolve.alias['@fls-components'] = path.join(
      __dirname,
      'components',
    );

    return config;
  },
});
