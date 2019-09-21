const merge = require('webpack-merge');
const config = require('./webpack.disttpl');
const path = require('path');
const webpack = require('webpack');
var sprintf = require('sprintf-js').sprintf
var fs = require('fs-extra')
var {log} = require('../utils')

module.exports = function(env) {
  log(config.resolve.extensions);
  conf=[]
  let conf1 = merge(config, {
    context: env.uifolder,
    entry: {
      index: 'index'
    },
    output: {
      library: env.library,
      libraryTarget: "amd",
      filename: 'scripts/index.js',
      path: env.uiBuildFolder,
      publicPath: '/'
    }
  })
  if (env.externals) {
    conf1 = merge(conf1, {externals: env.externals})
  }
  conf.push(conf1);
  if (env.dependencies) {
    var tmpfile = '/build/tmp.js';
    fs.removeSync(tmpfile)
    let str = ""
    var rules=[]
    var initStr = ""
    log("dependencies", env.dependencies)
    Object.keys(env.dependencies).forEach(function(k, index){
      let modName = "m"+index
      str= sprintf("%svar %s=require('%s');", str, modName, k)
      initStr = sprintf("%s; console.log('Initializing %s', %s); f('%s',[], function(){return %s;});", initStr, k, modName, k, modName)
      //str= str + " console.log('found module "+ k + " "+modName+"', " +modName+" );"
    });
    var initFunc = "function Initialize(app, settings, f){"+initStr+";}; export {Initialize}"
    fs.writeFileSync(tmpfile, str+initFunc);
    let conf2 = merge(config, {
      context: env.uifolder,
      entry: {
        index: tmpfile
      },
      externals: env.externals,
      output: {
        library: env.library+"_vendor",
        libraryTarget: "amd",
        filename: 'scripts/vendor.js',
        path: env.uiBuildFolder,
        publicPath: '/'
      }
    })
    conf.push(conf2);
  }
  return conf
};
/*module: {
  rules: [{
          test: require.resolve('jquery'),
          use: [{
              loader: 'expose-loader',
              options: 'jQuery'
          },{
              loader: 'expose-loader',
              options: '$'
          }]
      }]
}
*/
