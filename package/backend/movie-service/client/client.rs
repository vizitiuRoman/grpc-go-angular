#[path = "../src/grpc_proto/mod.rs"]
mod grpc_proto;

#[path = "../src/grpc_proto/movie.rs"]
mod movie;

#[path = "../src/grpc_proto/movie_grpc.rs"]
mod movie_grpc;

extern crate futures;

use crate::movie_grpc::MovieServiceClient;

use grpc::ClientStubExt;
use movie::{MovieStub};
use futures::executor;

fn main() {
    let port = 50051;
    let client_conf = Default::default();
    // create a client
    let client = MovieServiceClient::new_plain("::1", port, client_conf).unwrap();
    // create request
    let req = MovieStub::new();
    // send the request
    let resp = client
        .get_movies(grpc::RequestOptions::new(), req)
        .join_metadata_result();
    // wait for response
    println!("{:?}", executor::block_on(resp))
}
