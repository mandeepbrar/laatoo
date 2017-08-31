// Common Webpack configuration used by webpack.config.development and webpack.config.production

const path = require('path');
const webpack = require('webpack');
const autoprefixer = require('autoprefixer');

module.exports = {
  resolve: {
     modules: [
       path.join(__dirname, '../src'),
       path.join(__dirname, '../node_modules')
     ],
     alias: {
     },
     extensions: ['.js', '.jsx', '.json', '.scss', '.css', '.sass']
   },
   externals: {
       "babel-polyfill": "babel-polyfill",
       "react": "react",
       "react-dom":"react-dom",
       "redux-actions": "redux-actions",
       "react-redux":"react-redux",
       "tcomb-form":"tcomb-form",
       "md5":"md5",
       "redux-director":"redux-director",
       "director":"director",
       "redux":"redux",
       "redux-saga":"redux-saga",
       "react-tinymce-input":"react-tinymce-input",
       "react-player":"react-player",
       "react-dropzone":"react-dropzone",
       "react-bootstrap":"react-bootstrap",
       "react-pagify-preset-bootstrap":"react-pagify-preset-bootstrap",
       "segmentize":"segmentize",
       "react-pagify":"react-pagify"
   },
   module: {
    loaders: [
      // JavaScript / ES6
      {
        test: /\.(js|jsx)$/,
        include: path.join(__dirname, '../src'),
        loader: 'babel-loader'
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
