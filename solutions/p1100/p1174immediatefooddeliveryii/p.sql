/* Write your T-SQL query statement below */
with a as (
    select *,
        case when order_date = customer_pref_delivery_date
            then 1
            else 0
        end as immediate
    from (
        select *,
            r = rank() over (partition by customer_id order by order_date asc)
        from Delivery
    ) a
    where a.r = 1
)
select round(sum(cast(immediate as decimal(10,2)))*100/count(*), 2) as immediate_percentage
from a
