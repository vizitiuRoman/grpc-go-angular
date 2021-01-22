use crate::store::store::Store;
use crate::services::services::MovieService;
use crate::models::movie::{MoviesFromAPI, Movie};
use crate::store::repository::MovieRepository;

use easy_http_request::{HttpRequestError, DefaultHttpRequest};
use async_trait::async_trait;
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

#[async_trait]
impl MovieService for MovieSrv {
    fn fetch_movies(&self) -> Result<Vec<Movie>, HttpRequestError> {
        let response = DefaultHttpRequest::get_from_url_str("https://api.themoviedb.org/3/movie/popular?api_key=8e762f584dc6993fb94182714cbc8c96&language=en-US&page=1")
            .unwrap()
            .send()
            .unwrap();
        let movies_from_api: MoviesFromAPI = serde_json::from_slice(&response.body[..]).unwrap();
        Ok(movies_from_api.results)
    }

    async fn get_movie(&self) -> Movie {
        let c = self.store.movie_repo.create_movie(Movie {
            adult: false,
            backdrop_path: "fefewf".to_string(),
            genre_ids: vec![1, 2, 3],
            id: 10,
            original_language: "123123123".to_string(),
            original_title: "123123123".to_string(),
            overview: "123123123".to_string(),
            popularity: 11.0,
            poster_path: "123123123".to_string(),
            release_date: "123123123".to_string(),
            title: "123123123".to_string(),
            video: false,
            vote_average: 11.0,
            vote_count: 11,
        }).await.unwrap();
        println!("Good");
        self.store.movie_repo.get_movie()
    }
}
