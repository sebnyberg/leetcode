with a as (
    select distinct user_id, activity_date
    from Activity
    where activity_date
    BETWEEN dateadd(day, -29, '2019-07-27')
        AND '2019-07-28'
)
/* Write your T-SQL query statement below */
select
    format(
        dateadd(DAY, 0, datediff(day, 0, activity_date)),
        'yyyy-MM-dd'
     ) as day,
    count(*) as active_users
from a
GROUP BY dateadd(DAY, 0, datediff(day, 0, activity_date))
