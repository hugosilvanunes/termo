use super::model::{Word, WordLength};
use crate::{
    error,
    one_word::model::{AttemptCharResponse, AttemptRequest, AttemptStatus},
    state::AppState,
};
use actix_web::{get, post, web, HttpResponse};

#[get("")]
pub async fn new_game(data: web::Data<AppState>) -> HttpResponse {
    let word = sqlx::query_as::<_, Word>("SELECT * FROM words")
        .fetch_one(&data.db)
        .await;

    match word {
        Ok(value) => HttpResponse::Ok().json(WordLength {
            length: value.name.len(),
        }),
        Err(err) => HttpResponse::InternalServerError().json(error::APIError {
            error: err.to_string(),
        }),
    }
}

#[post("/attempt")]
pub async fn attempt(
    data: web::Data<AppState>,
    attempt: web::Json<AttemptRequest>,
) -> HttpResponse {
    let attempt = attempt.attempt.trim();
    if attempt.is_empty() {
        return HttpResponse::InternalServerError().json(error::APIError {
            error: "attempt is empty".to_string(),
        });
    }

    let word = match sqlx::query_as::<_, Word>("SELECT * FROM words")
        .fetch_one(&data.db)
        .await
    {
        Ok(w) => w,
        Err(err) => {
            return HttpResponse::InternalServerError().json(error::APIError {
                error: err.to_string(),
            })
        }
    };

    if attempt.len() != word.name.len() {
        return HttpResponse::InternalServerError().json(error::APIError {
            error: "length is not equals".to_string(),
        });
    }

    let mut res: Vec<AttemptCharResponse> = Vec::new();

    // for (i, v) in word.name.chars().enumerate() {
    // if v != attempt_chars[i] {
    // res.push(AttemptCharResponse {
    // index: i,
    // char: '-',
    // });
    // continue;
    // }
    //
    // res.push(AttemptCharResponse { index: i, char: v });
    // }

    for (k, j) in attempt.chars().enumerate() {
        let mut is_incorrect = true;
        for (i, v) in word.name.chars().enumerate() {
            if v == j {
                if i == k {
                    res.push(AttemptCharResponse {
                        index: k,
                        char: v,
                        status: AttemptStatus::CORRECT,
                    });
                    is_incorrect = false;
                    break;
                }

                res.push(AttemptCharResponse {
                    index: k,
                    char: v,
                    status: AttemptStatus::SEMICORRECT,
                });
                is_incorrect = false;
                break;
            }
        }

        if is_incorrect {
            res.push(AttemptCharResponse {
                index: k,
                char: j,
                status: AttemptStatus::INCORRECT,
            });
        }
    }

    HttpResponse::Ok().json(res)
}
