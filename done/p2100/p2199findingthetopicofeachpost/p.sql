with a as (
  select distinct post_id, topic_id
  from Posts p
  left join Keywords k
    on p.content = k.word
    or p.content like '% ' + k.word + ' %'
    or p.content like '% ' + k.word
    or p.content like k.word + ' %'
)
select
  post_id,
  coalesce(string_agg(a.topic_id, ',') within group(order by topic_id), 'Ambiguous!') as topic
from a
group by post_id
