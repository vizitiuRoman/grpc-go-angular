use crate::grpc_proto::movie_grpc::MovieService;
use crate::grpc_proto::movie::{MovieStub, MoviesRes, MovieRes};
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
    fn get_movies(&self, _: ServerHandlerContext, _: ServerRequestSingle<MovieStub>, resp: ServerResponseUnarySink<MoviesRes>) -> grpc::Result<()> {
        let mut r = MoviesRes::new();
        let movies = async_std::task::block_on(self.services.movie_service.get_movies()).unwrap();
        let movie_res = movies.into_iter().map(|movie| MovieRes{
            id: movie.id,
            adult: movie.adult,
            backdrop_path: movie.backdrop_path,
            original_language: movie.original_language,
            original_title: movie.original_title,
            overview: movie.overview,
            poster_path: movie.poster_path,
            release_date: movie.release_date,
            title: movie.title,
            video: movie.video,
            popularity: movie.popularity,
            vote_average: movie.vote_average,
            vote_count: movie.vote_count,
            unknown_fields: Default::default(),
            cached_size: Default::default()
        }).collect();
        r.set_movies(movie_res);
        resp.finish(r)
    }
}
