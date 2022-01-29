DROP TABLE IF EXISTS articles CASCADE;
DROP TABLE IF EXISTS emails CASCADE;

CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE
EXTENSION IF NOT EXISTS CITEXT;
-- CREATE EXTENSION IF NOT EXISTS postgis;
-- CREATE EXTENSION IF NOT EXISTS postgis_topology;


CREATE TABLE articles
(
    uuid        VARCHAR(128) PRIMARY KEY NOT NULL CHECK ( uuid <> '' ),
    user_uuid   VARCHAR(128)             NOT NULL CHECK ( user_uuid <> '' ),
    title       VARCHAR(32)              NOT NULL CHECK ( title <> '' ),
    slug        VARCHAR(64) UNIQUE       NOT NULL CHECK ( slug <> '' ),
    description TEXT                     NOT NULL,
    body        TEXT                     NOT NULL,
    about       VARCHAR(1024)                     DEFAULT '',
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS articles_slug_id_idx ON articles (slug);

CREATE TABLE emails
(
    uuid       VARCHAR(128) PRIMARY KEY NOT NULL CHECK ( uuid <> '' ),
    from       VARCHAR(128)             NOT NULL CHECK ( from <> '' ),
    to         TEXT                     NOT NULL CHECK ( to <> '' ),
    subject    VARCHAR(128)             NOT NULL CHECK ( subject <> '' ),
    body       TEXT                     NOT NULL, CHECK ( body <> '' ),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

-- CREATE INDEX IF NOT EXISTS emails_uuid_id_idx ON emails (uuid);
