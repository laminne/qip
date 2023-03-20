CREATE TABLE IF NOT EXISTS "follows"
(
    createdAt timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,

    userID    varchar     NOT NULL,
    FOREIGN KEY (userID) REFERENCES users (id),
    targetID  varchar     NOT NULL,
    FOREIGN KEY (targetID) REFERENCES users (id)
);