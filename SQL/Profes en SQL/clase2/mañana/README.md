### PRIMERA PARTE

1. ¿A qué se denomina JOIN en una base de datos?

```
La sentencia INNER JOIN es la sentencia JOIN por defecto, y consiste en combinar datos de una tabla con datos de la otra tabla a partir de una o varias condiciones en común.
```

2. Nombre y explique 2 tipos de JOIN.

```
INNER JOIN: recupera la instersección de dos o más tablas de acuerdo a identificadores coincidentes.

LEFT JOIN: damos prioridad a la tabla de la izquierda, y buscamos en la tabla derecha. Si no existe ninguna coincidencia para alguna de las filas de la tabla de la izquierda, de igual forma todos los resultados de la primera tabla se muestran
```

3. ¿Para qué se utiliza el GROUP BY?

```
La sentencia GROUP BY identifica una columna seleccionada para utilizarla para agrupar resultados. Divide los datos en grupos por los valores de la columna especificada, y devuelve una fila de resultados para cada grupo. Puede utilizar GROUP BY con más de un nombre de columna (separe los nombres de columna con comas).
```

4. ¿Para qué se utiliza el HAVING?

```
La cláusula HAVING aplica una condición al conjunto de resultados agrupado intermedio que una consulta devuelve. La condición HAVING se utiliza con funciones incorporadas. Determina si se incluye un grupo completo. HAVING va siempre seguida de una función de columna (como SUM, AVG, MAX, MIN o COUNT). HAVING también puede ir seguida de una subconsulta que busque un valor agrupado para completar la condición HAVING.
```

5. Dado los siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:

```
1. INNER JOIN

2. LEFT JOIN
```

6. Escriba una consulta genérica por cada uno de los diagramas a continuación:

1. RIGTH JOIN

2. FULL OUTER JOIN


Imagen de aplicación de diferentes tipos de join:

![sqljoin](sqljoin.jpeg)
Link de información e ilustraciones complementarios: https://programacionymas.com/blog/como-funciona-inner-left-right-full-join

https://ingenieriadesoftware.es/tipos-sql-join-guia-referencia/

### SEGUNDA PARTE

1. Mostrar el título y el nombre del género de todas las series.

```sql
SELECT m.title, g.name FROM movies AS m INNER JOIN genres AS g 
	ON m.genre_id = g.id;
```

2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.

```sql
SELECT e.title AS "Episodio", concat(a.first_name, " ", a.last_name) AS "Actor" FROM episodes AS e 
	INNER JOIN actor_episode AS a_ep ON a_ep.episode_id = e.id
    INNER JOIN actors AS a ON a_ep.actor_id = a.id
```

3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.

```sql 
SELECT s.title AS "Serie", COUNT(t.id) AS "Total temporadas" 
	FROM series AS s 
	RIGHT JOIN seasons t ON t.serie_id = s.id     
	GROUP BY s.title;
```

4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.

```sql
SELECT g.name, COUNT(m.id) FROM movies AS m 
	LEFT JOIN genres AS g ON m.genre_id = g.id 
    GROUP BY g.name
    HAVING COUNT(m.id) >= 3
```
5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.

```sql
SELECT a.first_name, a.last_name FROM actors AS a 
	INNER JOIN actor_movie AS am ON a.id = am.actor_id
    INNER JOIN movies AS m ON m.id = am.movie_id
    WHERE m.title LIKE "La Guerra de las Galaxias%"
    GROUP BY a.first_name, a.last_name;

```
