DROP TYPE IF EXISTS rating CASCADE;
DROP TABLE IF EXISTS actors_movies, directors_movies, company, persons, movies;

CREATE TYPE rating AS ENUM ('PG-10', 'PG-13', 'PG-18');

CREATE TABLE company
(
    id   SERIAL PRIMARY KEY,   -- первичный ключ
    name VARCHAR(128) NOT NULL -- название студии
);

CREATE TABLE persons
(
    id         SERIAL PRIMARY KEY,   -- первичный ключ
    first_name VARCHAR(32) NOT NULL, -- имя
    last_name  VARCHAR(32) NOT NULL, -- фамилия
    birthday   DATE        NOT NULL  -- дата рождения
);

CREATE TABLE movies
(
    id           SERIAL PRIMARY KEY,                                                  -- первичный ключ
    name         VARCHAR(128) NOT NULL,                                               -- название фильма
    release_year INTEGER      NOT NULL DEFAULT 1800 CHECK (release_year >= 1800),     -- год выхода
    company_id   INTEGER REFERENCES company (id) ON DELETE CASCADE ON UPDATE CASCADE, -- студия
    gross        BIGINT       NOT NULL DEFAULT 0,                                     -- сборы
    rating       rating,                                                              -- рейтинг
    UNIQUE (name, release_year)
);

CREATE TABLE actors_movies
(
    id        BIGSERIAL PRIMARY KEY,
    person_id INTEGER REFERENCES persons (id) ON DELETE CASCADE ON UPDATE CASCADE,
    movie_id  INTEGER REFERENCES movies (id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (person_id, movie_id)
);

CREATE TABLE directors_movies
(
    id        BIGSERIAL PRIMARY KEY,
    person_id INTEGER REFERENCES persons (id) ON DELETE CASCADE ON UPDATE CASCADE,
    movie_id  INTEGER REFERENCES movies (id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (person_id, movie_id)
);

CREATE INDEX IF NOT EXISTS company_name_idx ON company (lower(name));
CREATE INDEX IF NOT EXISTS movies_name_idx ON movies (lower(name));
