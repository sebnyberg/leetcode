WITH Ranking AS (
    SELECT
        row_number() over (
            PARTITION BY DepartmentId
            ORDER BY Salary DESC
        ) rn,
        Salary,
       DepartmentId
    FROM Employee
), TopSalaries AS (
    SELECT Salary, DepartmentId
    FROM Ranking
    WHERE rn = 1
)
SELECT D.Name as Department, e.Name AS Employee, e.Salary as Salary
FROM Employee e
RIGHT JOIN TopSalaries ts ON ts.Salary = e.Salary AND ts.DepartmentId = e.DepartmentId
JOIN Department D on e.DepartmentId = D.Id
ORDER BY d.Name