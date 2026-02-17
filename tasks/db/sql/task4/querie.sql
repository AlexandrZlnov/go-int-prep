/*
Задача: 
самый дорогой заказ у каждого пользователя

Таблицы:
orders
------
id | user_id | amount
1  | 1       | 100
2  | 1       | 300
3  | 2       | 250

Условие:
Для каждого пользователя найти максимальную сумму заказа

*/

//Вариант1 - если нужен тольк user_id, без id заказа
SELECT user_id, MAX(amount) AS amount_MAX
FROM orders
GROUP BY user_id

//Вариант2 - результат с id заказа через оконную функцию
SELECT user_id, id, amount
FROM (
    SELECT *,
           ROW_NUMBER() OVER (
               PARTITION BY user_id
               ORDER BY amount DESC
           ) AS rn
    FROM orders
) t
WHERE rn = 1;

//Вариант3 - с id заказ через CTE
WITH ranked_orders AS (
    SELECT *,
           ROW_NUMBER() OVER (
               PARTITION BY user_id
               ORDER BY amount DESC
           ) AS rn
    FROM orders
)
SELECT user_id, id, amount
FROM ranked_orders
WHERE rn = 1;