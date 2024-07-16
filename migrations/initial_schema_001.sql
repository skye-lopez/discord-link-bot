CREATE TABLE links (
    linkID TEXT PRIMARY KEY,
    userID TEXT,
    channelID TEXT,
    guildID TEXT,
    l TEXT
);

CREATE OR REPLACE FUNCTION insert_link (
    g_linkID TEXT, 
    g_userID TEXT, 
    g_channelID TEXT, 
    g_guildID TEXT, 
    g_link TEXT
) RETURNS VOID AS $$
DECLARE r RECORD;
BEGIN
    SELECT * FROM links WHERE linkID = g_linkID INTO r;
    IF (r.userID IS NULL) THEN
        INSERT INTO links (linkID, userID, channelID, guildID, l) VALUES (g_linkID, g_userID, g_channelID, g_guildID, g_link);
    END IF;
END; $$ language plpgsql;
