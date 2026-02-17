/*
Задача:
дерево + рекурсия

Таблица:
nodes
-----
id | parent_id | value
1  | NULL      | 10
2  | 1         | 5
3  | 1         | 7
4  | 2         | 3

Условие:
Для узла id = 1 посчитать сумму всех дочерних узлов (включая его самого)

*/

WITH RECURSIVE summ_value AS (
    SELECT id, parent_id, value
    FROM nodes
    WHERE parent_id IS NULL

    UNION ALL

    SELECT n.id, n.parent_id, n.value
    FROM nodes n
    JOIN summ_value sv ON n.parent_id = sv.id
)
SELECT SUM(value) AS total_summ
FROM summ_value;

