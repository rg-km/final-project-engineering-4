const { alias } = require('react-app-rewire-alias');

module.exports = function override(config) {
  alias({
    '@images': './src/images',
    '@config': './src/config',
    '@styles': './src/styles',
    '@routes': './src/routes',
    '@services': './src/services',
    '@store': './src/store',
    '@utils': './src/utils',
    '@tests': './src/tests',
  })(config);

  return config;
};
