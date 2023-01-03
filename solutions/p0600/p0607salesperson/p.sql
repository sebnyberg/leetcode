SELECT DISTINCT [name]
FROM SalesPerson sp
WHERE sp.[name] NOT IN
(
    SELECT s.[name]
    FROM Company c
    INNER JOIN Orders o ON o.com_id = c.com_id
    INNER JOIN SalesPerson s ON s.sales_id = o.sales_id
    WHERE c.[name] = 'RED'
)
