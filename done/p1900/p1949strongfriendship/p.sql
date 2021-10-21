with f as (
    select *
    from Friendship
    union
    select user2_id as user1_id, user1_id as user2_id
    from Friendship
)
select
       a.user1_id as user1_id,
       b.user2_id as user2_id,
       count(b.user1_id) as common_friend
from f a
join f b on a.user2_id = b.user1_id
                and a.user1_id < b.user2_id -- friends in common
join Friendship c on c.user1_id = a.user1_id
                         and  c.user2_id = b.user2_id
where a.user1_id < b.user2_id
group by a.user1_id, b.user2_id
having count(a.user2_id) >= 3;
