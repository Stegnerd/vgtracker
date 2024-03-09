// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use tokio;

mod cmd;
mod internal;

// Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
#[tauri::command]
fn greet(name: &str) -> String {
    format!("Hello, {}! You've been greeted from Rust!", name)
}

#[tokio::main]
async fn main() {
    tauri::Builder::default()
        .setup(move |app| {
            internal::config::load_or_initial();

            Ok(())
        })
        .invoke_handler(tauri::generate_handler![
            greet,
            cmd::config::get_user_config
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
