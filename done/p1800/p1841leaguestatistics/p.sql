with points as (
    select
           home_team_id,
           away_team_id,
           home_team_goals,
           away_team_goals,
           home_team_points =
    case
        when home_team_goals < away_team_goals then 0
        when home_team_goals = away_team_goals then 1
        when home_team_goals > away_team_goals then 3
    end,
        away_team_points =
     case
         when away_team_goals < home_team_goals then 0
         when away_team_goals = home_team_goals then 1
         when away_team_goals > home_team_goals then 3
     end
    from Matches
), t1 as (
    select home_team_id as team_id,
           sum(home_team_goals) as goal_for,
           sum(away_team_goals) as goal_against,
           count(home_team_id)  as matches_played,
            sum(home_team_points) as points
    from points
    group by home_team_id
    union all
    select away_team_id as team_id,
           sum(away_team_goals) as goal_for,
           sum(home_team_goals) as goal_against,
           count(away_team_id)  as matches_played,
            sum(away_team_points) as points
    from points
    group by away_team_id
), t2 as (
    select team_id,
           sum(goal_for) as goal_for,
           sum(goal_against) as goal_against,
           sum(matches_played) as matches_played,
           sum(goal_for)-sum(goal_against) as goal_diff,
            sum(points) as points
    from t1
    group by team_id
)
select
    team_name,
    matches_played,
       points,
    goal_for,
       goal_against,
       goal_diff
from t2 join Teams t on t.team_id = t2.team_id
order by points desc, goal_diff desc, team_name asc
