fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_client(false)
        .build_server(true)
        .compile(&["tfplugin5/tfplugin5.2.proto"], &["tfplugin5"])?;
    Ok(())
}
