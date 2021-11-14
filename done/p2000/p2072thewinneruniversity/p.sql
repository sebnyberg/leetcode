with c as (
    select coalesce(count(*), 0) as num, state = 'California'
    from California
    where score >= 90
), n as (
    select coalesce(count(*), 0) as num, state = 'NewYork'
    from NewYork
    where score >= 90
)
select
    case
        when c.num > n.num then 'California University'
        when n.num = c.num then 'No Winner'
        else 'New York University'
    end as winner
from c cross join n
