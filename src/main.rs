use actix_web::{middleware, web, App, HttpServer};
use dicio_api::WordResponse;

mod config;
mod db;
mod dicio_api;
mod error;
mod one_word;
mod state;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init_from_env(env_logger::Env::new().default_filter_or("info"));

    let cfg = envy::from_env::<config::Config>().expect("cannot read config from env");

    let pool = db::connect().await.expect("cannot connect to database");

    let mut tx = pool.begin().await.expect("cannot init the transaction");

    sqlx::query("CREATE TABLE IF NOT EXISTS words(id INTEGER PRIMARY KEY, name VARCHAR(20))")
        .execute(&mut tx)
        .await
        .expect("migrations failed");

    tx.commit().await.expect("cannot commit the transaction");

    let resp = reqwest::get(cfg.dicio_api_url)
        .await
        .unwrap()
        .json::<WordResponse>()
        .await
        .unwrap();

    println!("{:#?}", resp);

    sqlx::query("INSERT INTO words (name) VALUES ($1)")
        .bind(resp.word)
        .execute(&pool)
        .await
        .unwrap();

    let app_state = web::Data::new(state::AppState { db: pool });

    log::info!("starting HTTP server at http://localhost:8080");

    HttpServer::new(move || {
        App::new()
            .wrap(middleware::Compress::default())
            .wrap(middleware::Logger::default())
            .app_data(web::JsonConfig::default().limit(4096)) // <- limit size of the payload (global configuration)
            .app_data(app_state.clone())
            .service(one_word::router::router())
    })
    .bind(("127.0.0.1", cfg.port))?
    .run()
    .await
}
