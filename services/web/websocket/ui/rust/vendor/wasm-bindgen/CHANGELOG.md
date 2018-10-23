# `wasm-bindgen` Change Log

--------------------------------------------------------------------------------

## Unreleased

Released YYYY-MM-DD.

### Added

* TODO (or remove section if none)

### Changed

* TODO (or remove section if none)

### Deprecated

* TODO (or remove section if none)

### Removed

* TODO (or remove section if none)

### Fixed

* TODO (or remove section if none)

### Security

* TODO (or remove section if none)

--------------------------------------------------------------------------------

## 0.2.25

Released 2018-10-10.

### Fixed

* Using `wasm-bindgen` will no longer unconditionally pull in Rust's default
  allocator for Wasm (dlmalloc) regardless if you configured a custom global
  allocator (eg wee_alloc).
  [#947](https://github.com/rustwasm/wasm-bindgen/pull/947)

* Fixed web-sys build on some Windows machines.
  [#943](https://github.com/rustwasm/wasm-bindgen/issues/943)

* Fixed generated ES class bindings to Rust structs that were only referenced
  through struct fields.
  [#948](https://github.com/rustwasm/wasm-bindgen/issues/948)

--------------------------------------------------------------------------------

## 0.2.24

Released 2018-10-05.

### Added

* Constructors for types in `web-sys` should now have better documentation.

* A new `vendor_prefix` attribute in `#[wasm_bindgen]` is supported to bind APIs
  on the web which may have a vendor prefix (like `webkitAudioContext`). This is
  then subsequently used to fix `AudioContext` usage in Safari.

* The `#[wasm_bindgen(extends = Foo)]` attribute now supports full paths, so you
  can also say `#[wasm_bindgen(extends = foo::Bar)]` and such.

### Changed

* The `Closure<T>` type is now optimized when the underlying closure is a ZST.
  The type now no longer allocates memory in this situation.

* The documentation now has a list of caveats for browser support, including how
  `TextEncoder` and `TextDecoder` are not implemented in Edge. If you're using
  webpack there's a listed strategy available, and improvements to the polyfill
  strategy are always welcome!

* The `BaseAudioContext` and `AudioScheduledSourceNode` types in `web-sys` have
  been deprecated as they don't exist in Safari or Edge.

### Fixed

* Fixed the `#[wasm_bindgen_test]`'s error messages in a browser to correctly
  escape HTML-looking output.

* WebIDL Attributes on `Window` are now correctly bound to not go through
  `Window.prototype` which doesn't exist but instead use a `structural`
  definition.

* Fixed a codegen error when the `BorrowMut` trait was in scope.

* Fixed TypeScript generation for constructors of classes, it was accidentally
  producing a syntactially invalid file!

--------------------------------------------------------------------------------

## 0.2.23

Released 2018-09-26.

### Added

* [This is the first release of the `web-sys`
  crate!](https://rustwasm.github.io/2018/09/26/announcing-web-sys.html)

* Added support for unions of interfaces and non-interfaces in the WebIDL
  frontend.

* Added a policy for inclusion of new ECMAScript features in `js-sys`: the
  feature must be in stage 4 or greater for us to support it.

* Added some documentation about size profiling and optimization with
  `wasm-bindgen` to the guide.

* Added the `Clamped<T>` type for generating JavaScript `Uint8ClampedArray`s.

* CI is now running on beta! Can't wait for the `rustc` release trains to roll
  over, so we can run CI on stable too!

* Added the `js_sys::try_iter` function, which checks arbitrary JS values for
  compliance with the JS iteration protocol, and if they are iterable, converts
  them into an iterator over the JS values that they yield.

### Changed

* We now only generate null checks on methods on the JS side when in debug
  mode. For safety we will always null check on the Rust side, however.

* Improved error messages when defining setters that don't start with `set_` and
  don't use `js_name = ...`.

* Improved generated code for classes in a way that avoids an unnecessary
  allocation with static methods that return `Self` but are not the "main"
  constructor.

* **BREAKING:** `js_sys::Reflect` APIs are all fallible now. This is because
  reflecting on `Proxy`s whose trap handlers throw an exception can cause any of
  the reflection APIs to throw. Accordingly, `js_sys` has been bumped from
  `0.2.X` to `0.3.X`.

### Fixed

* The method of ensuring that `__wbindgen_malloc` and `__wbindgen_free` are
  always emitted in the `.wasm` binary, regardless of seeming reachability is
  now zero-overhead.

--------------------------------------------------------------------------------

## 0.2.22

Released 2018-09-21

### Added

* The `IntoIterator` trait is now implemented for JS `Iterator` types
* A number of variadic methods in `js-sys` have had explicit arities added.
* The guide has been improved quite a bit as well as enhanced with more examples
* The `js-sys` crate is now complete! Thanks so much to everyone involved to
  help fill out all the APIs.
* Exported Rust functions with `#[wasm_bindgen]` can now return a `Result` where
  the `Err` payload is raised as an exception in JS.

### Fixed

* An issue with running `wasm-bindgen` on crates that have been compiled with
  LTO has been resolved.

--------------------------------------------------------------------------------

## 0.2.21

Released 2018-09-07

### Added

* Added many more bindings for `WebAssembly` in the `js-sys` crate.

### Fixed

* The "names" section of the wasm binary is now correctly preserved by
  wasm-bindgen.

--------------------------------------------------------------------------------

## 0.2.20

Released 2018-09-06

### Added

* All of `wasm-bindgen` is configured to compile on stable Rust as of the
  upcoming 1.30.0 release, scheduled for October 25, 2018.
* The underlying `JsValue` of a `Closure<T>` type can now be extracted at any
  time.
* Initial and experimental support was added for modules that have shared memory
  (use atomic instructions).

### Removed

* The `--wasm2asm` flag of `wasm2es6js` was removed because the `wasm2asm` tool
  has been removed from upstream Binaryen. This is replaced with the new
  `wasm2js` tool from Binaryen.

### Fixed

* The "schema" version for wasm-bindgen now changes on all publishes, meaning we
  can't forget to update it. This means that the crate version and CLI version
  must exactly match.
* The `wasm-bindgen` crate now has a `links` key which forbids multiple versions
  of `wasm-bindgen` from being linked into a dependency graph, fixing obscure
  linking errors with a more first-class error message.
* Binary releases for Windows has been fixed.

--------------------------------------------------------------------------------

## 0.2.19 (and 0.2.18)

Released 2018-08-27.

### Added

* Added bindings to `js-sys` for some `WebAssembly` types.
* Added bindings to `js-sys` for some `Intl` types.
* Added bindings to `js-sys` for some `String` methods.
* Added an example of using the WebAudio APIs.
* Added an example of using the `fetch` API.
* Added more `extends` annotations for types in `js-sys`.
* Experimental support for `WeakRef` was added to automatically deallocate Rust
  objects when gc'd.
* Added support for executing `wasm-bindgen` over modules that import their
  memory.
* Added a global `memory()` function in the `wasm-bindgen` crate for accessing
  the JS object that represent wasm's own memory.

### Removed

* Removed `AsMut` implementations for imported objects.

### Fixed

* Fixed the `constructor` and `catch` attributes combined on imported types.
* Fixed importing the same-named static in two modules.

--------------------------------------------------------------------------------

## 0.2.17

Released 2018-08-16.

### Added

* Greatly expanded documentation in the wasm-bindgen guide.
* Added bindings to `js-sys` for `Intl.DateTimeFormat`
* Added a number of `extends` attributes for types in `js-sys`

### Fixed

* Fixed compile on latest nightly with latest `proc-macro2`
* Fixed compilation in some scenarios on Windows with paths in `module` paths

--------------------------------------------------------------------------------

## 0.2.16

Released 2018-08-13.

### Added

* Added the `wasm_bindgen::JsCast` trait, as described in [RFC #2][rfc-2].
* Added [the `#[wasm_bindgen(extends = ...)]` attribute][extends-attr] to
  describe inheritance relationships, as described in [RFC #2][rfc-2].
* Added support for receiving `Option<&T>` parameters from JavaScript in
  exported Rust functions and methods.
* Added support for receiving `Option<u32>` and other option-wrapped scalars.
* Added reference documentation to the guide for every `#[wasm_bindgen]`
  attribute and how it affects the generated bindings.
* Published the `wasm-bindgen-futures` crate for converting between JS
  `Promise`s and Rust `Future`s.

### Changed

* Overhauled the guide's documentation on passing JS closures to Rust, and Rust
  closures to JS.
* Overhauled the guide's documentation on using serde to serialize complex data
  to `JsValue` and deserialize `JsValue`s back into complex data.
* Static methods are now always bound to their JS class, as is required for
  `Promise`'s static methods.

### Removed

* Removed internal usage of `syn`'s `visit-mut` cargo feature, which should
  result in faster build times.

### Fixed

* Various usage errors for the `#[wasm_bindgen]` proc-macro are now properly
  reported with source span information, rather than `panic!()`s inside the
  proc-macro.
* Fixed a bug where taking a struct by reference and returning a slice resulted
  in lexical variable redeclaration errors in the generated JS glue. [#662][]
* The `#[wasm_bindgen(js_class = "....")]` attribute for binding methods to
  renamed imported JS classes now properly works with constructors.

[rfc-2]: https://rustwasm.github.io/rfcs/002-wasm-bindgen-inheritance-casting.html
[extends-attr]: https://rustwasm.github.io/wasm-bindgen/reference/attributes/on-js-imports/extends.html
[#662]: https://github.com/rustwasm/wasm-bindgen/pull/662

--------------------------------------------------------------------------------

## 0.2.15

Released 2018-07-26.

### Fixed

* Fixed `wasm-bindgen` CLI version mismatch checks that got broken in the last
  point release.

--------------------------------------------------------------------------------

## 0.2.14

Released 2018-07-25.

### Fixed

* Fixed compilation errors on targets that use
  Mach-O. [#545](https://github.com/rustwasm/wasm-bindgen/issues/545)

--------------------------------------------------------------------------------

## 0.2.13

Released 2018-07-22.

### Added

* Support the `#[wasm_bindgen(js_name = foo)]` attribute on exported functions
  and methods to allow renaming an export to JS. This allows JS to call it by
  one name and Rust to call it by another, for example using `camelCase` in JS
  and `snake_case` in Rust

### Fixed

* Compilation with the latest nightly compiler has been fixed (nightlies on and
  after 2018-07-21)

--------------------------------------------------------------------------------

## 0.2.12

Released 2018-07-19.

This release is mostly internal refactorings and minor improvements to the
existing crates and functionality, but the bigs news is an upcoming `js-sys` and
`web-sys` set of crates. The `js-sys` crate will expose [all global JS
bindings][js-all] and the `web-sys` crate will be generated from WebIDL to
expose all APIs browsers have. More info on this soon!

[js-all]: https://github.com/rustwasm/wasm-bindgen/issues/275

### Added

* Support for `Option<T>` was added where `T` can be a number of slices or
  imported types.
* Comments in Rust are now preserved in generated JS bindings, as well as
  comments being generated to indicate the types of arguments/return values.
* The online documentation has been reorganized [into a book][book].
* The generated JS is now formatted better by default for readability.
* A `--keep-debug` flag has been added to the CLI to retain debug sections by
  default. This happens by default when `--debug` is passed.

[book]: https://rustwasm.github.io/wasm-bindgen/

### Fixed

* Compilation with the latest nightly compiler has been fixed (nightlies on and
  after 2018-07-19)
* Declarations of an imported function in multiple crates have been fixed to not
  conflict.
* Compilation with `#![deny(missing_docs)]` has been fixed.

--------------------------------------------------------------------------------

## 0.2.11

Released 2018-05-24.

--------------------------------------------------------------------------------

## 0.2.10

Released 2018-05-17.

--------------------------------------------------------------------------------

## 0.2.9

Released 2018-05-11.
