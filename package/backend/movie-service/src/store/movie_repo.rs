use crate::store::repository::MovieRepository;
use crate::models::movie::Movie;

use sqlx::PgPool;

pub struct MovieRepo {
    pool: PgPool,
}

impl MovieRepo {
    pub fn new(pool: PgPool) -> MovieRepo {
        MovieRepo { pool }
    }
}

impl MovieRepository for MovieRepo {
    fn create_movie(&self, _new_movie: Movie) -> Movie {
        unimplemented!()
    }
}
