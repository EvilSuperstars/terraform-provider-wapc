extern crate wapc_guest as guest;

use guest::prelude::*;

wapc_handler!(handle_wapc);

fn handle_wapc(operation: &str, msg: &[u8]) -> CallResult {
    match operation {
        "hello" => hello(msg),
        op => Err(format!("unknown operation {}", op).into()),
    }
}

fn hello(msg: &[u8]) -> CallResult {
    guest::console_log("[DEBUG] Enter hello");

    guest::console_log("[DEBUG] Exit hello");
    Ok(vec![])
}
