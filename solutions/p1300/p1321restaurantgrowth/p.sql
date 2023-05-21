with a as (
  select
    visited_on,
    sum(amount) over (order by visited_on rows between 6 preceding and current row) as amount,
    round(avg(amount) over (order by visited_on rows between 6 preceding and current row), 2) as average_amount
  from (
    select
      visited_on,
      cast(sum(amount) as float) as amount
    from customer
    group by visited_on
  ) b
)
select *
from a
where visited_on >= (
  select dateadd(day, 6, min(visited_on))
  from customer
)
