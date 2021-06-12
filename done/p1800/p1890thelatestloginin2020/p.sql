with a as (
    select *,
        row_number() over(partition by user_id order by time_stamp desc) as rank
    from logins
    where time_stamp between '2020-01-01' and '2021-01-01'
)
select user_id, time_stamp as last_stamp
from a
where rank = 1