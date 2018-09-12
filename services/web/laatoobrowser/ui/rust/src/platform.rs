extern crate laatoocore;
extern crate wasm_bindgen;
extern crate web_sys;
extern crate js_sys;
extern crate futures;
extern crate wasm_bindgen_futures;

use wasm_bindgen::prelude::*;
use wasm_bindgen::JsCast;
use futures::{future, Future};
//use wasm_bindgen_futures::future_to_promise;
use web_sys::{Request as WSRequest, RequestInit, RequestMode, Response, Window};
use laatoocore::{http, platform::Platform, platform::SuccessCallback, platform::ErrorCallback};
use wasm_bindgen_futures::JsFuture;
use std::marker::{Sync, Send};
use js_sys::Promise;
/*

#[wasm_bindgen]
extern "C" {
    static window: Window;
}*/

pub struct LaatooBrowser {
}

impl Platform for LaatooBrowser {

    fn http_call(&self, url: String, method: http::HttpMethod, req: http::HttpRequest, success: SuccessCallback, error: ErrorCallback) {
        let mut request_options = RequestInit::new();
        request_options.method(&method.to_string());
       // request_options.mode(web_sys::RequestMode::Cors);
        let req_to_send = WSRequest::new_with_str_and_init(&url, &request_options).unwrap();
        for (hdr_name, hdr_value) in req.headers.iter() {
            req_to_send.headers().set(&hdr_name, &hdr_value).unwrap();
        }
        
        let req_promise = Window::fetch_with_request(&req_to_send);
        JsFuture::from(req_promise).and_then(|resp_value| {
                // resp_value is a Response object
                assert!(resp_value.is_instance_of::<Response>());
                let resp: Response = resp_value.dyn_into().unwrap();

                resp.json()


            }).and_then(|json_value: Promise| {
                // convert this other promise into a rust Future
                JsFuture::from(json_value)
            });/*.and_then(|json| {
                // Use serde to parse this into a struct
                //let branch_info: Branch = json.into_serde().unwrap();

                // Send the Branch struct back to javascript as an object
                //future::ok(JsValue::from_serde(&branch_info).unwrap())
                //success()
            });*/

    }

}

unsafe impl Sync for LaatooBrowser{}
unsafe impl Send for LaatooBrowser{}