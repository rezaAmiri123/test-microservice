DROP TABLE IF EXISTS articles CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;
-- CREATE EXTENSION IF NOT EXISTS postgis;
-- CREATE EXTENSION IF NOT EXISTS postgis_topology;


CREATE TABLE articles
(
    uuid         VARCHAR(128) PRIMARY KEY    NOT NULL CHECK ( uuid <> '' ),
    user_uuid    VARCHAR(128)                NOT NULL CHECK ( user_uuid <> '' ),
    title        VARCHAR(32)                 NOT NULL CHECK ( title <> '' ),
    slug         VARCHAR(64) UNIQUE          NOT NULL CHECK ( slug <> '' ),
    description  TEXT                NOT NULL,
    body         TEXT                 NOT NULL,
    about        VARCHAR(1024)                        DEFAULT '',
    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS articles_slug_id_idx ON articles (slug);
