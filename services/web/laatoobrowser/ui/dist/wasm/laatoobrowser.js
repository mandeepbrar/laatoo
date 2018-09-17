
(function() {
    var wasm;
    const __exports = {};
    /**
    * @returns {Application}
    */
    __exports.initialize = function() {
        return Application.__construct(wasm.initialize());
    };
    
    const __wbg_log_4609d9aeb31a73e9_target = console.log;
    
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
    
    __exports.__wbg_log_4609d9aeb31a73e9 = function(arg0, arg1) {
        let varg0 = getStringFromWasm(arg0, arg1);
        __wbg_log_4609d9aeb31a73e9_target(varg0);
    };
    
    let cachedEncoder = new TextEncoder('utf-8');
    
    function passStringToWasm(arg) {
        
        const buf = cachedEncoder.encode(arg);
        const ptr = wasm.__wbindgen_malloc(buf.length);
        getUint8Memory().set(buf, ptr);
        return [ptr, buf.length];
    }
    
    const __widl_f_set_Headers_target = Headers.prototype.set || function() {
        throw new Error(`wasm-bindgen: Headers.prototype.set does not exist`);
    };
    
    let cachegetUint32Memory = null;
    function getUint32Memory() {
        if (cachegetUint32Memory === null || cachegetUint32Memory.buffer !== wasm.memory.buffer) {
            cachegetUint32Memory = new Uint32Array(wasm.memory.buffer);
        }
        return cachegetUint32Memory;
    }
    
    const slab = [{ obj: undefined }, { obj: null }, { obj: true }, { obj: false }];
    
    let slab_next = slab.length;
    
    function addHeapObject(obj) {
        if (slab_next === slab.length) slab.push(slab.length + 1);
        const idx = slab_next;
        const next = slab[idx];
        
        slab_next = next;
        
        slab[idx] = { obj, cnt: 1 };
        return idx << 1;
    }
    
    const stack = [];
    
    function getObject(idx) {
        if ((idx & 1) === 1) {
            return stack[idx >> 1];
        } else {
            const val = slab[idx >> 1];
            
            return val.obj;
            
        }
    }
    
    __exports.__widl_f_set_Headers = function(arg0, arg1, arg2, arg3, arg4, exnptr) {
        let varg1 = getStringFromWasm(arg1, arg2);
        let varg3 = getStringFromWasm(arg3, arg4);
        try {
            __widl_f_set_Headers_target.call(getObject(arg0), varg1, varg3);
        } catch (e) {
            const view = getUint32Memory();
            view[exnptr / 4] = 1;
            view[exnptr / 4 + 1] = addHeapObject(e);
            
        }
    };
    
    __exports.__widl_f_new_with_str_and_init_Request = function(arg0, arg1, arg2, exnptr) {
        let varg0 = getStringFromWasm(arg0, arg1);
        try {
            return addHeapObject(new Request(varg0, getObject(arg2)));
        } catch (e) {
            const view = getUint32Memory();
            view[exnptr / 4] = 1;
            view[exnptr / 4 + 1] = addHeapObject(e);
            
        }
    };
    
    function GetOwnOrInheritedPropertyDescriptor(obj, id) {
        while (obj) {
            let desc = Object.getOwnPropertyDescriptor(obj, id);
            if (desc) return desc;
            obj = Object.getPrototypeOf(obj);
        }
        throw new Error(`descriptor for id='${id}' not found`);
    }
    
    const __widl_f_headers_Request_target = GetOwnOrInheritedPropertyDescriptor(Request.prototype, 'headers').get || function() {
        throw new Error(`wasm-bindgen: GetOwnOrInheritedPropertyDescriptor(Request.prototype, 'headers').get does not exist`);
    };
    
    __exports.__widl_f_headers_Request = function(arg0) {
        return addHeapObject(__widl_f_headers_Request_target.call(getObject(arg0)));
    };
    
    __exports.__widl_f_fetch_with_request_ = function(arg0) {
        return addHeapObject(fetch(getObject(arg0)));
    };
    
    __exports.__wbg_new_70bbf770ccc49620 = function() {
        return addHeapObject(new Object());
    };
    
    const __wbg_set_278d77d031e5f63c_target = Reflect.set.bind(Reflect) || function() {
        throw new Error(`wasm-bindgen: Reflect.set.bind(Reflect) does not exist`);
    };
    
    __exports.__wbg_set_278d77d031e5f63c = function(arg0, arg1, arg2) {
        return __wbg_set_278d77d031e5f63c_target(getObject(arg0), getObject(arg1), getObject(arg2)) ? 1 : 0;
    };
    
    const __wbg_then_52c189ae4cd8db1a_target = Promise.prototype.then || function() {
        throw new Error(`wasm-bindgen: Promise.prototype.then does not exist`);
    };
    
    let cachedGlobalArgumentPtr = null;
    function globalArgumentPtr() {
        if (cachedGlobalArgumentPtr === null) {
            cachedGlobalArgumentPtr = wasm.__wbindgen_global_argument_ptr();
        }
        return cachedGlobalArgumentPtr;
    }
    
    function getGlobalArgument(arg) {
        const idx = globalArgumentPtr() / 4 + arg;
        return getUint32Memory()[idx];
    }
    
    __exports.__wbg_then_52c189ae4cd8db1a = function(arg0, arg1, arg2) {
        let idxarg1 = getUint32Memory()[arg1 / 4];
        if (idxarg1 === 0xffffffff) {
            let cbarg1 = function(arg0) {
                let a = this.a;
                this.a = 0;
                try {
                    return this.f(a, addHeapObject(arg0));
                    
                } finally {
                    this.a = a;
                    
                }
                
            };
            cbarg1.f = wasm.__wbg_function_table.get(getGlobalArgument(0));
            cbarg1.a = getGlobalArgument(1);
            let real = cbarg1.bind(cbarg1);
            real.original = cbarg1;
            idxarg1 = getUint32Memory()[arg1 / 4] = addHeapObject(real);
        }
        let idxarg2 = getUint32Memory()[arg2 / 4];
        if (idxarg2 === 0xffffffff) {
            let cbarg2 = function(arg0) {
                let a = this.a;
                this.a = 0;
                try {
                    return this.f(a, addHeapObject(arg0));
                    
                } finally {
                    this.a = a;
                    
                }
                
            };
            cbarg2.f = wasm.__wbg_function_table.get(getGlobalArgument(2));
            cbarg2.a = getGlobalArgument(3);
            let real = cbarg2.bind(cbarg2);
            real.original = cbarg2;
            idxarg2 = getUint32Memory()[arg2 / 4] = addHeapObject(real);
        }
        return addHeapObject(__wbg_then_52c189ae4cd8db1a_target.call(getObject(arg0), getObject(idxarg1), getObject(idxarg2)));
    };
    
    const __wbg_log_157f92906a030fef_target = console.log;
    
    __exports.__wbg_log_157f92906a030fef = function(arg0, arg1) {
        let varg0 = getStringFromWasm(arg0, arg1);
        __wbg_log_157f92906a030fef_target(varg0);
    };
    /**
    */
    __exports.HttpMethod = Object.freeze({ GET:0,POST:1,PUT:2,DELETE:3, });
    
    function freeApplication(ptr) {
        
        wasm.__wbg_application_free(ptr);
    }
    /**
    */
    class Application {
        
        static __construct(ptr) {
            return new Application(ptr);
        }
        
        constructor(ptr) {
            this.ptr = ptr;
            
        }
        
        free() {
            const ptr = this.ptr;
            this.ptr = 0;
            freeApplication(ptr);
        }
        /**
        * @param {string} arg0
        * @param {string} arg1
        * @returns {string}
        */
        js_get_registered_item(arg0, arg1) {
            if (this.ptr === 0) {
                throw new Error('Attempt to use a moved value');
            }
            const [ptr0, len0] = passStringToWasm(arg0);
            const [ptr1, len1] = passStringToWasm(arg1);
            const retptr = globalArgumentPtr();
            wasm.application_js_get_registered_item(retptr, this.ptr, ptr0, len0, ptr1, len1);
            const mem = getUint32Memory();
            const rustptr = mem[retptr / 4];
            const rustlen = mem[retptr / 4 + 1];
            
            const realRet = getStringFromWasm(rustptr, rustlen).slice();
            wasm.__wbindgen_free(rustptr, rustlen * 1);
            return realRet;
            
        }
    }
    __exports.Application = Application;
    
    function freeBrowser(ptr) {
        
        wasm.__wbg_browser_free(ptr);
    }
    /**
    */
    class Browser {
        
        static __construct(ptr) {
            return new Browser(ptr);
        }
        
        constructor(ptr) {
            this.ptr = ptr;
            
        }
        
        free() {
            const ptr = this.ptr;
            this.ptr = 0;
            freeBrowser(ptr);
        }
        /**
        * @param {string} arg0
        * @returns {void}
        */
        static log(arg0) {
            const [ptr0, len0] = passStringToWasm(arg0);
            return wasm.browser_log(ptr0, len0);
        }
        /**
        * @param {string} arg0
        * @param {string} arg1
        * @param {string} arg2
        * @returns {void}
        */
        static execute_service(arg0, arg1, arg2) {
            const [ptr0, len0] = passStringToWasm(arg0);
            const [ptr1, len1] = passStringToWasm(arg1);
            const [ptr2, len2] = passStringToWasm(arg2);
            return wasm.browser_execute_service(ptr0, len0, ptr1, len1, ptr2, len2);
        }
    }
    __exports.Browser = Browser;
    
    function dropRef(idx) {
        
        idx = idx >> 1;
        if (idx < 4) return;
        let obj = slab[idx];
        
        obj.cnt -= 1;
        if (obj.cnt > 0) return;
        
        // If we hit 0 then free up our space in the slab
        slab[idx] = slab_next;
        slab_next = idx;
    }
    
    __exports.__wbindgen_object_drop_ref = function(i) {
        dropRef(i);
    };
    
    __exports.__wbindgen_string_new = function(p, l) {
        return addHeapObject(getStringFromWasm(p, l));
    };
    
    __exports.__wbindgen_number_get = function(n, invalid) {
        let obj = getObject(n);
        if (typeof(obj) === 'number') return obj;
        getUint8Memory()[invalid] = 1;
        return 0;
    };
    
    __exports.__wbindgen_is_null = function(idx) {
        return getObject(idx) === null ? 1 : 0;
    };
    
    __exports.__wbindgen_is_undefined = function(idx) {
        return getObject(idx) === undefined ? 1 : 0;
    };
    
    __exports.__wbindgen_boolean_get = function(i) {
        let v = getObject(i);
        if (typeof(v) === 'boolean') {
            return v ? 1 : 0;
        } else {
            return 2;
        }
    };
    
    __exports.__wbindgen_is_symbol = function(i) {
        return typeof(getObject(i)) === 'symbol' ? 1 : 0;
    };
    
    __exports.__wbindgen_string_get = function(i, len_ptr) {
        let obj = getObject(i);
        if (typeof(obj) !== 'string') return 0;
        const [ptr, len] = passStringToWasm(obj);
        getUint32Memory()[len_ptr / 4] = len;
        return ptr;
    };
    
    __exports.__wbindgen_cb_drop = function(i) {
        let obj = getObject(i).original;
        obj.a = obj.b = 0;
        dropRef(i);
    };
    
    __exports.__wbindgen_throw = function(ptr, len) {
        throw new Error(getStringFromWasm(ptr, len));
    };
    
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

