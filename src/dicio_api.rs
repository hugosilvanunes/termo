use serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct WordResponse {
    pub id: i64,
    pub word: String,
    pub count: u8,
    pub character: char,
}
