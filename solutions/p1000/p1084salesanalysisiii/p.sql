/* Write your T-SQL query statement below */
select distinct p.product_id, p.product_name
from Sales s
join Product p on s.product_id = p.product_id
where p.product_id in (
    select product_id
    from Sales
    where sale_date between '2019-01-01' and '2019-04-01'
)
and p.product_id not in (
    select product_id
    from Sales
    where sale_date not between '2019-01-01' and '2019-04-01'
)
