use serde::Serialize;

#[derive(Debug, Serialize)]
pub struct APIError {
    pub error: String,
}
