/*Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.*/
SELECT e.puesto, d.nombre_depto, d.localidad
FROM empleado e INNER JOIN departamento d ON d.depto_nro = e.depto_nro; 

/*Visualizar los departamentos con más de cinco empleados.*/
SELECT d.nombre_depto AS "Departamento", COUNT(e.cod_emp) AS "Cantidad de empleados"
FROM departamento d INNER JOIN empleado e ON e.depto_nro = d.depto_nro 
GROUP BY d.nombre_depto
HAVING COUNT(e.cod_emp) > 5;

/*Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’*/
SELECT e.nombre AS "Empleado", e.salario AS "Salario", d.nombre_depto AS "Departamento"
FROM empleado e INNER JOIN departamento d ON e.depto_nro = d.depto_nro
WHERE e.puesto IN (SELECT puesto 
				   FROM empleado 
                   WHERE nombre = "Mito" AND apellido = "Barchuk")
AND e.nombre != "Mito" AND e.apellido != "Barchuk";

/*Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.*/
SELECT CONCAT(e.nombre, " ", e.apellido) AS "Empleado", d.nombre_depto AS "Departamento"
FROM empleado e INNER JOIN departamento d ON e.depto_nro = d.depto_nro
WHERE d.nombre_depto = "Contabilidad"
ORDER BY e.nombre;

/*Mostrar el nombre del empleado que tiene el salario más bajo.*/
SELECT CONCAT(e.nombre, " ", e.apellido) AS "Empleado", e.salario AS "Salario"
FROM empleado e
ORDER BY e.salario ASC
LIMIT 1;

/*Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.*/
SELECT CONCAT(e.nombre, " ", e.apellido) AS "Empleado", e.salario AS "Salario"
FROM empleado e INNER JOIN departamento d on e.depto_nro = d.depto_nro
WHERE d.nombre_depto = "Ventas"
ORDER BY e.salario DESC
LIMIT 1;
