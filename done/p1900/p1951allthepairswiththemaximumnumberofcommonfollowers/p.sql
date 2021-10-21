select 
     user1_id,
     user2_id
from (
     select 
          a.user_id as user1_id,
          b.user_id as user2_id,
          rank() over (order by count(a.user_id) desc) as global_rank
     from Relations a
     join Relations b 
          on a.follower_id = b.follower_id 
               and a.user_id < b.user_id
     group by a.user_id, b.user_id
) q
where global_rank = 1
