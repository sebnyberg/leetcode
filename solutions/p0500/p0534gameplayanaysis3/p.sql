select
  player_id,
  event_date,
  games_played_so_far = sum(games_played) over (partition by player_id order by event_date)
from Activity
order by player_id, event_date