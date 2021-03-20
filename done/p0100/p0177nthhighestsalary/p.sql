CREATE FUNCTION getNthHighestSalary(@N INT) RETURNS INT AS
BEGIN
    RETURN (
      SELECT Salary FROM (
        SELECT
          row_number() OVER (ORDER BY Salary DESC) AS row_num,
          emp.Salary
        FROM (SELECT DISTINCT Salary FROM Employee) emp
      ) result
      WHERE row_num = @N
    );
END