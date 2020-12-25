use crate::models::movie::Movie;

pub trait MovieRepository {
    fn create_movie(&self, new_movie: Movie) -> Movie;
}
