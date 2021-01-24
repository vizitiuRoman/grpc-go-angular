use crate::models::movie::Movie;

use async_trait::async_trait;

#[async_trait]
pub trait MovieRepository {
    async fn create_movie(&self, create_movie: Movie) -> Result<Movie, sqlx::Error>;
    async fn get_movie(&self, id: i64) -> Result<Movie, sqlx::Error>;
    async fn get_movies(&self) -> Result<Vec<Movie>, sqlx::Error>;
}
