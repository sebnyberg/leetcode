WITH top_sales AS (
    SELECT transaction_id,
           RANK() over (PARTITION BY dateadd(DAY, datediff(DAY, 0, day), 0) ORDER BY amount DESC) as rank
    FROM Transactions
)
SELECT transaction_id
FROM top_sales
WHERE rank = 1
ORDER BY transaction_id ASC
