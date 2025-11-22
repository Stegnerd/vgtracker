use serde::{Deserialize, Serialize};
use std::str;
use std::sync::Mutex;

use crate::AppState;

#[derive(Debug)]
pub struct TwitchConfiguration {
    pub client_id: String,
    pub client_secret: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct TwitchTokenResponse {
    pub access_token: String,
    pub expires_in: u64,
    pub token_type: String,
}

impl TwitchConfiguration {
    pub fn from_env() -> Self {
        Self {
            client_id: std::env::var("vgt_twitch_client_id").unwrap_or_default(),
            client_secret: std::env::var("vgt_twitch_client_secret").unwrap_or_default(),
        }
    }

    pub fn get_token(&self) -> String {
        if self.client_id.is_empty() || self.client_secret.is_empty() {
            println!("Twitch client ID or secret not found in environment variables. Twitch API calls will not work.");
            return "".to_string();
        }

        let client = reqwest::blocking::Client::new();
        let response = client
            .post("https://id.twitch.tv/oauth2/token")
            .query(&[
                ("client_id", self.client_id.as_str()),
                ("client_secret", self.client_secret.as_str()),
                ("grant_type", "client_credentials"),
            ])
            .send()
            .expect("Failed to send request to Twitch API");

        if !response.status().is_success() {
            // TODO: maybe handle this in a more graceful way?
            // maybe set to emtpy and then raise notification in the frontend?
            panic!(
                "Failed to get Twitch token: {}",
                response.text().unwrap_or_default()
            );
        }

        let json: TwitchTokenResponse = response
            .json::<TwitchTokenResponse>()
            .expect("Failed to parse Twitch token response");
        json.access_token.clone()
    }
}

#[tauri::command]
pub fn is_token_available(state: tauri::State<'_, Mutex<AppState>>) -> bool {
    let state = state.lock().unwrap();
    state.twitch_access_token != ""
}

#[tauri::command]
pub fn refresh_token(state: tauri::State<'_, Mutex<AppState>>) -> bool {
    // reload the .env file to get the latest values
    dotenvy::dotenv().expect("Failed to load .env file");

    let mut state = state.lock().unwrap();
    let twitch_config = TwitchConfiguration::from_env();
    if twitch_config.client_id.is_empty() || twitch_config.client_secret.is_empty() {
        println!("Twitch client ID or secret not found in environment variables. Twitch API calls will not work.");
        return false;
    }
    let token = twitch_config.get_token();
    if !token.is_empty() {
        state.twitch_access_token = twitch_config.get_token();
        return true;
    }

    false
}
