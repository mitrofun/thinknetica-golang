-- выборка фильмов с названием студии
SELECT m.name as movie, c.name as company
FROM movies m
         JOIN company c on m.company_id = c.id;

-- выборка фильмов для некоторого актёра
SELECT m.name
FROM movies m
         JOIN actors_movies ON m.id = actors_movies.movie_id
         JOIN persons p ON actors_movies.person_id = p.id
WHERE p.first_name = 'Liv'
  and p.last_name = 'Tyler';

-- подсчёт фильмов для некоторого режиссёра
SELECT count(m.id)
FROM movies m
         JOIN directors_movies ON m.id = directors_movies.movie_id
         JOIN persons p ON directors_movies.person_id = p.id
WHERE p.first_name = 'Peter'
  AND p.last_name = 'Jackson';

-- подсчёт количества фильмов со сборами больше 1000
SELECT count(m.id)
FROM movies m
WHERE m.gross > 1000;

-- выборка различных фамилий актёров
SELECT DISTINCT p.last_name
FROM actors_movies
         JOIN persons p ON actors_movies.person_id = p.id
ORDER BY p.last_name;
