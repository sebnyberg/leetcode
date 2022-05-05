SELECT Email
FROM (
         SELECT COUNT(Id) as N, Email
         FROM Person
         GROUP BY Email
     ) res
WHERE N >= 2