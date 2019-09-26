// Common Webpack configuration used by webpack.config.development and webpack.config.production

const path = require('path');
const webpack = require('webpack');
//const autoprefixer = require('autoprefixer');
const dependencies = require('./DependenciesPlugin');
require("@babel/plugin-transform-runtime");


const presets = [
  ["@babel/preset-env", { "modules": false }],
  "@babel/preset-react"
];
const plugins = [
  "@babel/plugin-transform-runtime",
  "@babel/plugin-transform-flow-strip-types",
  "@babel/plugin-proposal-class-properties"
];

module.exports = {
  resolve: {
     modules: [
       './src',
       './js',
       '/nodemodules/node_modules'
     ],
     alias: {
     },
     mainFields: ['browserify', 'browser', 'module', 'main', 'style'],
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
        loader: "babel-loader",
        options: {
          configFile: "/nodemodules/babel.config.js"
          /*config: presets,
          plugins: plugins*/
        }
      },
      // Images
      // Inline base64 URLs for <=8k images, direct URLs for the rest
      {
        test: /\.(png|jpg|jpeg|gif|svg)$/,
        loader: 'url-loader',
        query: {
          limit: 8192,
          name: '[name].[ext]?[hash]',
          outputPath: '/images',
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
          outputPath: '/fonts',
          publicPath: 'fonts'
        }
      }
    ]
  }
};
