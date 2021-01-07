use crate::models::movie::Movie;

use easy_http_request::HttpRequestError;

pub trait MovieService {
    fn fetch_movies(&self) -> Result<Vec<Movie>, HttpRequestError>;
    fn get_movie(&self) -> Movie;
}

