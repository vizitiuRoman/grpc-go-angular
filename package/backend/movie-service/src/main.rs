mod services;
mod models;
mod controller;
mod store;
mod grpc_proto;

use crate::controller::controller::Controller;
use crate::services::manager::Manager;
use crate::store::store::Store;
use crate::grpc_proto::movie_grpc::MovieServiceServer;

#[async_std::main]
async fn main() {
    let port = 9092;
    // creating server
    let mut server = grpc::ServerBuilder::new_plain();
    // adding port to server for http
    server.http.set_port(port);
    // adding say service to server
    let store = Store::new();
    let manager = Manager::new(store);
    let controller = Controller::new(manager);
    server.add_service(MovieServiceServer::new_service_def(controller));
    // running the server
    let _server = server.build().expect("server");
    println!(
        "greeter server started on port {}",
        port,
    );
    // stopping the program from finishing
    loop {
        std::thread::park();
    }
}
