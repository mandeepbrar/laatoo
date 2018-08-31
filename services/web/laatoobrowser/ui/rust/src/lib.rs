extern crate laatoocore;
extern crate wasm_bindgen;
use wasm_bindgen::prelude::*;
use futures::{future, Future};
use wasm_bindgen_futures::future_to_promise;
extern crate web_sys;

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

impl laatoocore::Platform for Browser {

    fn http_call(&self, req request::Request, success success_callback, error error_callback) {
        let mut request_options = web_sys::RequestInit::new();
        request_options.method(req.Method);
        request_options.mode(web_sys::RequestMode::Cors);
        let req_to_send = web_sys::Request::new_using_usv_str_and_request_init(req.URL, &request_options).unwrap();
        req_to_send.headers().set("Accept", "application/vnd.github.v3+json").unwrap();
        let req_promise = web_sys::window.fetch_using_request(&req);
        let to_return = JsFuture::from(req_promise).and_then(|resp_value| {
                // resp_value is a Response object
                assert!(resp_value.is_instance_of::<Response>());
                let resp: Response = resp_value.dyn_into().unwrap();

                resp.json()


            }).and_then(|json_value: Promise| {
                // convert this other promise into a rust Future
                JsFuture::from(json_value)
            }).and_then(|json| {
                // Use serde to parse this into a struct
                let branch_info: Branch = json.into_serde().unwrap();

                // Send the Branch struct back to javascript as an object
                future::ok(JsValue::from_serde(&branch_info).unwrap())
            });

    }

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