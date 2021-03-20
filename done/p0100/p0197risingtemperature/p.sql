-- try 1
WITH pre AS (
    SELECT *,
        LAG(Temperature) OVER (ORDER BY RecordDate) AS PreviousTemp,
        LAG(RecordDate) OVER (ORDER BY RecordDate) AS PreviousDate
    FROM Weather
)
SELECT Id
FROM pre
WHERE Temperature > PreviousTemp
AND DATEDIFF(day, PreviousDate, RecordDate) = 1

-- try 2
WITH rownum AS (
    SELECT *, ROW_NUMBER() OVER (ORDER BY RecordDate) rn
    FROM Weather
), pre AS (
    SELECT *,
        LAG(Temperature) OVER (ORDER BY rn) AS PreviousTemp,
        LAG(RecordDate) OVER (ORDER BY rn) AS PreviousDate
    FROM rownum
)
SELECT Id
FROM pre
WHERE Temperature > PreviousTemp
AND DATEDIFF(day, PreviousDate, RecordDate) = 1
