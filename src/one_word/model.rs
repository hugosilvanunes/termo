use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, sqlx::FromRow)]
pub struct Word {
    pub name: String,
}

#[derive(Debug, Serialize)]
pub struct WordLength {
    pub length: usize,
}

#[derive(Debug, Deserialize)]
pub struct AttemptRequest {
    pub attempt: String,
}

#[derive(Debug, Serialize)]
pub struct AttemptCharResponse {
    pub index: usize,
    pub char: char,
    pub status: AttemptStatus,
}

#[derive(Debug, Serialize)]
pub enum AttemptStatus {
    CORRECT,
    INCORRECT,
    SEMICORRECT,
}
