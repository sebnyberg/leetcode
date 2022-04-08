with a as (
    select *, row_number() over(partition by company order by salary) as r,
        count(id) over (partition by company) as cnt
    from Employee
)
select id, company, salary
from a
where r = (cnt+2)/2
    or r = (cnt+1)/2