const merge = require('webpack-merge');
const config = require('./webpack.disttpl');
const path = require('path');
const webpack = require('webpack');
var fs = require('fs-extra')

module.exports = function(env) {
  conf=[]
  let conf1 = merge(config, {
    context: env.uifolder,
    entry: {
      index: 'index'
    },
    output: {
      library: env.library,
      libraryTarget: "amd",
      filename: 'dist/scripts/index.js',
      path: env.uifolder,
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
    Object.keys(env.dependencies).forEach(function(k, index){
      let modName = "m"+index
      str= str + "var "+modName+"=require('" +k+"');"
      initStr = initStr + "f('"+k+"',[],function(){return "+modName+";});"
    });
    var initFunc = "function Initialize(app, settings, f){"+initStr+";}; export {Initialize}"
    fs.writeFileSync(tmpfile, str+initFunc);
    let conf2 = merge(config, {
      context: env.uifolder,
      entry: {
        index: tmpfile
      },
      output: {
        library: env.library+"_vendor",
        libraryTarget: "amd",
        filename: 'dist/scripts/vendor.js',
        path: env.uifolder,
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
