select count(*) as rich_count
from (
    select distinct customer_id
    from store
    where amount > 500
) t