const merge = require('webpack-merge');
const config = require('./webpack.devtpl');

module.exports = merge(config, {
  output: {
    filename: 'scripts/index.js',
    publicPath: '/'
  }
});
