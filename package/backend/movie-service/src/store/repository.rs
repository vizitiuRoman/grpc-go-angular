use crate::models::movie::Movie;

pub trait MovieRepository {
    fn get_movie(&self) -> Movie;
}