use super::handler;
use actix_web::web;

pub fn router() -> actix_web::Scope {
    web::scope("/one_word")
        .service(handler::new_game)
        .service(handler::attempt)
}
