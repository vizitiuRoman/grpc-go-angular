use crate::grpc_proto::movie_grpc::MovieService;
use crate::grpc_proto::movie::{MovieStub, MoviesRes};
use crate::services::manager::Manager;

use grpc::{ServerHandlerContext, ServerRequestSingle, ServerResponseUnarySink};
use protobuf::RepeatedField;

pub struct Controller {
    services: Manager
}

impl Controller {
    pub fn new(manager: Manager) -> Controller {
        Controller {
            services: manager,
        }
    }
}

impl MovieService for Controller {
    fn get_movies(&self, o: ServerHandlerContext, req: ServerRequestSingle<MovieStub>, resp: ServerResponseUnarySink<MoviesRes>) -> grpc::Result<()> {
        let mut r = MoviesRes::new();
        resp.finish(r)
    }
}
