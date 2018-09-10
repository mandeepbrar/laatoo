
(function() {
    var wasm;
    const __exports = {};
    
    function init(wasm_path) {
        const fetchPromise = fetch(wasm_path);
        let resultPromise;
        if (typeof WebAssembly.instantiateStreaming === 'function') {
            resultPromise = WebAssembly.instantiateStreaming(fetchPromise, { './laatoobrowser': __exports });
        } else {
            resultPromise = fetchPromise
            .then(response => response.arrayBuffer())
            .then(buffer => WebAssembly.instantiate(buffer, { './laatoobrowser': __exports }));
        }
        return resultPromise.then(({instance}) => {
            wasm = init.wasm = instance.exports;
            return;
        });
    };
    self.laatoobrowser_wasm = Object.assign(init, __exports);
})();

