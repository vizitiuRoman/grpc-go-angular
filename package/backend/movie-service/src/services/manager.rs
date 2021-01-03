use crate::services::movie_service::MovieSrv;
use crate::store::store::Store;

pub struct Manager {
    pub user_service: MovieSrv
}

impl Manager {
    pub fn new(store: Store) -> Manager {
        Manager {
            user_service: MovieSrv::new(store)
        }
    }
}
