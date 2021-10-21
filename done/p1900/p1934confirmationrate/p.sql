select
    s.user_id,
    round(
      cast(sum(case when c.action = 'confirmed' then 1 else 0 end) as float)
        / count(s.user_id), 2
      ) as confirmation_rate

from confirmations c
right join signups s
on c.user_id = s.user_id
group by s.user_id
