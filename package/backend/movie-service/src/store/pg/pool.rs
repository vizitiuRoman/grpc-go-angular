use sqlx::{Pool, Postgres, postgres};

pub type PoolConnection = Pool<Postgres>;

pub async fn create_connection_pool() -> Result<PoolConnection, sqlx::Error> {
    let pool = postgres::PgPoolOptions::new()
        .max_connections(5)
        .connect("postgres://movie:movie@localhost:5432/movie").await?;

    sqlx::query("CREATE TABLE IF NOT EXISTS movies
        (
            id                INT8                          NOT NULL,
            backdrop_path     TEXT                          NOT NULL,
            adult             BOOLEAN                       NOT NULL,
            video             BOOLEAN                       NOT NULL,
            original_language varchar(255)                  NOT NULL,
            original_title    varchar(255)                  NOT NULL,
            title             varchar(255)                  NOT NULL,
            overview          TEXT                          NOT NULL,
            poster_path       TEXT                          NOT NULL,
            release_date      TEXT                          NOT NULL,
            popularity        FLOAT                         NOT NULL,
            vote_average      FLOAT                         NOT NULL,
            vote_count        INT8                          NOT NULL
        );
    ")
        .execute(&pool).await?;

    Ok(pool)
}
