use crate::models::movie::Movie;

use async_trait::async_trait;

#[async_trait]
pub trait MovieService {
    async fn synchronize_movies(&self) -> Result<(), easy_http_request::HttpRequestError>;
    async fn get_movie(&self, id: i64) -> Result<Movie, sqlx::Error>;
}

