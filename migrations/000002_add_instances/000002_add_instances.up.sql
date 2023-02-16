CREATE TABLE IF NOT EXISTS instances
(
    id        varchar PRIMARY KEY NOT NULL,
    host      varchar UNIQUE      NOT NULL,
    userCount integer             NOT NULL,
    postCount integer NOT NULL,
    watchingCount integer NOT NULL,
    watchedCount integer NOT NULL,
    softwareName varchar NOT NULL,
    version varchar NOT NULL,
    name varchar NOT NULL,
    description varchar NOT NULL,
    maintainerName varchar NOT NULL,
    maintainerEmail varchar NOT NULL
)