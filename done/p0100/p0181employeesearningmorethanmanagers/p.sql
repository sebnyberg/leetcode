SELECT NotManager.Name
FROM Employee e
JOIN
(
    SELECT *
    FROM Employee e
    WHERE e.ManagerId IS NOT NULL
) NotManager ON NotManager.ManagerId = e.Id
WHERE NotManager.Salary > e.Salary

