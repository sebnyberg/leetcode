SELECT problem_id
FROM Problems
WHERE cast(likes as float)/(cast(likes + dislikes as float)) < 0.6
ORDER BY problem_id ASC