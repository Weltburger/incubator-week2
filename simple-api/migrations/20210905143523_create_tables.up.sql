CREATE TABLE IF NOT EXISTS users (
  "id"                int4 NOT NULL PRIMARY KEY,
  "email"             VARCHAR (255) NOT NULL UNIQUE,
  "password"          VARCHAR (255) NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    "user_id"           int4 NOT NULL,
    "id"                int4 NOT NULL PRIMARY KEY,
    "title"             VARCHAR (255) NOT NULL UNIQUE,
    "body"              VARCHAR (255) NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
    "post_id"           int4 NOT NULL,
    "id"                int4 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (255) NOT NULL,
    "email"             VARCHAR (255) NOT NULL,
    "body"              VARCHAR (255) NOT NULL
);