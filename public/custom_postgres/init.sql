CREATE DATABASE users;

\c users;

CREATE TABLE users (
  -- id SERIAL PRIMARY KEY,
  user_id VARCHAR(255),
  name VARCHAR(255),
  number INTEGER,
  password VARCHAR(255),
  reshte VARCHAR(255),
  vorudi VARCHAR(255),
  courses TEXT[],
  groups TEXT[],
  role INTEGER
);

CREATE TABLE groups (
  -- id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  subgroups TEXT[],
  courses TEXT[],
  members TEXT[],
  owner VARCHAR(255),
  news TEXT[]
);

CREATE TABLE subgroups (
  -- id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  members TEXT[],
  courses TEXT[],
  owner VARCHAR(255)
);

CREATE TABLE courses (
  -- id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  exercises TEXT[],
  members TEXT[],
  owner VARCHAR(255)
);


-- INSERT INTO users (user_id, name, password, reshte, vorudi, courses, groups, role)
-- VALUES
--   ('John1', 'Doe', "password", 3123123, 'ce', '1400', ARRAY['math', 'algebra'], ARRAY['groupA', 'groupB'],  1);

