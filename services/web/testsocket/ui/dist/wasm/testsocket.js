
(function() {
    var wasm;
    const __exports = {};
    
    
    const __wbg_log_f74b4eaeae4e758a_target = console.log;
    
    let cachedDecoder = new TextDecoder('utf-8');
    
    let cachegetUint8Memory = null;
    function getUint8Memory() {
        if (cachegetUint8Memory === null || cachegetUint8Memory.buffer !== wasm.memory.buffer) {
            cachegetUint8Memory = new Uint8Array(wasm.memory.buffer);
        }
        return cachegetUint8Memory;
    }
    
    function getStringFromWasm(ptr, len) {
        return cachedDecoder.decode(getUint8Memory().subarray(ptr, ptr + len));
    }
    
    __exports.__wbg_log_f74b4eaeae4e758a = function(arg0, arg1) {
        let varg0 = getStringFromWasm(arg0, arg1);
        __wbg_log_f74b4eaeae4e758a_target(varg0);
    };
    
    function init(wasm_path) {
        return fetch(wasm_path)
        .then(response => response.arrayBuffer())
        .then(buffer => WebAssembly.instantiate(buffer, { './testsocket': __exports }))
        .then(({instance}) => {
            wasm = init.wasm = instance.exports;
            return;
        });
    };
    self.testsocket_wasm = Object.assign(init, __exports);
})();

