use tauri::command;

use crate::internal;
use crate::internal::config::ReadConfigOutput;

#[command]
pub fn get_user_config() -> ReadConfigOutput {
    internal::config::get_user_config()
}
