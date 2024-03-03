use std::path::PathBuf;
use std::{fs, path};

use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
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

impl Default for Config {
    fn default() -> Self {
        Self {
            twitch_client_id: String::new(),
            twitch_client_secret: String::new(),
            theme: Theme::Dark,
        }
    }
}

pub fn get_user_config_path() -> PathBuf {
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

pub fn get_user_config() -> Option<Config> {
    let user_config_path = get_user_config_path();

    if !user_config_path.exists() {
        fs::File::create(&user_config_path).expect("create user config failed");
    }

    let content = fs::read_to_string(&user_config_path).unwrap_or_else(|_| "".to_string());

    let data: Option<Config> = toml::from_str(&content).unwrap_or_else(|_| None);

    data
}

pub fn load_or_initial() -> Option<Config> {
    let user_config_path = get_user_config_path();

    if !user_config_path.exists() {
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

    Some(data.try_into::<Config>().expect("config data error"))
}
