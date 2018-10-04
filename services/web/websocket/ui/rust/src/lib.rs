#![feature(use_extern_macros)]

extern crate wasm_bindgen;
extern crate laatoocore;

use wasm_bindgen::prelude::*;
//use utils::{StringMap, StringMapValue};
use laatoocore::redux::{Store, Action, Reducer};
use laatoocore::event::{EventListener, Event};
use laatoocore::platform::{Platform, SuccessCallback, ErrorCallback};

#[wasm_bindgen]
extern {
    #[wasm_bindgen(js_namespace = window)]
    fn alert(s: &str);

}

/*
// A macro to provide `println!(..)`-style syntax for `console.log` logging.
macro_rules! log {
    ($($t:tt)*) => (log(&format!($($t)*)))
}*/

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}


#[no_mangle]
pub extern fn add_one(a: u32) -> u32 {
    alert("add one");
    a + 1
}

#[derive(Clone, Debug)]
struct TestData {
    testdata: Vec<i32>,
}
impl Store for TestData {
    fn initialize(&self) {
        //TestData{testdata: vec![]}
    }
    fn get_id(&self) -> &'static str {
        "Test Data"
    }
    fn as_any(&self) -> &dyn Any {
        self
    }

}

impl Reducer for TestData {
    fn reduce(&mut self, action: &Action) -> Result<bool, String> {
        match (*action).as_any().downcast_ref::<TestStoreAction>() {
            Some(act) =>  {
                println!("Hello World{:?}", act);
            },
            None => panic!("Wrong action type!"),
        };
        /*let act = action as &TestStoreAction;
        match act {
            TestStoreAction::Add(val) => {
                self.testdata.push(*val);
            }
        }*/
        println!("reduced");
        Ok(true)
    }

}

#[derive(Debug)]
enum TestStoreAction {
    Add(i32),
}

    impl Action for TestStoreAction {
        fn get_type(&self)->&'static str {
            return "TestStoreAction";
        }
        fn as_any(&self) -> &dyn Any {
            self
        }
    }

    #[test]
    fn store_works() {
        let mut app = create_application();
        let str = Box::new(TestData{testdata: vec![]});
        let act = TestStoreAction::Add(2);
        let str_id = str.get_id();
        app.register_store(str, act.get_type());
       // let lsr = Box::new(TestListener{});
        app.register_listener(str_id, |stor| {
            println!("event received {:?}", stor);
        });
        app.dispatch(&act);
        assert_eq!(2 + 2, 4);
    }

    fn create_application() -> Application {
        Application::new(Box::new(TestPlatform{}))
    }
