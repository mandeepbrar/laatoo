use proc_macro2::{Ident, Span};
use shared;
use syn;

use Diagnostic;

/// An abstract syntax tree representing a rust program. Contains
/// extra information for joining up this rust code with javascript.
#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq))]
#[derive(Default, Clone)]
pub struct Program {
    /// rust -> js interfaces
    pub exports: Vec<Export>,
    /// js -> rust interfaces
    pub imports: Vec<Import>,
    /// rust enums
    pub enums: Vec<Enum>,
    /// rust structs
    pub structs: Vec<Struct>,
    /// rust consts
    pub consts: Vec<Const>,
    /// "dictionaries", generated for WebIDL, which are basically just "typed
    /// objects" in the sense that they represent a JS object with a particular
    /// shape in JIT parlance.
    pub dictionaries: Vec<Dictionary>,
}

/// A rust to js interface. Allows interaction with rust objects/functions
/// from javascript.
#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct Export {
    /// The javascript class name.
    pub class: Option<Ident>,
    /// The type of `self` (either `self`, `&self`, or `&mut self`)
    pub method_self: Option<MethodSelf>,
    /// Whether or not this export is flagged as a constructor, returning an
    /// instance of the `impl` type
    pub is_constructor: bool,
    /// The rust function
    pub function: Function,
    /// Comments extracted from the rust source.
    pub comments: Vec<String>,
    /// The name of the rust function/method on the rust side.
    pub rust_name: Ident,
}

