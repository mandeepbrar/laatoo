'use strict';

let path = require('path');
let webpack = require('webpack');

let baseConfig = require('./base');
let defaultSettings = require('./defaults');

// Add needed plugins here
let BowerWebpackPlugin = require('bower-webpack-plugin');

let config = Object.assign({}, baseConfig, {
  entry: {
    laatoo: path.join(__dirname, '../src/laatoo'),
    //vendor:["react", "react-dom","react-redux", "redux", "core-js", "normalize.css"]
  },
  output: {
      // export itself to a global var
      // name of the global var: "Foo"
      library: "laatoo",
      path: path.join(__dirname, '/../dist/assets'),
      libraryTarget: "commonjs2",
      filename: 'index.js',
      publicPath: `.${defaultSettings.publicPath}`
  },
  externals: {
      // require("jquery") is external and available
      //  on the global var jQuery
      "babel-polyfill": "babel-polyfill",
      "react": "react",
      "react-dom":"react-dom",
      "core-js": "core-js",
      "normalize.css": "normalize.css",
      "redux-actions": "redux-actions",
      "react-redux":"react-redux",
      "tcomb-form":"tcomb-form",
      "md5":"md5",
      "redux-director":"redux-director",
      "director":"director",
      "redux":"redux",
      "redux-saga":"redux-saga",
      "core-js":"core-js",
      "react-tinymce-input":"react-tinymce-input",
      "react-player":"react-player",
      "react-dropzone":"react-dropzone",
      "react-bootstrap":"react-bootstrap",
      "normalize.css":"normalize.css"
  },
  cache: false,
  devtool: 'sourcemap',
  plugins: [
    new webpack.optimize.DedupePlugin(),
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': '"production"'
    }),
    new BowerWebpackPlugin({
      searchResolveModulesDirectories: false
    }),
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

config.module.loaders.push({
  test: /\.(js|jsx)$/,
  loader: 'strip-loader?strip[]=console.log'
});
module.exports = config;
