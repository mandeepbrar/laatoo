

pub trait Platform {
    fn http_call(&self, request::Request, success_callback, error_callback);
}
