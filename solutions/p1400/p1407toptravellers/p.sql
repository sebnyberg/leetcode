select
    name,
    coalesce(travelled_distance, 0) as travelled_distance
from Users u
left join (
     select user_id, sum(distance) as travelled_distance
     from Rides
     group by user_id
) r
  on u.id = r.user_id
order by travelled_distance desc, name asc
