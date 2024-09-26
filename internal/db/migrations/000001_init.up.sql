CREATE TABLE IF NOT EXISTS "songs" (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    group_name VARCHAR(50) NOT NULL,
    song VARCHAR(50) NOT NULL,
    release_date VARCHAR(50) NOT NULL DEFAULT 'Not found',
    text VARCHAR(255) NOT NULL DEFAULT 'Not found',
    patronymic VARCHAR(255) NOT NULL DEFAULT 'Not found'
)