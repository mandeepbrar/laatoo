extern crate laatoocore;
extern crate wasm_bindgen;
extern crate web_sys;
extern crate js_sys;
extern crate futures;
extern crate wasm_bindgen_futures;

use wasm_bindgen::prelude::*;
use futures::{future, Future};
//use wasm_bindgen_futures::future_to_promise;
use web_sys::{Request as WSRequest, RequestInit, RequestMode, Response, Window};
use laatoocore::{application::Application, request::Request, platform::Platform, platform::SuccessCallback, platform::ErrorCallback};
use wasm_bindgen_futures::JsFuture;
use js_sys::Promise;

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}



#[wasm_bindgen]
extern "C" {
    static window: Window;
}

pub struct Browser {
}

impl Platform for Browser {

    fn http_call(&self, req: laatoocore::request::HttpRequest, success: SuccessCallback, error: ErrorCallback) {
        let mut request_options = RequestInit::new();
        request_options.method(req.Method);
       // request_options.mode(web_sys::RequestMode::Cors);
        let req_to_send = WSRequest::new_with_str_and_init(&req.URL, &request_options).unwrap();
        for hdr in req.Headers {
            req_to_send.headers().set(&hdr.Name, &hdr.Value).unwrap();
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
            }).and_then(|json| {
                // Use serde to parse this into a struct
                //let branch_info: Branch = json.into_serde().unwrap();

                // Send the Branch struct back to javascript as an object
                //future::ok(JsValue::from_serde(&branch_info).unwrap())
                //success()
            });

    }

}

fn initialize() -> Box<Application> {
        let app = Application{pf: Box::new(Browser{})};
        Box::new(app)
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