with a as (
  select top 1 [name] as results
  from MovieRating mr
  join Users u
    on u.user_id = mr.user_id
  group by u.user_id, [name]
  order by count([name]) desc, [name]
), b as (
  select top 1 title as results
  from Movies m
  join MovieRating mr
      on m.movie_id = mr.movie_id
  where mr.created_at >= '2020-02-01' and mr.created_at < '2020-03-01'
  group by title
  order by sum(cast(rating as float))/cast(count([title]) as float) desc, [title]
)
select *
from a
union all
select *
from b
