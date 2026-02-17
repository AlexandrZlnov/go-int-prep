/*
Задача: 
пользователи с суммой заказов > 10 000

Таблицы:
orders
------
id | user_id | amount
1  | 1       | 6000
2  | 1       | 5000
3  | 2       | 3000

Условие:
Найти пользователей, у которых сумма заказов больше 10 000

*/

// Выриат 1 - если достаточно просто id
SELECT o.user_id,
	SUM(amount) AS total_amount
FROM orders o 
GROUP BY o.user_id
HAVING SUM(amount) > 10000

// Выриат 2 - есл нужно имя пользователя
SELECT u.id,
	u.name,
	SUM(o.amount) AS total_amount
FROM orders o 
JOIN users u ON o.user_id = u.id
GROUP BY u.id, u.name
HAVING SUM(o.amount) > 10000