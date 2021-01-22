use crate::store::pg;
use crate::store::pg::pool::PoolConnection;

pub struct Store {
    pub movie_repo: pg::movie_repo::MovieRepo
}

impl Store {
    pub fn new(pool: PoolConnection) -> Store {
        Store {
            movie_repo: pg::movie_repo::MovieRepo::new(pool)
        }
    }
}
