/* Write your T-SQL query statement below */
select top 1 person_name
from (
  select *,
    SUM(weight) over (order by turn) as cum
  from Queue
) a
where a.cum <= 1000
order by a.cum desc
