use crate::store::store::Store;
use crate::services::services::MovieService;
use crate::models::movie::Movie;
use crate::store::repository::MovieRepository;

use easy_http_request::{HttpRequestError, DefaultHttpRequest};

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
    fn fetch_movies(&self) -> Result<(), HttpRequestError> {
        let response = DefaultHttpRequest::get_from_url_str("https://magiclen.org")
            .unwrap()
            .send()
            .unwrap();

        println!("{}", response.status_code);
        println!("{:?}", response.headers);
        Ok(())
    }

    fn get_movie(&self) -> Movie {
        self.store.movie_repo.get_movie()
    }
}
