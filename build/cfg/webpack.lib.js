const merge = require('webpack-merge');
const config = require('./webpack.disttpl');
const path = require('path');

module.exports = function(env) {
  let conf = merge(config, {
    context: env.uifolder,
    output: {
      library: env.library,
      libraryTarget: "amd",
      filename: 'dist/scripts/index.js',
      path: env.uifolder,
      publicPath: '/'
    }
  })
  if (env.externals) {
    conf = merge(conf, {externals: env.externals})
  }
  return conf
};
