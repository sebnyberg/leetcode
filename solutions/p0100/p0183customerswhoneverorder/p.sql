SELECT Name as Customers
FROM Customers c
WHERE c.ID NOT IN (SELECT DISTINCT CustomerId FROM Orders)