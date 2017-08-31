const merge = require('webpack-merge');
const config = require('./webpack.disttpl');
const path = require('path');

module.exports = merge(config, {
  output: {
    libraryTarget: "umd",
    filename: 'scripts/index.js',
    path: path.resolve(__dirname, '../dist/'),
    publicPath: '/'
  }
});
