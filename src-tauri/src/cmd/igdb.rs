use tauri::command;

use crate::internal;
use crate::internal::igdb::{SearchGamesInput, SearchGamesOutput};

#[command]
pub async fn search_games(input: SearchGamesInput) -> Result<SearchGamesOutput, String> {
    let result = internal::igdb::search_games(input).await;
    if let Ok(games) = result {
        Ok(games)
    } else {
        Err(result.unwrap_err())
    }
}
