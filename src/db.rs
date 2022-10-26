use sqlx::{Pool, Sqlite, SqlitePool};

pub async fn connect() -> Result<Pool<Sqlite>, sqlx::Error> {
    SqlitePool::connect("sqlite::memory:").await
}
