with visits_and_purchases as (
    select M.member_id,
           M.name,
           sum(IIF(V.visit_id is null, 0, 1)) as n_visits,
           sum(IIF(P.visit_id is null, 0, 1)) as n_purchases
    from Members M
             left join Visits V on M.member_id = V.member_id
             left join Purchases P on V.visit_id = P.visit_id
    group by M.member_id, M.name
), conversion as (
    select member_id,
           name,
           coalesce((100 * n_purchases) / nullif(n_visits, 0), 0) as conversion,
           n_visits
    from visits_and_purchases
)
select member_id,
       name,
       case
        when n_visits = 0 then 'Bronze'
       when conversion < 50 then 'Silver'
       when conversion < 80 then 'Gold'
       else 'Diamond'
        end as category
from conversion


