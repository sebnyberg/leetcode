WITH UnbannedUsers AS (
    SELECT Users_Id
    FROM Users
    WHERE Banned <> 'No'
)
SELECT request_at AS day,
       CAST(SUM(IIF(Status <> 'completed', 1.0, 0.0))/COUNT(*) AS DECIMAL(4,2)) AS 'cancellation rate'
FROM Trips
WHERE Client_Id NOT IN (SELECT * FROM UnbannedUsers)
    AND Driver_Id NOT IN (SELECT * FROM UnbannedUsers)
    AND Request_at IN ('2013-10-01', '2013-10-02', '2013-10-03')
GROUP BY Request_at