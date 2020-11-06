extern crate wapc_guest as guest;

use guest::prelude::*;
use serde::{Deserialize, Serialize};
use serde_json::json;

#[no_mangle]
pub extern "C" fn wapc_init() {
  register_function("hello", hello);
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
