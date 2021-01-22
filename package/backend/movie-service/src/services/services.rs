use crate::models::movie::Movie;

use easy_http_request::HttpRequestError;
use async_trait::async_trait;

#[async_trait]
pub trait MovieService {
    fn fetch_movies(&self) -> Result<Vec<Movie>, HttpRequestError>;
    async fn get_movie(&self) -> Movie;
}

