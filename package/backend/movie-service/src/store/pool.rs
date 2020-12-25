use sqlx::postgres::PgPoolOptions;
use sqlx::Error;

pub async fn create_connection_pool() -> Result<i64, Error> {
    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect("postgres://postgres:password@localhost/test").await?;

    // Make a simple query to return the given parameter
    let row: (i64, ) = sqlx::query_as("SELECT $1")
        .bind(150_i64)
        .fetch_one(&pool).await?;

    Ok(1)
}
