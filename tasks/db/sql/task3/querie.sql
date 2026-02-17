/*
Задача:
пользователи без заказов

Таблицы:
users
-----
id | name
1  | Alice
2  | Bob
3  | Charlie

orders
------
id | user_id | amount
1  | 1       | 100
2  | 2       | 200

Условие:
Найти пользователей, которые не сделали ни одного заказа
(Charlie — правильный ответ)

*/

// Решение

// Вариант 1
SELECT u.id u.name 
FROM users u
LEFT JOIN orders o ON o.user_id = u.id  
WHERE o.id IS NILL
ORDER BY u.name DESC

// Вариант 2 - лучше в таком случае
SELECT u.id, u.name
FROM users u
WHERE NOT EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.user_id = u.id
);