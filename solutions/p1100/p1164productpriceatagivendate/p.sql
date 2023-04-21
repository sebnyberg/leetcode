/* Write your T-SQL query statement below */
with a as (
    select product_id, new_price as price,
        rank() over (partition by product_id order by change_date desc) as r
    from Products
    where change_date <= '2019-08-16'
)
select b.product_id,
    case when
        a.price is not null then a.price
        else b.price
    end as price
from (
    select distinct product_id, 10 as price
    from Products
) b
left join a on a.product_id = b.product_id
where a.r = 1 or a.r is null
