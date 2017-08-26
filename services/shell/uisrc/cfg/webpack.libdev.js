const merge = require('webpack-merge');
const config = require('./webpack.devtpl');
const path = require('path');

module.exports = merge(config, {
  output: {
    libraryTarget: "umd",
    path: path.resolve(__dirname, '../dev/'),
    filename: 'scripts/index.js',
    publicPath: '/'
  }
});
