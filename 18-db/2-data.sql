-- company
INSERT INTO company (id, name)
VALUES  (1, 'New Line Cinema'),
        (2, 'Metro-Goldwyn-Mayer (MGM)'),
        (3, 'Universal Pictures');

-- persons
INSERT INTO persons (id, first_name, last_name, birthday)
VALUES  (1, 'Elijah', 'Wood', '1981-01-26'),
        (2, 'Ian', 'McKellen', '1939-05-25'),
        (3, 'Sean', 'Astin', '1971-02-25'),
        (4, 'Viggo', 'Mortensen', '1958-10-20'),
        (5, 'Billy', 'Boyd', '1968-08-28'),
        (6, 'Dominic', 'Monaghan', '1976-12-08'),
        (7, 'John', 'Rhys-Davies', '1944-05-05'),
        (8, 'Orlando', 'Bloom', '1977-01-13'),
        (9, 'Sean', 'Bean', '1959-04-17'),
        (10, 'Liv', 'Tyler', '1977-07-01'),
        (11, 'Cate', 'Blanchett', '1969-05-14'),
        (12, 'Hugo', 'Weaving', '1960-04-04'),
        (13, 'Peter', 'Jackson', '1961-10-31'),
        (14, 'Martin', 'Freeman', '1971-09-08'),
        (15, 'Ian', 'Holm', '1931-09-12'),
        (16, 'Robert', 'Zemeckis', '1951-05-14');

-- movies
INSERT INTO movies (id, name, release_year, company_id, gross, rating)
VALUES  (2, 'The Hobbit: An Unexpected Journey', 2012, 2, 1017003568, 'PG-13'),
        (4, 'The Lord of the Rings: The Return of the King', 2003, 1, 1140682011, 'PG-13'),
        (3, 'The Lord of the Rings: The Two Towers', 2002, 1, 936689735, 'PG-13'),
        (1, 'The Lord of the Rings: The Fellowship of the Ring', 2001, 1, 880839846, 'PG-13'),
        (6, 'The Hobbit: The Battle of the Five Armies', 2014, 2, 956019788, 'PG-13'),
        (5, 'The Hobbit: The Desolation of Smaug', 2013, 2, 958366855, 'PG-13'),
        (7, 'Back to the Future Part II', 1989, 3, 331950002, 'PG-10');

-- directors
INSERT INTO director_movies (person_id, movie_id)
VALUES  (13, 1),
        (13, 2),
        (13, 3),
        (13, 4),
        (13, 5),
        (13, 6),
        (16, 7);

-- actors
INSERT INTO public.actor_movies (person_id, movie_id)
VALUES  (1, 1),
        (2, 1),
        (3, 1),
        (4, 1),
        (5, 1),
        (6, 1),
        (7, 1),
        (8, 1),
        (9, 1),
        (10, 1),
        (11, 1),
        (12, 1),
        (1, 3),
        (2, 3),
        (3, 3),
        (4, 3),
        (5, 3),
        (6, 3),
        (7, 3),
        (8, 3),
        (9, 3),
        (10, 3),
        (11, 3),
        (12, 3),
        (1, 4),
        (2, 4),
        (3, 4),
        (4, 4),
        (5, 4),
        (6, 4),
        (7, 4),
        (8, 4),
        (9, 4),
        (10, 4),
        (11, 4),
        (12, 4),
        (14, 2),
        (14, 5),
        (14, 6),
        (2, 2),
        (2, 5),
        (2, 6),
        (15, 1),
        (15, 2),
        (15, 3),
        (15, 4),
        (15, 5),
        (1, 2),
        (1, 7);
