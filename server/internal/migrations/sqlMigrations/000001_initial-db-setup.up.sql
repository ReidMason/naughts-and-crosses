CREATE TABLE users (
  id   SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  token TEXT NOT NULL,
  wins bigint default 0 NOT NULL,
  losses bigint default 0 NOT NULL
);
