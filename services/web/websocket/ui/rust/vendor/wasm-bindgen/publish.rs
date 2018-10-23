//! Helper script to publish the wasm-bindgen suite of crates
//!
//! Usage:
//!
//! * First, compile this script
//! * Next, set cwd to the root of the wasm-bindgen repository
//! * Execute `./publish bump` to bump versions
//! * Send a PR
//! * Merge when green
//! * Execute `./publish publish` to publish crates

use std::env;
use std::fs;
use std::path::{Path, PathBuf};
use std::process::Command;

const CRATES_TO_PUBLISH: &[&str] = &[
    "wasm-bindgen",
    "js-sys",
    "wasm-bindgen-backend",
    "wasm-bindgen-cli",
    "wasm-bindgen-cli-support",
    "wasm-bindgen-futures",
    "wasm-bindgen-gc",
    "wasm-bindgen-macro",
    "wasm-bindgen-macro-support",
    "wasm-bindgen-shared",
    "wasm-bindgen-test",
    "wasm-bindgen-test-macro",
    "wasm-bindgen-wasm-interpreter",
    "wasm-bindgen-webidl",
    "web-sys",
];

const CRATES_TO_AVOID_PUBLISH: &[&str] = &[
    // We'll publish these when they're ready one day
    "wasm-bindgen-typescript",

    // These are internal crates, unlikely to ever be published
    "ui-tests",
    "sample",
    "webidl-tests",
];

struct Crate {
    manifest: PathBuf,
    name: String,
    version: String,
    next_version: String,
}

fn main() {
    let mut crates = Vec::new();
    crates.push(read_crate("./Cargo.toml".as_ref()));
    find_crates("crates".as_ref(), &mut crates);

    match &env::args().nth(1).expect("must have one argument")[..] {
        "bump" => {
            for krate in crates.iter() {
                bump_version(&krate, &crates);
            }
        }

        "publish" => {
            for krate in crates.iter() {
                publish(&krate);
            }
        }

        s => panic!("unknown command: {}", s),
    }
}

fn find_crates(dir: &Path, dst: &mut Vec<Crate>) {
    if dir.join("Cargo.toml").exists() {
        let krate = read_crate(&dir.join("Cargo.toml"));
        if CRATES_TO_PUBLISH.iter()
            .chain(CRATES_TO_AVOID_PUBLISH)
            .any(|c| krate.name == *c)
        {
            dst.push(krate);
        } else {
            panic!("failed to find {:?} in whitelist or blacklist", krate.name);
        }
    }

    for entry in dir.read_dir().unwrap() {
        let entry = entry.unwrap();
        if entry.file_type().unwrap().is_dir() {
            find_crates(&entry.path(), dst);
        }
    }
}

fn read_crate(manifest: &Path) -> Crate {
    let mut name = None;
    let mut version = None;
    for line in fs::read_to_string(manifest).unwrap().lines() {
        if name.is_none() && line.starts_with("name = \"") {
            name = Some(line.replace("name = \"", "")
                .replace("\"", "")
                .trim()
                .to_string());
        }
        if version.is_none() && line.starts_with("version = \"") {
            version = Some(line.replace("version = \"", "")
                .replace("\"", "")
                .trim()
                .to_string());
        }
    }
    let name = name.unwrap();
    let version = version.unwrap();
    let next_version = if CRATES_TO_PUBLISH.contains(&&name[..]) {
        bump(&version)
    } else {
        version.clone()
    };
    Crate {
        manifest: manifest.to_path_buf(),
        name,
        version,
        next_version,
    }
}

fn bump_version(krate: &Crate, crates: &[Crate]) {
    let contents = fs::read_to_string(&krate.manifest).unwrap();

    let mut new_manifest = String::new();
    let mut is_deps = false;
    for line in contents.lines() {
        let mut rewritten = false;
        if line.starts_with("version =") {
            if CRATES_TO_PUBLISH.contains(&&krate.name[..]) {
                println!("bump `{}` {} => {}",
                         krate.name,
                         krate.version,
                         krate.next_version);
                new_manifest.push_str(&line.replace(&krate.version, &krate.next_version));
                rewritten = true;
            }
        }

        is_deps = if line.starts_with("[") {
            line.contains("dependencies")
        } else {
            is_deps
        };

        for other in crates {
            if !is_deps || !line.starts_with(&format!("{} ", other.name)) {
                continue
            }
            if !line.contains(&other.version) {
                if !line.contains("version =") {
                    continue
                }
                panic!("{:?} has a dep on {} but doesn't list version {}",
                       krate.manifest,
                       other.name,
                       other.version);
            }
            rewritten = true;
            new_manifest.push_str(&line.replace(&other.version, &other.next_version));
            break
        }
        if !rewritten {
            new_manifest.push_str(line);
        }
        new_manifest.push_str("\n");
    }
    fs::write(&krate.manifest, new_manifest).unwrap();
}

fn bump(version: &str) -> String {
    let mut iter = version.split('.').map(|s| s.parse::<u32>().unwrap());
    let major = iter.next().expect("major version");
    let minor = iter.next().expect("minor version");
    let patch = iter.next().expect("patch version");
    format!("{}.{}.{}", major, minor, patch + 1)
}

fn publish(krate: &Crate) {
    if !CRATES_TO_PUBLISH.iter().any(|s| *s == krate.name) {
        return
    }
    let status = Command::new("cargo")
        .arg("publish")
        .current_dir(krate.manifest.parent().unwrap())
        .arg("--no-verify")
        .arg("--allow-dirty")
        .status()
        .expect("failed to run cargo");
    if !status.success() {
        println!("FAIL: failed to publish `{}`: {}", krate.name, status);
    }
}
