-- ユーザー
CREATE TABLE IF NOT EXISTS users
(
    id             varchar PRIMARY KEY NOT NULL,
    host           varchar,
    name           varchar             NOT NULL,
    password       varchar,
    screenName     varchar             NOT NULL,
    summary        varchar             NOT NULL,
    privateKey     varchar             NOT NULL,
    publicKey      varchar             NOT NULL,
    watcherCount   int                 NOT NULL,
    postsCount     int                 NOT NULL,
    headerImageUrl varchar,
    iconImageUrl   varchar,

    createdAt      timestamptz         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt      timestamptz
);

-- 投稿
CREATE TABLE IF NOT EXISTS posts
(
    id         varchar PRIMARY KEY NOT NULL,
    body       text                NOT NULL,
    type       varchar             NOT NULL,
    mergeCount int                 NOT NULL,
    visibility varchar             NOT NULL,
    createdAt  timestamptz         NOT NULL DEFAULT CURRENT_TIMESTAMP,

    userID     varchar             NOT NULL,
    FOREIGN KEY (userID) REFERENCES users (id)
);

-- リアクション
CREATE TABLE IF NOT EXISTS reactions
(
    reactedUserID varchar     NOT NULL,
    postID        varchar     NOT NULL,
    createdAt     timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (reactedUserID) REFERENCES users (id),
    FOREIGN KEY (postID) REFERENCES posts (id)
);

-- ファイル
CREATE TABLE IF NOT EXISTS files
(
    ID        varchar PRIMARY KEY NOT NULL,
    host      varchar             NOT NULL,
    md5Hash   varchar             NOT NULL,
    mimeType  varchar             NOT NULL,
    fileSize  integer             NOT NULL,
    url       varchar             NOT NULL,
    isNSFW    boolean             NOT NULL,
    blurHash  varchar             NOT NULL,
    createdAt timestamptz         NOT NULL DEFAULT CURRENT_TIMESTAMP,

    userID    varchar             NOT NULL,
    FOREIGN KEY (userID) REFERENCES users (id)
);

-- ウォッチ
CREATE TABLE IF NOT EXISTS watches
(
    userID    varchar     NOT NULL,
    targetID  varchar     NOT NULL,
    createdAt timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (userID) REFERENCES users (id),
    FOREIGN KEY (targetID) REFERENCES users (id)
);

