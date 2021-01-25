use crate::store::pg;

pub struct Store {
    pub movie_repo: pg::movie_repo::MovieRepo
}

impl Store {
    pub fn new(pool: pg::pool::PoolConnection) -> Store {
        Store {
            movie_repo: pg::movie_repo::MovieRepo::new(pool)
        }
    }
}
