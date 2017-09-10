const path = require('path');
const merge = require('webpack-merge');
const webpack = require('webpack');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const config = require('./webpack.base');

const GLOBALS = {
  'process.env': {
    'NODE_ENV': JSON.stringify('production')
  },
  __DEV__: JSON.stringify(JSON.parse(process.env.DEBUG || 'false'))
};

let cssaliases = {
  "common": "/nodemodules/node_modules/reactwebcommon/files/app"
}
let cssoptions = {
  includePaths: [
    path.resolve("/nodemodules/node_modules/")
  ],
  alias: cssaliases
}
module.exports = merge(config, {
  devtool: 'cheap-module-source-map',
  plugins: [
    // Avoid publishing files when compilation fails
    new webpack.NoErrorsPlugin(),
    new webpack.DefinePlugin(GLOBALS),
    new webpack.optimize.UglifyJsPlugin({
      compress: {
        warnings: false
      },
      output: {
        comments: false
      },
      sourceMap: false
    }),
    new webpack.LoaderOptionsPlugin({
      minimize: true,
      debug: false
    }),
    new ExtractTextPlugin({
      filename: 'dist/css/app.css',
      allChunks: true
    })
  ],
  module: {
    noParse: /\.min\.js$/,
    loaders: [
      // Sass
      {
        test: /\.scss$/,
        loader: ExtractTextPlugin.extract({
          fallback: 'style-loader',
          use: [
            { loader: 'css-loader', options: merge({ sourceMap: true }, cssoptions) },
            { loader: 'resolve-url-loader'},
            { loader: 'sass-loader', options: merge({ outputStyle: 'compressed'}, cssoptions ) }
          ]
        })
      },
      {
        test: /\.sass$/,
        loader: ExtractTextPlugin.extract({
          fallback: 'style-loader',
          use: [
            { loader: 'css-loader', options: merge( { sourceMap: true }, cssoptions) },
            { loader: 'resolve-url-loader'},
            { loader: 'sass-loader', options: merge( { outputStyle: 'compressed' }, cssoptions) }
          ]
        })
      },
      {
        test: /\.css$/,
        loader: ExtractTextPlugin.extract({
          fallback: 'style-loader',
          use: [
             { loader: 'css-loader', options: cssoptions },
             { loader: 'resolve-url-loader'}
          ]
        })
      }
    ]
  },
});
