WITH todelete AS (
    SELECT p1.Id
    FROM Person p1, Person p2
    WHERE p1.Email = p2.Email AND p1.Id > p2.Id
)
DELETE FROM Person
WHERE Id IN (SELECT Id FROM todelete)
