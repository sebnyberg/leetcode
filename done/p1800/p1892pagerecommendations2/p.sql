with a as (
    select user1_id as first, user2_id as friend
    from Friendship
    union all
    select user2_id as first, user1_id as friend
    from Friendship
)
select a.first as user_id, page_id, count(first) as friends_likes
from a
join Likes l on a.friend = l.user_id
where a.first not in (select user_id from Likes where page_id = l.page_id)
group by a.first, page_id
order by count(first) desc, a.first;