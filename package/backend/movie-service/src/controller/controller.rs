use crate::grpc_proto::movie_grpc::MovieService;
use crate::grpc_proto::movie::{MovieStub, MoviesRes};
use crate::services::manager::Manager;
use crate::services::services::MovieService as MovieSrv;

use grpc::{ServerHandlerContext, ServerRequestSingle, ServerResponseUnarySink};

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
        let d = async_std::task::block_on(self.services.movie_service.get_movie(8587)).unwrap();
        println!("{:?}", d);
        resp.finish(r)
    }
}
