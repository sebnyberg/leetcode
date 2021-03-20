-- Version 1
SELECT s.Score as score, rank as Rank
FROM (
     SELECT
            row_number() over (ORDER BY Score DESC) as rank,
            s.Score
     FROM (SELECT DISTINCT Score FROM Scores) s
) ranks
JOIN Scores s  ON ranks.Score = s.Score
ORDER BY rank


-- Version 2
SELECT Score, DENSE_RANK() OVER(ORDER BY Score DESC) rank
FROM Scores
ORDER BY rank