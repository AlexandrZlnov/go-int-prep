/*
Задача:
сумма заказов по пользователям

Таблицы:
users
-----
id | name
1  | Alice
2  | Bob

orders
------
id | user_id | amount
1  | 1       | 100
2  | 1       | 250
3  | 2       | 400

Условие:
Посчитать общую сумму заказов для каждого пользователя

*/

// Решение

// Вариант 1
SELECT u.id, 
	u.name, 
	SUM(o.amount) AS total_amount 
FROM orders o 
JOIN users u ON o.user_id = u.id
GROUP BY u.id, u.name
ORDER BY total_amounе DESC;

