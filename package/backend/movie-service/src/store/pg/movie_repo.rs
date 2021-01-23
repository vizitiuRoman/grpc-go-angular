use crate::store::repository::MovieRepository;
use crate::models::movie::Movie;
use crate::store::pg::pool::PoolConnection;

use sqlx::{Pool, Postgres, Row};
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
    async fn create_movie(&self, create_movie: Movie) -> Result<Movie, sqlx::Error> {
        sqlx::query_as(r#"
                INSERT INTO movies
                (
                    id,
                    backdrop_path,
                    adult,
                    video,
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
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
                RETURNING *
            "#
        )
            .bind(create_movie.id)
            .bind(create_movie.backdrop_path)
            .bind(create_movie.adult)
            .bind(create_movie.video)
            .bind(create_movie.original_language)
            .bind(create_movie.original_title)
            .bind(create_movie.title)
            .bind(create_movie.overview)
            .bind(create_movie.poster_path)
            .bind(create_movie.release_date)
            .bind(create_movie.popularity)
            .bind(create_movie.vote_average)
            .bind(create_movie.vote_count)
            .fetch_one(&self.pool)
            .await
    }

    async fn get_movie(&self, id: i64) -> Result<Movie, sqlx::Error> {
        sqlx::query_as("SELECT * FROM movies WHERE id = $1")
            .bind(id)
            .fetch_one(&self.pool)
            .await
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
