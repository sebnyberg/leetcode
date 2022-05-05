/* Write your T-SQL query statement below */
with unioned as (
    select product_id, 'store1' as store, store1 as price
    from products
union all
    select product_id, 'store2' as store, store2 as price
    from products
union all
    select product_id, 'store3' as store, store3 as price
    from products
)
select *
from unioned u
where price is not null