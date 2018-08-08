#![feature(use_extern_macros)]

extern crate wasm_bindgen;

use wasm_bindgen::prelude::*;

#[wasm_bindgen]
extern {
   // fn testalert(s: &str);

    #[wasm_bindgen(js_namespace = console)]
    fn log(msg: &str);
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}


#[no_mangle]
pub extern fn my_func(a: u32) -> u32 {
 //   testalert("add one");
    log("my func");
    a + 1
}