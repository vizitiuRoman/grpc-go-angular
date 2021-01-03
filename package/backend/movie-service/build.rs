fn main() {
    // compile protocol buffer using protoc
    protoc_rust_grpc::Codegen::new()
        .out_dir("src/grpc_proto")
        .input("../../../grpc-proto/movie/movie.proto")
        .include("../../../grpc-proto")
        .rust_protobuf(true)
        .run()
        .expect("error compiling protocol buffer");
}
