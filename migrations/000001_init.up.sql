-- Qip Migration file

-- Instance
CREATE TABLE IF NOT EXISTS "instances"
(
    id              varchar PRIMARY KEY NOT NULL,
    name            varchar             NOT NULL,
    softwareName    varchar             NOT NULL,
    softwareVersion varchar             NOT NULL,
    host            varchar             NOT NULL UNIQUE,
    description     varchar             NOT NULL,
    state           int                 NOT NULL,
    createdAt       timestamptz         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt       timestamptz
);

-- File
CREATE TABLE IF NOT EXISTS "files"
(
    id           varchar PRIMARY KEY NOT NULL,
    fileName     varchar             NOT NULL,
    filePath     varchar,
    fileURL      varchar             NOT NULL,
    thumbnailURL varchar,
    blurhash     varchar             NOT NULL,
    isNSFW       boolean             NOT NULL,
    mimeType     varchar             NOT NULL,
    createdAt    timestamptz         NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- User
CREATE TABLE IF NOT EXISTS "users"
(
    id            varchar PRIMARY KEY NOT NULL,
    name          varchar             NOT NULL,
    displayName   varchar             NOT NULL,
    role          int                 NOT NULL,
    bio           varchar,
    isFroze       boolean             NOT NULL,
    inboxURL      varchar             NOT NULL,
    outboxURL     varchar             NOT NULL,
    followURL     varchar             NOT NULL,
    followersURL  varchar             NOT NULL,
    secretKey     varchar,
    publicKey     varchar             NOT NULL,
    password      varchar             NOT NULL,
    isLocalUser   boolean             NOT NULL,
    createdAt     timestamptz         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt     timestamptz,


    instanceID    varchar             NOT NULL,
    FOREIGN KEY (instanceID) REFERENCES instances (id),
    headerImageID varchar,
    iconImageID   varchar
);

-- Post
CREATE TABLE IF NOT EXISTS "posts"
(
    id         varchar PRIMARY KEY NOT NULL,
    body       text                NOT NULL,
    visibility int                 NOT NULL,
    createdAt  timestamptz         NOT NULL DEFAULT CURRENT_TIMESTAMP,

    authorID   varchar             NOT NULL,
    FOREIGN KEY (authorID) REFERENCES users (id)
);


