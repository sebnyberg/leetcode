with a as (
   select *,
        lag(id, 2, -10) over (order by id) as prev2,
        lag(id, 1, -10) over (order by id) as prev1,
        lead(id, 1, -10) over (order by id) as next1,
        lead(id, 2, -10) over (order by id) as next2
    from stadium
    where people >= 100
)
select a.id, a.visit_date, a.people
from a
where (prev2 = id-2 and prev1 = id-1)
    or (prev1 = id-1 and next1 = id+1)
    or (next1 = id+1 and next2 = id+2)
order by visit_date
