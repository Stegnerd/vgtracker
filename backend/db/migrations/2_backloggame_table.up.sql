CREATE TABLE backlog_game
(
    game_id           INTEGER PRIMARY KEY AUTOINCREMENT,
    name              VARCHAR(255) NOT NULL DEFAULT '',
    finished          BOOLEAN               DEFAULT 0,
    finished_on       DATE                  DEFAULT CURRENT_TIMESTAMP,
    platform_owned_on CHAR(20)     NOT NULL DEFAULT ''
);

