use crate::store::store::Store;
use crate::services::services::MovieService;
use crate::models::movie::Movie;
use crate::store::repository::MovieRepository;

pub struct MovieSrv {
    store: Store
}

impl MovieSrv {
    pub fn new(store: Store) -> MovieSrv {
        MovieSrv {
            store
        }
    }
}

impl MovieService for MovieSrv {
    fn get_movie(&self) -> Movie {
        self.store.movie_repo.get_movie()
    }
}
