use tauri::command;

use crate::internal;
use crate::internal::config::{ReadConfigOutput, Theme, UpdateConfigInput};

#[command]
pub fn get_user_configuration() -> ReadConfigOutput {
    internal::config::read_user_configuration()
}

#[command]
pub fn update_user_configuration(input: UpdateConfigInput) -> ReadConfigOutput {
    internal::config::update_user_configuration(input)
}

#[command]
pub fn update_theme(theme_change: Theme) {
    internal::config::update_theme(theme_change);
}
