extern crate wapc_guest as guest;

use guest::prelude::*;
use serde_json::json;

wapc_handler!(handle_wapc);

fn handle_wapc(operation: &str, msg: &[u8]) -> CallResult {
    match operation {
        "hello" => hello(msg),
        op => Err(format!("unknown operation {}", op).into()),
    }
}

fn hello(msg: &[u8]) -> CallResult {
    guest::console_log("[DEBUG] Enter hello");

    let t = match serde_json::from_slice(msg)? {
        serde_json::Value::Array(_) => "Array",
        serde_json::Value::Bool(_) => "Bool",
        serde_json::Value::Null => "Null",
        serde_json::Value::Number(_) => "Number",
        serde_json::Value::Object(_) => "Object",
        serde_json::Value::String(_) => "String",
    };
    guest::console_log(&format!("[DEBUG] JSON {}", t));

    let result = json!({"key": 123});

    guest::console_log("[DEBUG] Exit hello");
    Ok(serde_json::to_vec(&result)?)
}
