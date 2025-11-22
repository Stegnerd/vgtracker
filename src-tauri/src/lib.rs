use anyhow::{Context, Result};
use std::{env, sync::Mutex};
use tauri::Manager;
pub mod twitch;
use twitch::{is_token_available, refresh_token};

// Learn more about Tauri commands at https://tauri.app/develop/calling-rust/
#[tauri::command]
fn greet(name: &str) -> String {
    format!("Hello, {}! You've been greeted from Rust!", name)
}

#[derive(Debug)]
pub struct AppState {
    pub twitch_access_token: String,
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() -> Result<()> {
    dotenvy::dotenv().with_context(|| "Failed to load .env file on start up")?;

    tauri::Builder::default()
        // .plugin(tauri_plugin_dialog::init())
        // .plugin(tauri_plugin_notification::init())
        .setup(|app| {
            #[cfg(debug_assertions)]
            {
                let window = app.get_webview_window("main").unwrap();
                window.open_devtools();
                window.close_devtools();
            }

            let mut state = AppState {
                twitch_access_token: "".to_string(),
            };

            let twitch_config = twitch::TwitchConfiguration::from_env();
            if !twitch_config.client_id.is_empty() && !twitch_config.client_secret.is_empty() {
                state.twitch_access_token = twitch_config.get_token();
            }

            app.manage(Mutex::new(state));

            Ok(())
        })
        .plugin(tauri_plugin_opener::init())
        .invoke_handler(tauri::generate_handler![
            greet,
            is_token_available,
            refresh_token
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");

    Ok(())
}
