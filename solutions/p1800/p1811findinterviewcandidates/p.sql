with Medals as (
    select gold_medal as user_id, 'gold' as medal, contest_id from Contests
union all
    select silver_medal as user_id, 'silver' as medal, contest_id from Contests
union all
    select bronze_medal as user_id, 'bronze' as medal, contest_id from Contests
), Lagged as (
    select
        user_id, medal, contest_id,
        lag(contest_id, 2) over (partition by user_id order by contest_id) as lag_2,
        count(*) over (partition by user_id, medal) as medal_count
    from Medals
)
select distinct name, mail
from Lagged l
join Users u on u.user_id = l.user_id
where (medal = 'gold' and medal_count >= 3)
    or contest_id = lag_2 + 2


