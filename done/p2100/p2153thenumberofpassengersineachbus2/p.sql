with a as (
  -- calculate cumulative sum of passengers and create link between each current
  -- and the next bus based on arrival time
  select
    ntot = count(passenger_id),
    B.arrival_time,
    B.bus_id,
    B.capacity,
    next_bus_id = lead(B.bus_id, 1, 0) over (order by B.arrival_time),
    bus_order = rank() over (order by b.arrival_time)
  from Passengers
    right join Buses B on Passengers.arrival_time <= B.arrival_time
  group by B.arrival_time, B.bus_id, b.capacity
), b as (
  -- calculate number of new passengers as ntot[i] - ntot[i-1]
  select a.*,
    new = ntot - lag(ntot, 1, 0) over(order by arrival_time)
  from a
), prev as (
  -- prev = (the Anchor statement below)
  select
    bus_id,
    next_bus_id,
    new,
    passengers_cnt = iif(capacity >= new, new, capacity),
    remains = iif(capacity >= new, 0, new-capacity)
  from b
  where bus_order = 1
    union all
  -- Recursive statement (refers to previous result in 'prev')
  select
    b.bus_id,
    b.next_bus_id,
    b.new,
    passengers_cnt = iif(b.capacity >= b.new+prev.remains, b.new+prev.remains, b.capacity),
    remains = iif(b.capacity >= b.new+prev.remains, 0, b.new+prev.remains-b.capacity)
  from b
  inner join prev on b.bus_id = prev.next_bus_id
)
select bus_id, passengers_cnt
from prev
order by bus_id asc;
