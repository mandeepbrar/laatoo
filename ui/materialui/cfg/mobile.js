'use strict';

let path = require('path');
let webpack = require('webpack');

let baseConfig = require('./base');
let defaultSettings = require('./defaults');
var nodeExternals = require('webpack-node-externals');

// Add needed plugins here
let BowerWebpackPlugin = require('bower-webpack-plugin');
let config = Object.assign({}, baseConfig, {
  entry: {
    laatoomobile: path.join(__dirname, '../src/mobile'),
    //vendor:["react", "react-dom","react-redux", "redux", "core-js", "normalize.css"]
  },
  output: {
      // export itself to a global var
      // name of the global var: "Foo"
      library: "laatoo",
      path: path.join(__dirname, '/../dist/assets'),
      libraryTarget: "umd",
      filename: 'mobile.js',
      publicPath: `.${defaultSettings.publicPath}`
  },
  externals: {
      "babel-polyfill": "babel-polyfill",
      "react": "react",
      "core-js": "core-js",
      "normalize.css": "normalize.css",
      "redux-actions": "redux-actions",
      "react-redux":"react-redux",
      "md5":"md5",
      "redux":"redux",
      "redux-saga":"redux-saga",
      "core-js":"core-js",
      "segmentize":"segmentize",
      "normalize.css":"normalize.css"
  },
  optional: true,
  cache: false,
  devtool: 'sourcemap',
  plugins: [
    new webpack.optimize.DedupePlugin(),
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': '"production"'
    }),
    new BowerWebpackPlugin(),
//    new webpack.optimize.CommonsChunkPlugin("vendor", "react.bundle.js"),
    new webpack.optimize.UglifyJsPlugin(),
    new webpack.optimize.OccurenceOrderPlugin(),
    new webpack.optimize.AggressiveMergingPlugin(),
    new webpack.NoErrorsPlugin()
  ],
  module: defaultSettings.getDefaultModules()
});

// Add needed loaders to the defaults here
config.module.loaders.push({
  test: /\.(js|jsx)$/,
  loader: 'babel',
  include: [].concat(
    config.additionalPaths,
    [ path.join(__dirname, '/../src') ]
  )
});
/*
config.module.loaders.push({
  test: /\.(js|jsx)$/,
  loader: 'strip-loader?strip[]=console.log'
});*/
module.exports = config;
