ALTER TABLE game_details ADD COLUMN name TEXT DEFAULT '';
ALTER TABLE game_details ADD COLUMN thumbnail_url TEXT DEFAULT '';
ALTER TABLE game_details ADD COLUMN cover_url TEXT DEFAULT '';

CREATE INDEX idx_game_details_is_owned ON game_details(is_owned);
CREATE INDEX idx_game_details_is_beaten ON game_details(is_beaten);
CREATE INDEX idx_game_details_is_wishlisted ON game_details(is_wishlisted);
