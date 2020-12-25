use crate::store::repository::MovieRepository;
use crate::models::movie::Movie;

pub struct MovieService {
    movie_repo: Box<dyn MovieRepository + Send + Sync>,
}

impl MovieService {
    pub fn new(movie_repo: Box<dyn MovieRepository + Send + Sync>) -> MovieService {
        MovieService { movie_repo }
    }

    fn create_movie(&self, _new_movie: Movie) -> Movie {
        self.movie_repo.create_movie(_new_movie)
    }
}
