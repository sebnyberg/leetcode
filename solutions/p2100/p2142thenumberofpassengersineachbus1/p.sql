with a as (
    select passenger_id, min(B.arrival_time) as bus_arrival_time
    from Passengers
           join Buses B on Passengers.arrival_time <= B.arrival_time
    group by passenger_id
)
select bus_id, count(passenger_id) as passengers_cnt
from a
right join Buses b
on a.bus_arrival_time = b.arrival_time
group by b.bus_id
order by b.bus_id
