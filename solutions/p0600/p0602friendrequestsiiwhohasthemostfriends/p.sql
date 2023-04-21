select top 1 id, count(id) as num
from (
    select requester_id as id from RequestAccepted
    union all
    select accepter_id as id from RequestAccepted
) a
group by id
order by count(id)
