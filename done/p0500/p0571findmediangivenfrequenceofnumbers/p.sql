with a as (
    select *,
           sum(frequency) over (order by num) - frequency + 1 as start_sum,
           sum(frequency) over (order by num) as end_sum
    from Numbers
), b as (
    select *
    from a
    where a.start_sum <= (( select max(end_sum) from a ) + 2) / 2
      and a.end_sum >= (( select max(end_sum) from a ) + 2) / 2
    union all
    select *
    from a
    where a.start_sum <= (( select max(end_sum) from a ) + 1) / 2
      and a.end_sum >= (( select max(end_sum) from a ) + 1) / 2
)
select round(cast(sum(num) as float) / 2, 1) as median
from b