with a as (
    select *,
           row_number() over (partition by player_id order by match_day) as day
    from matches
), b as (
    select m.*,
           day - row_number() over (partition by player_id order by match_day) as win_group
    from a m
    where result = 'Win'
), c as (
    select player_id, count(win_group) as win_streak
    from b
    group by player_id, win_group
)
select p.player_id, coalesce(max(win_streak), 0) as longest_streak
from (select distinct(player_id) from matches m) p
         left join c on p.player_id = c.player_id
group by p.player_id