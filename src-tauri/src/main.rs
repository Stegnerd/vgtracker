// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]
use tauri::Manager;

mod cmd;
mod internal;

// Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
#[tokio::main]
async fn main() {
    tauri::Builder::default()
        .setup(move |app| {
            internal::config::load_or_initial();

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
