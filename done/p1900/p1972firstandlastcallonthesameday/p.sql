with both_ways as (
    select *
    from calls
    union all
    select recipient_id as caller_id,
        caller_id as recipient_id,
        call_time
    from calls
), distinct_both as (
    select distinct caller_id, recipient_id, call_time
    from both_ways
), first_calls as (
    select top 1 with ties caller_id, recipient_id, day(call_time) call_day
    from distinct_both
    order by row_number() over (partition by caller_id, day(call_time) order by call_time asc)
), last_calls as (
    select top 1 with ties caller_id, recipient_id, day(call_time) call_day
    from distinct_both
    order by row_number() over (partition by caller_id, day(call_time) order by call_time desc)
)
select distinct fc.caller_id as user_id
from first_calls fc
join last_calls lc
on fc.recipient_id = lc.recipient_id
    and fc.caller_id = lc.caller_id
    and fc.call_day = lc.call_day
