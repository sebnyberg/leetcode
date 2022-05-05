with a as (
    select departure_airport as airport_id, flights_count
    from flights
    union all
    select arrival_airport as airport_id, flights_count
    from flights
), b as (
    select airport_id, sum(flights_count) as flights
    from a
    group by airport_id
)
select airport_id
from b
where flights = (select max(flights) from b)