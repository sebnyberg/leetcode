with a as (
  select
    first_value(event_date) over (partition by player_id order by event_date) as f,
    *
  from Activity
), b as (
  select distinct player_id, 1 as val
  from a
  where event_date = dateadd(dd, 1, f)
), c as (
  select distinct a.player_id, b.val
  from a
    left join b on a.player_id = b.player_id
)
select round(cast(sum(val) as float) / count(*), 2) as fraction
from c