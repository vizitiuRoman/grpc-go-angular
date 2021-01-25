use crate::store::store::Store;
use crate::services::services::MovieService;
use crate::models::movie::{MoviesFromAPI, Movie};
use crate::store::repository::MovieRepository;

use async_trait::async_trait;

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
    async fn synchronize_movies(&self) -> Result<(), easy_http_request::HttpRequestError> {
        let mut page: u8 = 1;
        while page < 10 {
            let api_url = format!(
                r#"
            https://api.themoviedb.org/3/movie/popular?api_key=8e762f584dc6993fb94182714cbc8c96&language=en-US&page={}
        "#,
                page
            );
            let response = easy_http_request::DefaultHttpRequest::get_from_url_str(api_url)
                .unwrap()
                .send()
                .unwrap();

            match serde_json::from_slice(&response.body[..]) {
                Ok::<MoviesFromAPI, serde_json::Error>(movies_from_api) => {
                    for movie in movies_from_api.results {
                        self.store.movie_repo.create_movie(movie).await.ok();
                    }
                }
                Err(_) => {}
            }
            page += page;
        }
        Ok(())
    }

    async fn get_movie(&self, id: i64) -> Result<Movie, sqlx::Error> {
       self.store.movie_repo.get_movie(id).await
    }

    async fn get_movies(&self) -> Result<Vec<Movie>, sqlx::Error> {
       self.store.movie_repo.get_movies().await
    }
}
