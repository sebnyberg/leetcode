select s.product_id,
    s.year as first_year,
    s.quantity,
    s.price
from Sales s
inner join (
    select product_id, min(year) as year
    from Sales
    group by product_id
) b
on s.product_id = b.product_id and s.year = b.year
