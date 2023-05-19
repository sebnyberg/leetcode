select
  p.product_id,
  cast(cast(sum(p.price*u.units) as decimal(10,2))/cast(sum(u.units) as decimal(10, 2)) as decimal(10,2)) as average_price
from Prices p
join UnitsSold u
on p.product_id = u.product_id
where u.purchase_date between p.start_date and p.end_date
group by p.product_id
