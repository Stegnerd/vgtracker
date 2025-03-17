ALTER TABLE game_details DROP COLUMN name;
ALTER TABLE game_details DROP COLUMN thumbnail_url;
ALTER TABLE game_details DROP COLUMN cover_url;

DROP INDEX IF EXISTS idx_game_details_is_owned;
DROP INDEX IF EXISTS idx_game_details_is_beaten;
DROP INDEX IF EXISTS idx_game_details_is_wishlisted;
