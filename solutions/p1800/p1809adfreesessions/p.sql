select session_id
from Playback
where session_id not in
(
  select distinct session_id
  from ads a
  join playback p
  on a.customer_id = p.customer_id
    and a.timestamp >= p.start_time
    and a.timestamp <= p.end_time
)