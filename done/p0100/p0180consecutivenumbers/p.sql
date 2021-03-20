SELECT DISTINCT l1.Num as ConsecutiveNums
FROM Logs l1
JOIN Logs l2 on l1.Id+1 = l2.Id
JOIN Logs l3 on l2.Id+1 = l3.Id
WHERE l1.Num = l2.Num
    AND l2.Num = l3.Num
