extern crate wapc_guest as guest;

use guest::prelude::*;
use serde::{Deserialize, Serialize};
use serde_json::json;

wapc_handler!(handle_wapc);

fn handle_wapc(operation: &str, msg: &[u8]) -> CallResult {
    match operation {
        "hello" => hello(msg),
        op => Err(format!("unknown operation {}", op).into()),
    }
}

//
// Can be specified in HCL as:
// input = {
//   "Name"      = "waPC"
//   "Uppercase" = true
// }
//
#[derive(Serialize, Deserialize)]
struct Input {
    #[serde(rename = "Name")]
    name: String,
    #[serde(rename = "Uppercase")]
    #[serde(default)]
    uppercase: bool,
}

fn hello(msg: &[u8]) -> CallResult {
    guest::console_log("[DEBUG] Enter hello");

    let result = match serde_json::from_slice(msg) {
        Ok(input @ Input { .. }) => {
            let mut name = input.name;
            if input.uppercase {
                name = name.to_uppercase();
            }
            //
            // Can be referenced in HCL as `result["Name"]`.
            //
            serde_json::to_vec(&json!({ "Name": name }))?
        }
        _ => vec![],
    };

    guest::console_log("[DEBUG] Exit hello");
    Ok(result)
}
