#![feature(wasm_import_module)]

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}

#[wasm_import_module="websocket"]
extern {
    fn add_one(a: u32) -> u32 ;
}



#[no_mangle]
pub extern fn my_func() -> u32 {
    let mut myvar = unsafe { add_one(1) };
    myvar
}