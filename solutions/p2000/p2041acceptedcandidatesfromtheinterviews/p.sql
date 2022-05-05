select candidate_id
from Candidates c
join Rounds r on c.interview_id = r.interview_id
group by candidate_id, years_of_exp
having years_of_exp >= 2 and sum(score) >= 16