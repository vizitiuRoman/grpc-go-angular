use crate::models::movie::Movie;
use async_trait::async_trait;
use sqlx::Error;

#[async_trait]
pub trait MovieRepository {
    fn get_movie(&self) -> Movie;
    async fn get_t(&self);
}
