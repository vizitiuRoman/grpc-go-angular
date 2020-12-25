// use crate::schema::movies;

// #[derive(Identifiable, Queryable, Insertable)]
// #[table_name = "movies"]
pub struct Movie {
    pub adult: bool,
    pub backdrop_path: String,
    pub genre_ids: Vec<i64>,
    pub id: i64,
    pub original_language: String,
    pub original_title: String,
    pub overview: String,
    pub popularity: f64,
    pub poster_path: String,
    pub release_date: String,
    pub title: String,
    pub video: bool,
    pub vote_average: f64,
    pub vote_count: i64,
}
