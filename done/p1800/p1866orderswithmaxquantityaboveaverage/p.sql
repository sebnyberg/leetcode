with avgs as (
    select avg(quantity) as avgq, order_id
    from ordersdetails
    group by order_id
)
/* Write your T-SQL query statement below */
select distinct order_id
from ordersdetails
where quantity > (
    select max(avgq) maxavg
    from avgs
)