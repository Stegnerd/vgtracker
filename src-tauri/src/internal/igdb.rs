use igdb_api_rust::apicalypse_builder::ApicalypseBuilder;
use igdb_api_rust::client::Client;
use igdb_api_rust::igdb::GameResult;
use serde::{Deserialize, Serialize};
use ts_rs::TS;

#[derive(Debug, Deserialize, Serialize, TS)]
#[serde(rename_all = "camelCase")]
#[ts(export, export_to = "./igdb/")]
pub struct SearchGamesInput {
    pub name_query: String,
}

#[derive(Debug, Deserialize, Serialize, TS)]
#[serde(rename_all = "camelCase")]
#[ts(export, export_to = "./igdb/")]
pub struct SearchGamesOutput {
    pub games: Vec<GameSummary>,
}

#[derive(Debug, Deserialize, Serialize, TS)]
#[serde(rename_all = "camelCase")]
#[ts(export, export_to = "./igdb/")]
pub struct GameSummary {
    pub id: u64,
}

pub async fn search_games(input: SearchGamesInput) -> Result<SearchGamesOutput, String> {
    let mapped_result: Vec<GameSummary>;

    let config = crate::internal::config::load_or_initial();
    if config.twitch_client_id.is_empty() || config.twitch_client_secret.is_empty() {
        return Err("Twitch client ID or secret not set".to_string());
    }

    let mut client = Client::new(&config.twitch_client_id, &config.twitch_client_secret);

    let mut filter_string: String = String::new();
    if !input.name_query.is_empty() {
        let name_query: String = String::from("name = ") + &input.name_query;
        filter_string.push_str(&name_query);
    }

    let query = ApicalypseBuilder::default()
        .filter(&filter_string)
        .fields("*")
        .limit(10);

    let result = client.request::<GameResult>(&query).await;

    match result {
        Ok(game_result) => {
            mapped_result = game_result
                .games
                .into_iter()
                .map(|game| GameSummary { id: game.id })
                .collect();
            Ok(SearchGamesOutput {
                games: mapped_result,
            })
        }
        Err(_) => Err("Error fetching games".to_string()),
    }
}
