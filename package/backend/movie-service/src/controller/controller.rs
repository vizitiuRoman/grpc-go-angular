use crate::grpc_proto::movie_grpc::MovieService;
use crate::grpc_proto::movie::{MovieStub, MoviesRes};
use crate::services::manager::Manager;
use crate::services::services::MovieService as MovieSrv;

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
        self.services.movie_service.fetch_movies();
        resp.finish(r)
    }
}
