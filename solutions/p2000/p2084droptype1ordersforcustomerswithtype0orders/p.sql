select *
from orders
where order_type = 0
union all
select *
from orders
where order_type = 1
and customer_id not in (
  select distinct customer_id
  from orders
  where order_type = 0
) a