/* Write your T-SQL query statement below */
with a as(
    select
        buyer_id,
        join_date,
        count(*) as orders_in_2019
    from Users u
    left join Orders o
        on u.user_id = o.buyer_id
    where o.order_date between '2019-01-01' and '2020-01-01'
    group by buyer_id, join_date
), b as (
    select
        user_id as buyer_id,
        join_date,
        0 as orders_in_2019
    from Users
    where user_id not in (
        select distinct buyer_id
        from Orders
        where order_date between '2019-01-01' and '2020-01-01'
    )
)
select *
from a
union all
select *
from b
