const merge = require('webpack-merge');
const config = require('./webpack.devtpl');
const path = require('path');

module.exports = function(env) {
  return merge(config, {
    output: {
      library: env.library,
      libraryTarget: "amd",
      filename: 'scripts/index.js',
      path: '/plugin/dev',
      publicPath: '/'
    }
  })
};
