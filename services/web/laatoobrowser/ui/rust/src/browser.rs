use wasm_bindgen::prelude::*;

use laatoocore::utils::StringsMap;
#[wasm_bindgen]
pub struct Browser {

}
#[wasm_bindgen]
impl Browser {
    #[wasm_bindgen(constructor)]
    pub fn new() -> Browser {
        Browser {}
    }

    #[allow(dead_code)]
    pub fn execute_service(_service_name: String, _data_json: String, _config_json: String) {

    }    
}
