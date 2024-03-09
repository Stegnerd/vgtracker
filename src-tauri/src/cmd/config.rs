use tauri::command;

use crate::internal;
use crate::internal::config::Config;

#[command]
pub fn get_user_config() -> Config {
    internal::config::get_user_config()
}
