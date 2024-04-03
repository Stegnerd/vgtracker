use std::{fs, path};
use std::path::PathBuf;

use serde::{Deserialize, Serialize};
use ts_rs::TS;

#[derive(Debug, Serialize, Deserialize, TS)]
#[ts(export, export_to = "./config/")]
pub enum Theme {
    Light,
    Dark,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct Config {
    pub twitch_client_id: String,
    pub twitch_client_secret: String,
    pub theme: Theme,
}

#[derive(Debug, Deserialize, Serialize, TS)]
#[serde(rename_all = "camelCase")]
#[ts(export, export_to = "./config/")]
pub struct ReadConfigOutput {
    pub twitch_client_id: String,
    pub twitch_client_secret: String,
    pub theme: Theme,
}

#[derive(Debug, Deserialize, Serialize, TS)]
#[serde(rename_all = "camelCase")]
#[ts(export, export_to = "./config/")]
pub struct UpdateConfigInput {
    pub twitch_client_id: String,
    pub twitch_client_secret: String,
}

impl Default for Config {
    fn default() -> Self {
        Self {
            twitch_client_id: String::new(),
            twitch_client_secret: String::new(),
            theme: Theme::Dark,
        }
    }
}

impl From<Config> for ReadConfigOutput {
    fn from(c: Config) -> Self {
        Self {
            twitch_client_id: c.twitch_client_id,
            twitch_client_secret: c.twitch_client_secret,
            theme: c.theme,
        }
    }
}

fn get_user_config_path() -> PathBuf {
    let home_dir = &tauri::api::path::home_dir().unwrap();
    let user_config = path::Path::new(home_dir);
    let user_config = user_config.join(".vgtracker");

    if !user_config.exists() {
        fs::create_dir_all(&user_config).unwrap();
    }

    let user_config = user_config.join("config.toml");

    println!("-2->{:?}", user_config);

    user_config
}

fn get_user_config() -> Config {
    let user_config_path = get_user_config_path();

    if !user_config_path.exists() {
        fs::File::create(&user_config_path).expect("create user config failed");
    }

    let content = fs::read_to_string(&user_config_path).unwrap_or_else(|_| "".to_string());

    let data: Option<Config> = toml::from_str(&content).unwrap_or_else(|_| None);

    data.expect("failed to get config")
}

pub fn read_user_config() -> ReadConfigOutput {
    let base_config = get_user_config();

    base_config.into()
}

pub fn load_or_initial() -> Option<Config> {
    let user_config_path = get_user_config_path();

    let mut write_file = false;
    if !user_config_path.exists() {
        write_file = true;
        fs::File::create(&user_config_path).expect("create user config failed");
    }

    let content = fs::read_to_string(&user_config_path).unwrap_or_else(|_| "".to_string());

    let mut data = match content.parse::<toml::Table>() {
        Ok(data) => data,
        Err(err) => {
            println!("error ==> {:?}", err);
            toml::map::Map::new()
        }
    };

    if !data.contains_key("twitch_client_id") {
        data.insert(
            String::from("twitch_client_id"),
            toml::Value::try_from::<String>(String::new()).unwrap(),
        );
    }

    if !data.contains_key("twitch_client_secret") {
        data.insert(
            String::from("twitch_client_secret"),
            toml::Value::try_from::<String>(String::new()).unwrap(),
        );
    }

    if !data.contains_key("theme") {
        data.insert(
            String::from("theme"),
            toml::Value::try_from::<Theme>(Theme::Dark).unwrap(),
        );
    }

    // fix issue with ts wanting camelCase but toml wanting snake case
    let cfg = data.try_into::<Config>().expect("config data error");

    if write_file {
        update_user_config_internal(&cfg)
    }

    Some(cfg)
}

pub fn update_user_config(update_input: UpdateConfigInput) -> ReadConfigOutput {
    let mut current_config = get_user_config();
    current_config.twitch_client_id = update_input.twitch_client_id.to_string();
    current_config.twitch_client_secret = update_input.twitch_client_secret.to_string();

    update_user_config_internal(&current_config);

    current_config.into()
}

pub fn update_theme(theme_change: Theme) {
    let mut current_config = get_user_config();
    current_config.theme = theme_change;
    update_user_config_internal(&current_config);
}

fn update_user_config_internal(cfg: &Config) {
    let user_config_path = get_user_config_path();
    let content = toml::to_string(&cfg).unwrap();
    fs::write(user_config_path, &content).expect("update config error")
}
