use sqlx::postgres::PgPoolOptions;
use sqlx::{Pool, Error, Postgres};

// pub type PoolConnection = Pool<Postgres>;
pub type PoolConnection = i8;

pub fn create_connection_pool() -> PoolConnection {
    // let pool = PgPoolOptions::new()
    //     .max_connections(5)
    //     .connect("postgres://postgres:password@localhost/test").await?;

    // Make a simple query to return the given parameter
    /* let row: (i64, ) = sqlx::query_as("SELECT $1")
         .bind(150_i64)
         .fetch_one(&pool).await?;
    */
    1
}
