use crate::services::movie_service::MovieSrv;
use crate::store::store::Store;

pub struct Manager {
    pub movie_service: MovieSrv
}

impl Manager {
    pub fn new(store: Store) -> Manager {
        Manager {
            movie_service: MovieSrv::new(store)
        }
    }
}
