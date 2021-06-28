-- выборка фильмов с названием студии
SELECT m.name as movie, c.name as company
FROM movies m
         JOIN company c on m.company_id = c.id;

-- выборка фильмов для некоторого актёра
SELECT m.name
FROM movies m
         JOIN actor_movies am ON m.id = am.movie_id
         JOIN persons p ON am.person_id = p.id
WHERE p.first_name = 'Liv'
  and p.last_name = 'Tyler';

-- подсчёт фильмов для некоторого режиссёра
SELECT count(m.id)
FROM movies m
         JOIN director_movies dm ON m.id = dm.movie_id
         JOIN persons p ON dm.person_id = p.id
WHERE p.first_name = 'Peter'
  and p.last_name = 'Jackson';

-- подсчёт количества фильмов со сборами больше 1000
SELECT count(m.id)
FROM movies m
WHERE m.gross > 1000;

-- выборка различных фамилий актёров
SELECT DISTINCT p.last_name
FROM actor_movies am
         JOIN persons p ON am.person_id = p.id
ORDER BY p.last_name;
