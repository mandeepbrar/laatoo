//! Support for long-lived closures in `wasm-bindgen`
//!
//! This module defines the `Closure` type which is used to pass "owned
//! closures" from Rust to JS. Some more details can be found on the `Closure`
//! type itself.

#[cfg(feature = "nightly")]
use std::marker::Unsize;
use std::mem::{self, ManuallyDrop};
use std::prelude::v1::*;

use JsValue;
use convert::*;
use describe::*;
use throw_str;

/// A handle to both a closure in Rust as well as JS closure which will invoke
/// the Rust closure.
///
/// A `Closure` is the primary way that a `'static` lifetime closure is
/// transferred from Rust to JS. `Closure` currently requires that the closures
/// it's created with have the `'static` lifetime in Rust for soundness reasons.
///
/// This type is a "handle" in the sense that whenever it is dropped it will
/// invalidate the JS closure that it refers to. Any usage of the closure in JS
/// after the `Closure` has been dropped will raise an exception. It's then up
/// to you to arrange for `Closure` to be properly deallocate at an appropriate
/// location in your program.
///
/// The type parameter on `Closure` is the type of closure that this represents.
/// Currently this can only be the `Fn` and `FnMut` traits with up to 7
/// arguments (and an optional return value). The arguments/return value of the
/// trait must be numbers like `u32` for now, although this restriction may be
/// lifted in the future!
///
/// # Example
///
/// Sample usage of `Closure` to invoke the `setTimeout` API.
///
/// ```rust,no_run
/// #[wasm_bindgen]
/// extern {
///     fn setTimeout(closure: &Closure<FnMut()>, time: u32);
///
///     #[wasm_bindgen(js_namespace = console)]
///     fn log(s: &str);
/// }
///
/// #[wasm_bindgen]
/// pub struct ClosureHandle(Closure<FnMut()>);
///
/// #[wasm_bindgen]
/// pub fn run() -> ClosureHandle {
///     // First up we use `Closure::wrap` to wrap up a Rust closure and create
///     a JS closure.
///     let cb = Closure::wrap(Box::new(move || {
///         log("timeout elapsed!");
///     }) as Box<FnMut()>);
///
///     // Next we pass this via reference to the `setTimeout` function, and
///     // `setTimeout` gets a handle to the corresponding JS closure.
///     setTimeout(&cb, 1_000);
///
///     // If we were to drop `cb` here it would cause an exception to be raised
///     // when the timeout elapses. Instead we *return* our handle back to JS
///     // so JS can tell us later when it would like to deallocate this handle.
///     ClosureHandle(cb)
/// }
/// ```
///
/// Sample usage of the same example as above except using `web_sys` instead
///
/// ```rust,no_run
/// extern crate wasm_bindgen;
/// extern crate web_sys;
///
/// use wasm_bindgen::JsCast;
///
/// #[wasm_bindgen]
/// pub struct ClosureHandle(Closure<FnMut()>);
///
/// #[wasm_bindgen]
/// pub fn run() -> ClosureHandle {
///     let cb = Closure::wrap(Box::new(move || {
///         web_sys::console::log_1(&"timeout elapsed!".into());
///     }) as Box<FnMut()>);
///
///     let window = web_sys::window().unwrap();
///     window.set_timeout_with_callback_and_timeout_and_arguments_0(
///         // Note this method call, which uses `as_ref()` to get a `JsValue`
///         // from our `Closure` which is then converted to a `&Function`
///         // using the `JsCast::unchecked_ref` function.
///         cb.as_ref().unchecked_ref(),
///         1_000,
///     );
///
///     // same as above
///     ClosureHandle(cb)
/// }
/// ```
pub struct Closure<T: ?Sized> {
    js: ManuallyDrop<JsValue>,
    data: ManuallyDrop<Box<T>>,
}

union FatPtr<T: ?Sized> {
    ptr: *mut T,
    fields: (usize, usize),
}

