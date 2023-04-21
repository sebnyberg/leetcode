with a as (
    select distinct customer_id, product_key
    from Customer
), b as (
    select customer_id, count(product_key) as got
    from a
    group by customer_id
)
select customer_id
from b
where b.got >=
(
    select count(*) as want from Product
)
