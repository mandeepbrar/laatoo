// Common Webpack configuration used by webpack.config.development and webpack.config.production

const path = require('path');
const webpack = require('webpack');

const config = require('/nodemodules/dll.json');

module.exports = {
  entry: {
    vendor: config.packages
  },
  output: {
      filename: 'vendor.js',
      path: '/nodemodules/dll'
  },
  resolve: {
     modules: [
       '/nodemodules/node_modules'
     ],
     alias: {
     },
     extensions: ['.js', '.jsx', '.json', '.scss', '.css', '.sass']
   },
  resolveLoader: {
    modules: [
      '/nodemodules/node_modules'
    ]
  },
  plugins:[
    new webpack.DllPlugin({
        name: 'vendor',
        path: '/nodemodules/dll/vendor-manifest.json',
      }),
    new webpack.optimize.UglifyJsPlugin({
      compress: {
        warnings: false
      },
      output: {
        comments: false
      },
      sourceMap: false
    })      
  ],
  module: {
    loaders: [
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
          name: 'images/[name].[ext]?[hash]'
        }
      },
      // Fonts
      {
        test: /\.(woff|woff2|ttf|eot)(\?v=\d+\.\d+\.\d+)?$/,
        loader: 'url-loader',
        query: {
          limit: 8192,
          name: 'fonts/[name].[ext]?[hash]'
        }
      }
    ]
  }
};
