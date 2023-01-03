SELECT
    distinct t1.id AS [id],
    CASE
        WHEN t2.id IS NULL THEN 'Root'
        WHEN t3.id IS NULL THEN 'Leaf'
        ELSE 'Inner'
    END AS [type]
FROM Tree t1
LEFT JOIN Tree t2 on t1.p_id = t2.id
LEFT JOIN Tree t3 on t1.id = t3.p_id