impl<T> Closure<T>
    where T: ?Sized + WasmClosure,
{
    /// Creates a new instance of `Closure` from the provided Rust closure.
    ///
    /// Note that the closure provided here, `F`, has a few requirements
    /// associated with it:
    ///
    /// * It must implement `Fn` or `FnMut`
    /// * It must be `'static`, aka no stack references (use the `move` keyword)
    /// * It can have at most 7 arguments
    /// * Its arguments and return values are all wasm types like u32/f64.
    ///
    /// This is unfortunately pretty restrictive for now but hopefully some of
    /// these restrictions can be lifted in the future!
    ///
    /// *This method requires the `nightly` feature of the `wasm-bindgen` crate
    /// to be enabled, meaning this is a nightly-only API. Users on stable
    /// should use `Closure::wrap`.*
    #[cfg(feature = "nightly")]
    pub fn new<F>(t: F) -> Closure<T>
        where F: Unsize<T> + 'static
    {
        Closure::wrap(Box::new(t) as Box<T>)
    }

    /// A mostly internal function to wrap a boxed closure inside a `Closure`
    /// type.
    ///
    /// This is the function where the JS closure is manufactured.
    pub fn wrap(mut data: Box<T>) -> Closure<T> {
        assert_eq!(mem::size_of::<*const T>(), mem::size_of::<FatPtr<T>>());
        let (a, b) = unsafe {
            FatPtr { ptr: &mut *data as *mut T }.fields
        };

        // Here we need to create a `JsValue` with the data and `T::invoke()`
        // function pointer. To do that we... take a few unconventional turns.
        // In essence what happens here is this:
        //
        // 1. First up, below we call a function, `breaks_if_inlined`. This
        //    function, as the name implies, does not work if it's inlined.
        //    More on that in a moment.
        // 2. This function internally calls a special import recognized by the
        //    `wasm-bindgen` CLI tool, `__wbindgen_describe_closure`. This
        //    imported symbol is similar to `__wbindgen_describe` in that it's
        //    not intended to show up in the final binary but it's an
        //    intermediate state for a `wasm-bindgen` binary.
        // 3. The `__wbindgen_describe_closure` import is namely passed a
        //    descriptor function, monomorphized for each invocation.
        //
        // Most of this doesn't actually make sense to happen at runtime! The
        // real magic happens when `wasm-bindgen` comes along and updates our
        // generated code. When `wasm-bindgen` runs it performs a few tasks:
        //
        // * First, it finds all functions that call
        //   `__wbindgen_describe_closure`. These are all `breaks_if_inlined`
        //   defined below as the symbol isn't called anywhere else.
        // * Next, `wasm-bindgen` executes the `breaks_if_inlined`
        //   monomorphized functions, passing it dummy arguments. This will
        //   execute the function just enough to invoke the special import,
        //   namely telling us about the function pointer that is the describe
        //   shim.
        // * This knowledge is then used to actually find the descriptor in the
        //   function table which is then executed to figure out the signature
        //   of the closure.
        // * Finally, and probably most heinously, the call to
        //   `breaks_if_inlined` is rewritten to call an otherwise globally
        //   imported function. This globally imported function will generate
        //   the `JsValue` for this closure specialized for the signature in
        //   question.
        //
        // Later on `wasm-gc` will clean up all the dead code and ensure that
        // we don't actually call `__wbindgen_describe_closure` at runtime. This
        // means we will end up not actually calling `breaks_if_inlined` in the
        // final binary, all calls to that function should be pruned.
        //
        // See crates/cli-support/src/js/closures.rs for a more information
        // about what's going on here.

        extern fn describe<T: WasmClosure + ?Sized>() {
            inform(CLOSURE);
            T::describe()
        }

        #[inline(never)]
        unsafe fn breaks_if_inlined<T: WasmClosure + ?Sized>(
            a: usize,
            b: usize,
            invoke: u32,
            destroy: u32,
        ) -> u32 {
            super::__wbindgen_describe_closure(
                a as u32,
                b as u32,
                invoke,
                destroy,
                describe::<T> as u32,
            )
        }

        let idx = unsafe {
            breaks_if_inlined::<T>(a, b, T::invoke_fn(), T::destroy_fn())
        };

        Closure {
            js: ManuallyDrop::new(JsValue { idx }),
            data: ManuallyDrop::new(data),
        }
    }

    /// Leaks this `Closure` to ensure it remains valid for the duration of the
    /// entire program.
    ///
    /// > **Note**: this function will leak memory. It should be used sparingly
    /// > to ensure the memory leak doesn't affect the program too much.
    ///
    /// When a `Closure` is dropped it will invalidate the associated JS
    /// closure, but this isn't always desired. Some callbacks are alive for
    /// the entire duration of the program, so this can be used to conveniently
    /// leak this instance of `Closure` while performing as much internal
    /// cleanup as it can.
    pub fn forget(self) {
        unsafe {
            super::__wbindgen_cb_forget(self.js.idx);
            mem::forget(self);
        }
    }
}

impl<T: ?Sized> AsRef<JsValue> for Closure<T> {
    fn as_ref(&self) -> &JsValue {
        &self.js
    }
}

impl<T> WasmDescribe for Closure<T>
    where T: WasmClosure + ?Sized,
{
    fn describe() {
        inform(ANYREF);
    }
}

// `Closure` can only be passed by reference to imports.
impl<'a, T> IntoWasmAbi for &'a Closure<T>
    where T: WasmClosure + ?Sized,
{
    type Abi = u32;

    fn into_abi(self, extra: &mut Stack) -> u32 {
        (&*self.js).into_abi(extra)
    }
}

fn _check() {
    fn _assert<T: IntoWasmAbi>() {}
    _assert::<&Closure<Fn()>>();
    _assert::<&Closure<Fn(String)>>();
    _assert::<&Closure<Fn() -> String>>();
    _assert::<&Closure<FnMut()>>();
    _assert::<&Closure<FnMut(String)>>();
    _assert::<&Closure<FnMut() -> String>>();
}

