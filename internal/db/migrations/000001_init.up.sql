CREATE TABLE IF NOT EXISTS "songs" (
    id VARCHAR(255) NOT NULL PRIMARY KEY UNIQUE,
    group_name VARCHAR(50) NOT NULL,
    song VARCHAR(50) NOT NULL,
    release_date VARCHAR(50) NOT NULL,
    text VARCHAR(255) NOT NULL,
    link VARCHAR(255) NOT NULL ,

    CONSTRAINT unique_song UNIQUE(song, group_name)
)