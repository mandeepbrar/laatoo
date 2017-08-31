let webpack = require('webpack')
let rimraf = require('rimraf')

let config = {}

let args = process.env.argv

config = require('./cfg/webpack.lib')

console.log("config", config)

let compiler = webpack(config)

rimraf('dist', function() {
  console.log("Removed directory dist")
})

compiler.run(function(err, stats) {
  console.log("Errors: ", err);
});
