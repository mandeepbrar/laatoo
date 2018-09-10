extern crate laatoocore;
extern crate web_sys;
extern crate js_sys;
extern crate futures;
extern crate wasm_bindgen_futures;
extern crate wasm_bindgen;
use wasm_bindgen::prelude::*;

mod platform;

use std::sync::Once;
use laatoocore::{application::Application};

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}

static INIT: Once = Once::new();

#[cfg_attr(target_arch = "wasm32", wasm_bindgen)]
pub fn initialize() {
    //static mut App: Box<Application>;
    INIT.call_once(|| {
        Application.lock().unwrap().initialize(Box::new(platform::Browser{}));
        //let app= Application{pf: Box::new(platform::Browser{})};
        //App = Box::new(app);
    });
}

/*
        use web_sys::{Request, RequestInit, RequestMode, Response, Window};


#[wasm_bindgen]
pub fn run() -> Promise {


    // the RequestInit struct will eventually support setting headers, but that's missing right now


    // Convert this rust future back into a javascript promise.
    // Return it to javascript so that it can be driven to completion.
    future_to_promise(to_return)
}*/