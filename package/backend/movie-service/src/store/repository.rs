use crate::models::movie::Movie;
use async_trait::async_trait;
use sqlx::Error;

#[async_trait]
pub trait MovieRepository {
    async fn create_movie(&self, create_movie: Movie) -> Result<(), Error>;
    fn get_movie(&self) -> Movie;
    async fn get_t(&self);
}
