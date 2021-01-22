use sqlx::{Pool, Error, Postgres, Row, postgres};

pub type PoolConnection = Pool<Postgres>;

pub async fn create_connection_pool() -> Result<PoolConnection, Error> {
    let pool = postgres::PgPoolOptions::new()
        .max_connections(5)
        .connect("postgres://mtbbot:mtbbot@localhost:5432/mtbbot").await?;
    Ok(pool)
}
