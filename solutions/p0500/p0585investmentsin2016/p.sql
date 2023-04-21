select tiv_2016 = round(sum(tiv_2016), 2)
from Insurance a
where not exists (
  select *
  from Insurance b
  where a.pid <> b.pid and a.lat = b.lat and a.lon = b.lon
)
and exists (
  select *
  from Insurance b
  where a.pid <> b.pid and a.tiv_2015 = b.tiv_2015
)
