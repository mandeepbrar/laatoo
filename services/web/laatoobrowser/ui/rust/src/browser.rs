use wasm_bindgen::prelude::*;

use laatoocore::utils::StringsMap;


#[wasm_bindgen]
extern {
    #[wasm_bindgen(js_namespace = console)]
    fn log(msg: &str);
}

#[wasm_bindgen]
pub struct Browser {

}
#[wasm_bindgen]
impl Browser {

    #[allow(dead_code)]
    pub fn log(msg: String) {
    }

    #[allow(dead_code)]
    pub fn execute_service(_service_name: String, _data_json: String, _config_json: String) {
        log(&_service_name);
    }    
}
