mod services;
mod models;
mod controller;
mod store;
mod grpc_proto;

use crate::controller::controller::Controller;
use crate::services::manager::Manager;
use crate::store::store::Store;
use crate::grpc_proto::movie_grpc::MovieServiceServer;
use crate::store::pg::pool::{create_connection_pool};

#[async_std::main]
async fn main() {
    let port = 9092;
    // creating server
    let mut server = grpc::ServerBuilder::new_plain();
    server.http.set_port(port);

    let pool = create_connection_pool().await.unwrap();
    let store = Store::new(pool);
    let manager = Manager::new(store);
    let controller = Controller::new(manager);

    server.add_service(MovieServiceServer::new_service_def(controller));
    let _server = server.build().expect("server");
    println!(
        "Service started on port: {}",
        port,
    );

    loop {
        std::thread::park();
    }
}
