use crate::store::pg;

pub struct Store {
    pub movie_repo: pg::user_repo::MovieRepo
}

impl Store {
    pub fn new() -> Store {
        let poll = pg::pool::create_connection_pool();
        Store {
            movie_repo: pg::user_repo::MovieRepo::new(poll)
        }
    }
}
