CREATE TABLE links (
    linkID TEXT PRIMARY KEY,
    userID TEXT,
    channelID TEXT,
    guildID TEXT
);

CREATE OR REPLACE FUNCTION insert_link (g_linkID TEXT, g_userID TEXT, g_channelID TEXT, g_guildID TEXT)
RETURNS VOID AS $$
DECLARE r RECORD;
BEGIN
    SELECT * FROM links WHERE linkID = g_linkID INTO r;
    IF (r.userID IS NULL) THEN
        INSERT INTO links (linkID, userID, channelID, guildID) VALUES (g_linkID, g_userID, g_channelID, g_guildID);
    END IF;
END; $$ language plpgsql;
