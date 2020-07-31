fn main() {
    prost_build::compile_protos(&["src/tfplugin5/tfplugin5.2.proto"], &["src/"]).unwrap();
}
