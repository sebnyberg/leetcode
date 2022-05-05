with ranked as (
    select *,
        row_number() over (order by first_col asc) as first_n,
        row_number() over (order by second_col desc) as second_n
    from Data
)
select a.first_col, b.second_col
from ranked a
join ranked b
    on a.first_n = b.second_n