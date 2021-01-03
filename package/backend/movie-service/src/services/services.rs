use crate::models::movie::Movie;

pub trait MovieService {
    fn get_movie(&self) -> Movie;
}