/// The 3 types variations of `self`.
#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub enum MethodSelf {
    /// `self`
    ByValue,
    /// `&mut self`
    RefMutable,
    /// `&self`
    RefShared,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct Import {
    pub module: Option<String>,
    pub js_namespace: Option<Ident>,
    pub kind: ImportKind,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub enum ImportKind {
    Function(ImportFunction),
    Static(ImportStatic),
    Type(ImportType),
    Enum(ImportEnum),
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct ImportFunction {
    pub function: Function,
    pub rust_name: Ident,
    pub js_ret: Option<syn::Type>,
    pub catch: bool,
    pub variadic: bool,
    pub structural: bool,
    pub kind: ImportFunctionKind,
    pub shim: Ident,
    pub doc_comment: Option<String>,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub enum ImportFunctionKind {
    Method {
        class: String,
        ty: syn::Type,
        kind: MethodKind,
    },
    Normal,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub enum MethodKind {
    Constructor,
    Operation(Operation),
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct Operation {
    pub is_static: bool,
    pub kind: OperationKind,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub enum OperationKind {
    Regular,
    Getter(Option<Ident>),
    Setter(Option<Ident>),
    IndexingGetter,
    IndexingSetter,
    IndexingDeleter,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct ImportStatic {
    pub vis: syn::Visibility,
    pub ty: syn::Type,
    pub shim: Ident,
    pub rust_name: Ident,
    pub js_name: String,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct ImportType {
    pub vis: syn::Visibility,
    pub rust_name: Ident,
    pub js_name: String,
    pub attrs: Vec<syn::Attribute>,
    pub doc_comment: Option<String>,
    pub instanceof_shim: String,
    pub extends: Vec<syn::Path>,
    pub vendor_prefixes: Vec<Ident>,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct ImportEnum {
    /// The Rust enum's visibility
    pub vis: syn::Visibility,
    /// The Rust enum's identifiers
    pub name: Ident,
    /// The Rust identifiers for the variants
    pub variants: Vec<Ident>,
    /// The JS string values of the variants
    pub variant_values: Vec<String>,
    /// Attributes to apply to the Rust enum
    pub rust_attrs: Vec<syn::Attribute>,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct Function {
    pub name: String,
    pub name_span: Span,
    pub renamed_via_js_name: bool,
    pub arguments: Vec<syn::ArgCaptured>,
    pub ret: Option<syn::Type>,
    pub rust_attrs: Vec<syn::Attribute>,
    pub rust_vis: syn::Visibility,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct Struct {
    pub name: Ident,
    pub fields: Vec<StructField>,
    pub comments: Vec<String>,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct StructField {
    pub name: Ident,
    pub struct_name: Ident,
    pub readonly: bool,
    pub ty: syn::Type,
    pub getter: Ident,
    pub setter: Ident,
    pub comments: Vec<String>,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct Enum {
    pub name: Ident,
    pub variants: Vec<Variant>,
    pub comments: Vec<String>,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct Variant {
    pub name: Ident,
    pub value: u32,
}

#[derive(Copy, Clone, Debug, PartialEq, Eq)]
pub enum TypeKind {
    ByRef,
    ByMutRef,
    ByValue,
}

#[derive(Copy, Clone, Debug, PartialEq, Eq)]
pub enum TypeLocation {
    ImportArgument,
    ImportRet,
    ExportArgument,
    ExportRet,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq))]
#[derive(Clone)]
pub struct Const {
    pub vis: syn::Visibility,
    pub name: Ident,
    pub class: Option<Ident>,
    pub ty: syn::Type,
    pub value: ConstValue,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq))]
#[derive(Clone)]
/// same as webidl::ast::ConstValue
pub enum ConstValue {
    BooleanLiteral(bool),
    FloatLiteral(f64),
    SignedIntegerLiteral(i64),
    UnsignedIntegerLiteral(u64),
    Null,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct Dictionary {
    pub name: Ident,
    pub fields: Vec<DictionaryField>,
}

#[cfg_attr(feature = "extra-traits", derive(Debug, PartialEq, Eq))]
#[derive(Clone)]
pub struct DictionaryField {
    pub name: Ident,
    pub required: bool,
    pub ty: syn::Type,
}

impl Program {
    pub(crate) fn shared(&self) -> Result<shared::Program, Diagnostic> {
        let mut errors = Vec::new();
        let mut imports = Vec::new();
        for import in self.imports.iter() {
            match import.shared() {
                Ok(i) => imports.push(i),
                Err(e) => errors.push(e),
            }
        }
        Diagnostic::from_vec(errors)?;
        Ok(shared::Program {
            exports: self.exports.iter().map(|a| a.shared()).collect(),
            structs: self.structs.iter().map(|a| a.shared()).collect(),
            enums: self.enums.iter().map(|a| a.shared()).collect(),
            imports,
            version: shared::version(),
            schema_version: shared::SCHEMA_VERSION.to_string(),
        })
    }
}

impl Function {
    fn shared(&self) -> shared::Function {
        shared::Function {
            name: self.name.to_string(),
        }
    }
}

impl Export {
    /// Mangles a rust -> javascript export, so that the created Ident will be unique over function
    /// name and class name, if the function belongs to a javascript class.
    pub(crate) fn rust_symbol(&self) -> Ident {
        let mut generated_name = String::from("__wasm_bindgen_generated");
        if let Some(class) = &self.class {
            generated_name.push_str("_");
            generated_name.push_str(&class.to_string());
        }
        generated_name.push_str("_");
        generated_name.push_str(&self.function.name.to_string());
        Ident::new(&generated_name, Span::call_site())
    }

    /// This is the name of the shim function that gets exported and takes the raw
    /// ABI form of its arguments and converts them back into their normal,
    /// "high level" form before calling the actual function.
    pub(crate) fn export_name(&self) -> String {
        let fn_name = self.function.name.to_string();
        match &self.class {
            Some(class) => shared::struct_function_export_name(&class.to_string(), &fn_name),
            None => shared::free_function_export_name(&fn_name),
        }
    }

    fn shared(&self) -> shared::Export {
        let (method, consumed) = match self.method_self {
            Some(MethodSelf::ByValue) => (true, true),
            Some(_) => (true, false),
            None => (false, false),
        };
        shared::Export {
            class: self.class.as_ref().map(|s| s.to_string()),
            method,
            consumed,
            is_constructor: self.is_constructor,
            function: self.function.shared(),
            comments: self.comments.clone(),
        }
    }
}

impl Enum {
    fn shared(&self) -> shared::Enum {
        shared::Enum {
            name: self.name.to_string(),
            variants: self.variants.iter().map(|v| v.shared()).collect(),
            comments: self.comments.clone(),
        }
    }
}

impl Variant {
    fn shared(&self) -> shared::EnumVariant {
        shared::EnumVariant {
            name: self.name.to_string(),
            value: self.value,
        }
    }
}

impl Import {
    fn shared(&self) -> Result<shared::Import, Diagnostic> {
        Ok(shared::Import {
            module: self.module.clone(),
            js_namespace: self.js_namespace.as_ref().map(|s| s.to_string()),
            kind: self.kind.shared()?,
        })
    }
}

impl ImportKind {
    /// Whether this type can be inside an `impl` block.
    pub fn fits_on_impl(&self) -> bool {
        match *self {
            ImportKind::Function(_) => true,
            ImportKind::Static(_) => false,
            ImportKind::Type(_) => false,
            ImportKind::Enum(_) => false,
        }
    }

    fn shared(&self) -> Result<shared::ImportKind, Diagnostic> {
        Ok(match *self {
            ImportKind::Function(ref f) => shared::ImportKind::Function(f.shared()?),
            ImportKind::Static(ref f) => shared::ImportKind::Static(f.shared()),
            ImportKind::Type(ref f) => shared::ImportKind::Type(f.shared()),
            ImportKind::Enum(ref f) => shared::ImportKind::Enum(f.shared()),
        })
    }
}

impl ImportFunction {
    /// If the rust object has a `fn xxx(&self) -> MyType` method, get the name for a getter in
    /// javascript (in this case `xxx`, so you can write `val = obj.xxx`)
    fn infer_getter_property(&self) -> String {
        self.function.name.to_string()
    }

    /// If the rust object has a `fn set_xxx(&mut self, MyType)` style method, get the name
    /// for a setter in javascript (in this case `xxx`, so you can write `obj.xxx = val`)
    fn infer_setter_property(&self) -> Result<String, Diagnostic> {
        let name = self.function.name.to_string();

        // if `#[wasm_bindgen(js_name = "...")]` is used then that explicitly
        // because it was hand-written anyway.
        if self.function.renamed_via_js_name {
            return Ok(name);
        }

        // Otherwise we infer names based on the Rust function name.
        if !name.starts_with("set_") {
            bail_span!(
                syn::token::Pub(self.function.name_span),
                "setters must start with `set_`, found: {}",
                name,
            );
        }
        Ok(name[4..].to_string())
    }

    fn shared(&self) -> Result<shared::ImportFunction, Diagnostic> {
        let shared_operation = |operation: &Operation| -> Result<_, Diagnostic> {
            let is_static = operation.is_static;
            let kind = match &operation.kind {
                OperationKind::Regular => shared::OperationKind::Regular,
                OperationKind::Getter(g) => {
                    let g = g.as_ref().map(|g| g.to_string());
                    shared::OperationKind::Getter(g.unwrap_or_else(|| self.infer_getter_property()))
                }
                OperationKind::Setter(s) => {
                    let s = s.as_ref().map(|s| s.to_string());
                    shared::OperationKind::Setter(match s {
                        Some(s) => s,
                        None => self.infer_setter_property()?,
                    })
                }
                OperationKind::IndexingGetter => shared::OperationKind::IndexingGetter,
                OperationKind::IndexingSetter => shared::OperationKind::IndexingSetter,
                OperationKind::IndexingDeleter => shared::OperationKind::IndexingDeleter,
            };
            Ok(shared::Operation { is_static, kind })
        };

        let method = match self.kind {
            ImportFunctionKind::Method {
                ref class,
                ref kind,
                ..
            } => {
                let kind = match kind {
                    MethodKind::Constructor => shared::MethodKind::Constructor,
                    MethodKind::Operation(op) => {
                        shared::MethodKind::Operation(shared_operation(op)?)
                    }
                };
                Some(shared::MethodData {
                    class: class.clone(),
                    kind,
                })
            }
            ImportFunctionKind::Normal => None,
        };

        Ok(shared::ImportFunction {
            shim: self.shim.to_string(),
            catch: self.catch,
            variadic: self.variadic,
            method,
            structural: self.structural,
            function: self.function.shared(),
        })
    }
}

impl ImportStatic {
    fn shared(&self) -> shared::ImportStatic {
        shared::ImportStatic {
            name: self.js_name.to_string(),
            shim: self.shim.to_string(),
        }
    }
}

impl ImportType {
    fn shared(&self) -> shared::ImportType {
        shared::ImportType {
            name: self.js_name.clone(),
            instanceof_shim: self.instanceof_shim.clone(),
            vendor_prefixes: self.vendor_prefixes.iter().map(|s| s.to_string()).collect(),
        }
    }
}

impl ImportEnum {
    fn shared(&self) -> shared::ImportEnum {
        shared::ImportEnum {}
    }
}

impl Struct {
    fn shared(&self) -> shared::Struct {
        shared::Struct {
            name: self.name.to_string(),
            fields: self.fields.iter().map(|s| s.shared()).collect(),
            comments: self.comments.clone(),
        }
    }
}

impl StructField {
    fn shared(&self) -> shared::StructField {
        shared::StructField {
            name: self.name.to_string(),
            readonly: self.readonly,
            comments: self.comments.clone(),
        }
    }
}
