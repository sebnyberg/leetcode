with upd as (
    select tp.team_id,
           tp.points + coalesce(sum(pc.points_change), 0) as points
    from TeamPoints tp
             left join PointsChange pc on tp.team_id = pc.team_id
    group by tp.team_id, tp.points, tp.name
), before as (
    select
       team_id,
       name,
       rank() over (order by points desc, name asc) as rank
    from TeamPoints
), after as (
    select
        upd.team_id,
        tp.name,
        rank() over (order by upd.points desc, tp.name asc) as rank
    from upd join TeamPoints tp on upd.team_id = tp.team_id
)
select after.team_id,
       after.name,
       before.rank - after.rank as rank_diff
from after join before on after.team_id = before.team_id
