with winners as (
    select year, Wimbledon as winner, 'Wimbledon' as championship
    from Championships
    union all
    select year, Fr_open as winner, 'Fr_open' as championship
    from Championships
    union all
    select year, US_open as winner, 'US_open' as championship
    from Championships
    union all
    select year, Au_open as winner, 'Au_open' as championship
    from Championships
)
select player_id, player_name, count(player_name) as grand_slams_count
from winners w
join Players p
on w.winner = p.player_id
group by player_id, player_name