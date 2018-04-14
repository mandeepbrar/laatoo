// Common Webpack configuration used by webpack.config.development and webpack.config.production

const path = require('path');
const webpack = require('webpack');
//const autoprefixer = require('autoprefixer');
const dependencies = require('./DependenciesPlugin');

module.exports = {
  resolve: {
     modules: [
       './src',
       '/nodemodules/node_modules'
     ],
     alias: {
     },
     mainFields: ['browserify', 'browser', 'module', 'main'],
     extensions: ['.js', '.jsx', '.json', '.scss', '.css', '.sass']
  },
  resolveLoader: {
    modules: [
      '/nodemodules/node_modules'
    ]
  },
  module: {
    rules: [
      // JavaScript / ES6
      {
        test: /\.(js|jsx)$/,         // Match both .js and .jsx files
         exclude: /node_modules/,
         loader: "babel-loader"
      },
      // Images
      // Inline base64 URLs for <=8k images, direct URLs for the rest
      {
        test: /\.(png|jpg|jpeg|gif|svg)$/,
        loader: 'url-loader',
        query: {
          limit: 8192,
          name: '[name].[ext]?[hash]',
          outputPath: '/dist/images',
          publicPath: 'images'
        }
      },
      // Fonts
      {
        test: /\.(woff|woff2|ttf|eot)(\?v=\d+\.\d+\.\d+)?$/,
        loader: 'url-loader',
        query: {
          limit: 8192,
          name: '[name].[ext]?[hash]',
          outputPath: '/dist/fonts',
          publicPath: 'fonts'
        }
      }
    ]
  }
};