impl<T> Drop for Closure<T>
    where T: ?Sized,
{
    fn drop(&mut self) {
        unsafe {
            // this will implicitly drop our strong reference in addition to
            // invalidating all future invocations of the closure
            if super::__wbindgen_cb_drop(self.js.idx) != 0 {
                ManuallyDrop::drop(&mut self.data);
            }
        }
    }
}

/// An internal trait for the `Closure` type.
///
/// This trait is not stable and it's not recommended to use this in bounds or
/// implement yourself.
#[doc(hidden)]
pub unsafe trait WasmClosure: 'static {
    fn describe();

    fn invoke_fn() -> u32;
    fn destroy_fn() -> u32;
}

// The memory safety here in these implementations below is a bit tricky. We
// want to be able to drop the `Closure` object from within the invocation of a
// `Closure` for cases like promises. That means that while it's running we
// might drop the `Closure`, but that shouldn't invalidate the environment yet.
//
// Instead what we do is to wrap closures in `Rc` variables. The main `Closure`
// has a strong reference count which keeps the trait object alive. Each
// invocation of a closure then *also* clones this and gets a new reference
// count. When the closure returns it will release the reference count.
//
// This means that if the main `Closure` is dropped while it's being invoked
// then destruction is deferred until execution returns. Otherwise it'll
// deallocate data immediately.

macro_rules! doit {
    ($(
        ($($var:ident)*)
    )*) => ($(
        unsafe impl<$($var,)* R> WasmClosure for Fn($($var),*) -> R
            where $($var: FromWasmAbi + 'static,)*
                  R: ReturnWasmAbi + 'static,
        {
            fn describe() {
                <&Self>::describe();
            }

            #[inline]
            fn invoke_fn() -> u32 {
                #[allow(non_snake_case)]
                unsafe extern fn invoke<$($var: FromWasmAbi,)* R: ReturnWasmAbi>(
                    a: usize,
                    b: usize,
                    $($var: <$var as FromWasmAbi>::Abi),*
                ) -> <R as ReturnWasmAbi>::Abi {
                    if a == 0 {
                        throw_str("closure invoked recursively or destroyed already");
                    }
                    // Make sure all stack variables are converted before we
                    // convert `ret` as it may throw (for `Result`, for
                    // example)
                    let ret = {
                        let f: *const Fn($($var),*) -> R =
                            FatPtr { fields: (a, b) }.ptr;
                        let mut _stack = GlobalStack::new();
                        $(
                            let $var = <$var as FromWasmAbi>::from_abi($var, &mut _stack);
                        )*
                        (*f)($($var),*)
                    };
                    ret.return_abi(&mut GlobalStack::new())
                }
                invoke::<$($var,)* R> as u32
            }

            #[inline]
            fn destroy_fn() -> u32 {
                unsafe extern fn destroy<$($var: FromWasmAbi,)* R: ReturnWasmAbi>(
                    a: usize,
                    b: usize,
                ) {
                    debug_assert!(a != 0);
                    drop(Box::from_raw(FatPtr::<Fn($($var,)*) -> R> {
                        fields: (a, b)
                    }.ptr));
                }
                destroy::<$($var,)* R> as u32
            }
        }

        unsafe impl<$($var,)* R> WasmClosure for FnMut($($var),*) -> R
            where $($var: FromWasmAbi + 'static,)*
                  R: ReturnWasmAbi + 'static,
        {
            fn describe() {
                <&mut Self>::describe();
            }

            #[inline]
            fn invoke_fn() -> u32 {
                #[allow(non_snake_case)]
                unsafe extern fn invoke<$($var: FromWasmAbi,)* R: ReturnWasmAbi>(
                    a: usize,
                    b: usize,
                    $($var: <$var as FromWasmAbi>::Abi),*
                ) -> <R as ReturnWasmAbi>::Abi {
                    if a == 0 {
                        throw_str("closure invoked recursively or destroyed already");
                    }
                    // Make sure all stack variables are converted before we
                    // convert `ret` as it may throw (for `Result`, for
                    // example)
                    let ret = {
                        let f: *const FnMut($($var),*) -> R =
                            FatPtr { fields: (a, b) }.ptr;
                        let f = f as *mut FnMut($($var),*) -> R;
                        let mut _stack = GlobalStack::new();
                        $(
                            let $var = <$var as FromWasmAbi>::from_abi($var, &mut _stack);
                        )*
                        (*f)($($var),*)
                    };
                    ret.return_abi(&mut GlobalStack::new())
                }
                invoke::<$($var,)* R> as u32
            }

            #[inline]
            fn destroy_fn() -> u32 {
                unsafe extern fn destroy<$($var: FromWasmAbi,)* R: ReturnWasmAbi>(
                    a: usize,
                    b: usize,
                ) {
                    debug_assert!(a != 0);
                    drop(Box::from_raw(FatPtr::<FnMut($($var,)*) -> R> {
                        fields: (a, b)
                    }.ptr));
                }
                destroy::<$($var,)* R> as u32
            }
        }
    )*)
}

doit! {
    ()
    (A)
    (A B)
    (A B C)
    (A B C D)
    (A B C D E)
    (A B C D E F)
    (A B C D E F G)
}
