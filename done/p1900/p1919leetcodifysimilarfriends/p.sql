select distinct f.user1_id, f.user2_id
from Friendship f,
     Listens l1,
     Listens l2
where l1.day = l2.day
    and l1.song_id = l2.song_id
    and f.user1_id = l1.user_id
    and f.user2_id = l2.user_id
group by l1.day, f.user1_id, f.user2_id
having count(distinct l1.song_id) >= 3