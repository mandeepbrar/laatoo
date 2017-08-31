const merge = require('webpack-merge');
const config = require('./webpack.disttpl');
const path = require('path');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = merge(config, {
  output: {
    filename: 'scripts/index.js',
    path: path.resolve(__dirname, '../dist/'),
    publicPath: '/'
  },
  plugins:[
    new HtmlWebpackPlugin({
      template: path.join(__dirname, '../src/index.tpl.html')
    }),
    new CopyWebpackPlugin([
      {
        from: path.join(__dirname, '../src/images'),
        to: 'images'
      }
    ]),
  ]
});
