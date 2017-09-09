var PrintChunksPlugin = function() {};
PrintChunksPlugin.prototype.apply = function(compiler) {
    compiler.plugin('compilation', function(compilation, params) {
        compilation.plugin('after-optimize-chunk-assets', function(chunks) {
            Object.keys(chunks).forEach(function(k){
              console.log("=====================")
              console.log(chunks[k])
            });
        });
    });
};
module.exports = {
  PrintChunksPlugin
};
