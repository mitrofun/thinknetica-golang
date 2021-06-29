-- company
INSERT INTO company (id, name)
VALUES  (1, 'New Line Cinema'),
        (2, 'Metro-Goldwyn-Mayer (MGM)'),
        (3, 'Universal Pictures');

-- movies
INSERT INTO movies (id, name, release_year, company_id, gross, rating)
VALUES  (2, 'The Hobbit: An Unexpected Journey', 2012, 2, 1017003568, 'PG-13'),
        (4, 'The Lord of the Rings: The Return of the King', 2003, 1, 1140682011, 'PG-13'),
        (3, 'The Lord of the Rings: The Two Towers', 2002, 1, 936689735, 'PG-13'),
        (1, 'The Lord of the Rings: The Fellowship of the Ring', 2001, 1, 880839846, 'PG-13'),
        (6, 'The Hobbit: The Battle of the Five Armies', 2014, 2, 956019788, 'PG-13'),
        (5, 'The Hobbit: The Desolation of Smaug', 2013, 2, 958366855, 'PG-13'),
        (7, 'Back to the Future Part II', 1989, 3, 331950002, 'PG-10');
