/*TEMPORARY TABLES*/
use movies_db;
DROP TABLE TWD;
CREATE TEMPORARY TABLE TWD
SELECT e.*, t.serie_id
FROM episodes e INNER JOIN seasons t ON e.season_id = t.id
WHERE t.serie_id = (SELECT id
			   FROM series
			   WHERE title = "The Walking Dead");

SELECT * FROM TWD;

/*INDEXES*/
ALTER TABLE series ADD INDEX idx_series_title (title);
SHOW INDEXES FROM series;

/*Agregar una película a la tabla movies.*/
INSERT INTO movies values (33,"2022-10-01","2022-10-02","No respires 2",4,0,"2022-11-04",1,3);

/*Agregar un género a la tabla genres.*/
INSERT INTO genres values(13,"2022-10-01","2022-10-02","Independiente",13,1);

/*Asociar a la película del punto 1. con el género creado en el punto 2.*/
UPDATE movies set genre_id = 1 WHERE id = 33;

/*Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.*/
UPDATE actors set favorite_movie_id = 33 WHERE id = 3;

/*Crear una tabla temporal copia de la tabla movies.*/
DROP TABLE IF EXISTS temp_movies;
CREATE TEMPORARY TABLE temp_movies
select * from movies;

/*Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.*/
DELETE FROM temp_movies WHERE awards < 5;
select * from temp_movies;

/*Obtener la lista de todos los géneros que tengan al menos una película.*/
SELECT g.name, g.id
FROM genres g
WHERE EXISTS (SELECT m.genre_id FROM movies m WHERE m.genre_id = g.id);

/*Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.*/
SELECT a.first_name, a.last_name, a.favorite_movie_id, m.awards
FROM actors a INNER JOIN movies m ON a.favorite_movie_id = m.id
WHERE m.awards > 3;

/*Crear un índice sobre el nombre en la tabla movies.*/
ALTER TABLE movies ADD INDEX idx_movies_title (title);

/*Chequee que el índice fue creado correctamente.*/
SHOW INDEXES FROM movies;
