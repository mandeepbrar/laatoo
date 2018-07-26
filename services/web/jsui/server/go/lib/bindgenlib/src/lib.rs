#![feature(libc)]
extern crate wasm_bindgen_cli_support;

extern crate libc;
use libc::c_char;
use std::ffi::CStr;

use wasm_bindgen_cli_support::Bindgen;

#[no_mangle]
pub extern "C" fn bindgen(input: *const c_char, out_dir: *const c_char) ->  i32 {
    let inp_str: &CStr = unsafe{CStr::from_ptr(input)};
    let outdir_str: &CStr = unsafe{CStr::from_ptr(out_dir)};
    let mut b = Bindgen::new();
    let inp_name = inp_str.to_str().unwrap();
    println!("{}", inp_name);
    b.input_path(inp_name);
    /*
        .nodejs(args.flag_nodejs)
        .browser(args.flag_browser)
        .no_modules(args.flag_no_modules)
        .debug(args.flag_debug)
        .demangle(!args.flag_no_demangle)
        .keep_debug(args.flag_keep_debug)
        .typescript(typescript);
    if let Some(ref name) = args.flag_no_modules_global {
        b.no_modules_global(name);
    }*/

/*    let out_dir = match args.flag_out_dir {
        Some(ref p) => p,
        None => bail!("the `--out-dir` argument is now required"),
    };
*/
    let res = b.generate(outdir_str.to_str().unwrap());
    match res {
        Result::Ok(_) => println!("Wasm bindgen successfull"),
        Result::Err(err) => {println!("Error in wasm bindgen {}", err); return -1;}
    }
    //res
    return 0;
}


#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
