// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use std::sync::Mutex;
use igdb_api_rust::client::Client;
use tauri::{Manager, State};

mod cmd;
mod internal;

pub struct AppContext {
    igdb_client: Mutex<Client>,
}

// Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
#[tokio::main]
async fn main() {
    tauri::Builder::default()
        .setup(move |app| {
            app.manage(AppContext{ igdb_client: Mutex::new(Client::new("","")) });

            let setup_config = internal::config::load_or_initial();

            if !setup_config.twitch_client_id.is_empty() && !setup_config.twitch_client_secret.is_empty() {
                let app_state: State<AppContext> = app.state();
                let mut client = app_state.igdb_client.lock().unwrap();
                *client= Client::new(&setup_config.twitch_client_id, &setup_config.twitch_client_secret);
            }

            #[cfg(debug_assertions)] // only include this code on debug builds
            {
                let window = app.get_window("main").unwrap();
                window.open_devtools();
            }

            Ok(())
        })
        .invoke_handler(tauri::generate_handler![
            cmd::config::get_user_configuration,
            cmd::config::update_user_configuration,
            cmd::config::update_theme,
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
