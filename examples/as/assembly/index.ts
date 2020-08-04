// waPC guest package.
import { consoleLog, handleAbort, handleCall, register } from "wapc-guest-as";

import { JSON } from "assemblyscript-json";

// This must be present in the entry file to be exported from the Wasm module.
// It's invoked by the waPC host.
export function __guest_call(operation_size: usize, payload_size: usize): bool {
  return handleCall(operation_size, payload_size);
}

// Abort function.
// Compile with '--use abort=path/to/module/abort'.
function abort(message: string | null, fileName: string | null, lineNumber: u32, columnNumber: u32): void {
  handleAbort(message, fileName, lineNumber, columnNumber);
}

// Explicit start function to call.
export function _start(): void {
  register("hello", hello);
}

// waPC invokable operation.
function hello(payload: ArrayBuffer): ArrayBuffer {
  consoleLog("[DEBUG] Enter hello");

  const input = String.UTF8.decode(payload);
  let jsonObj: JSON.Value = JSON.parse(input);

  if (jsonObj instanceof JSON.Arr) {
    consoleLog("[DEBUG] JSON.Arr");
  } else if (jsonObj instanceof JSON.Str) {
    consoleLog("[DEBUG] JSON.Str");
  } else {
    consoleLog("[DEBUG] Unknown");
  }

  consoleLog("[DEBUG] " + input);

  consoleLog("[DEBUG] Exit hello");
  return String.UTF8.encode(JSON.from("Hello").toString());
}
