with cte as (
    select *,
           lag(time_stamp) over (partition by user_id order by time_stamp)
               as previous_confirm
    from confirmations
)
select distinct user_id
from cte
where previous_confirm is not null
    and datediff(second, previous_confirm, time_stamp) <= 60*60*24
