use crate::store::repository::MovieRepository;
use crate::models::movie::Movie;
use crate::store::pg::pool::PoolConnection;

use sqlx::{Pool, Postgres, Error, Row};
use async_trait::async_trait;

pub struct MovieRepo {
    pool: PoolConnection
}

impl MovieRepo {
    pub fn new(pool: PoolConnection) -> MovieRepo {
        MovieRepo {
            pool
        }
    }
}

#[async_trait]
impl MovieRepository for MovieRepo {
    async fn create_movie(&self, create_movie: Movie) -> Result<(), Error> {
        let movies = sqlx::query(r#"
                INSERT INTO movies
                (
                    id,
                    backdrop_path,
                    adult,
                    video,
                    genre_ids,
                    original_language,
                    original_title,
                    title,
                    overview,
                    poster_path,
                    release_date,
                    popularity,
                    vote_average,
                    vote_count
                )
                VALUES (1, 'path', true, true, '{1, 2, 3}', 'ru', 'title ru', 'title', 'overview', 'poster path', 'release date', 1.22, 12.44, 11.9)
                RETURNING *
            "#,)
            .execute(&self.pool)
            .await?;
        Ok(())
    }

    fn get_movie(&self) -> Movie {
        Movie {
            adult: false,
            backdrop_path: "".to_string(),
            genre_ids: vec![],
            id: 0,
            original_language: "".to_string(),
            original_title: "".to_string(),
            overview: "".to_string(),
            popularity: 0.0,
            poster_path: "".to_string(),
            release_date: "".to_string(),
            title: "".to_string(),
            video: false,
            vote_average: 0.0,
            vote_count: 0,
        }
    }

    async fn get_t(&self) {
        let rows = sqlx::query("SELECT * from users")
            .fetch_all(&self.pool).await.unwrap();
        for row in rows {
            let username: String = row.get("login"); // username
            println!("{}", username)
        }
    }
}
