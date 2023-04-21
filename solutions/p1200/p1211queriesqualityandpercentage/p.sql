select
  query_name,
  round(sum(quality)/count(*), 2) as quality,
  round(cast(sum(ispoor)*100 as decimal(10,2))/count(*), 2) as poor_query_percentage
from (
  select
    cast(rating as decimal(10,2)) / position as quality,
    case when rating < 3 then 1 else 0 end as ispoor,
    *
  from Queries
) a
group by query_name
