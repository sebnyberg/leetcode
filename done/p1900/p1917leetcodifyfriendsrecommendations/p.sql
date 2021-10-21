with res as (
    select
        distinct a.user_id as user_id,
        b.user_id as recommended_id
    from Listens a, Listens b
    where a.day = b.day
      and a.song_id = b.song_id
      and a.user_id < b.user_id
      and concat(a.user_id, '->', b.user_id) not in (
          select concat(user1_id, '->', user2_id)
          from Friendship
      )
    group by a.user_id, b.user_id, b.day
    having count(distinct a.song_id) >= 3
)
select * from res
union all
select recommended_id as user_id,
    user_id as recommended_id
from res
order by user_id, recommended_id
