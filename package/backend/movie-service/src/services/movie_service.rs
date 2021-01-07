use crate::store::store::Store;
use crate::services::services::MovieService;
use crate::models::movie::{MovieFromAPI, Movie};
use crate::store::repository::MovieRepository;

use easy_http_request::{HttpRequestError, DefaultHttpRequest};
use serde_json;

pub struct MovieSrv {
    store: Store
}

impl MovieSrv {
    pub fn new(store: Store) -> MovieSrv {
        MovieSrv {
            store
        }
    }
}

impl MovieService for MovieSrv {
    fn fetch_movies(&self) -> Result<Vec<Movie>, HttpRequestError> {
        let response = DefaultHttpRequest::get_from_url_str("https://api.themoviedb.org/3/movie/popular?api_key=8e762f584dc6993fb94182714cbc8c96&language=en-US&page=1")
            .unwrap()
            .send()
            .unwrap();
        let json: MovieFromAPI = serde_json::from_slice(&response.body[..]).unwrap();
        Ok(json.results)
    }

    fn get_movie(&self) -> Movie {
        self.store.movie_repo.get_movie()
    }
}
