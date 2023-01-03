with a as (
     select *,
        row_number() over (order by id) as rn
     from Seat
), b as (
    select
        a1.id as id1,
        a1.student as student1,
        a2.id as id2,
        a2.student as student2
    from a a1
    left join a a2 on a1.rn = a2.rn-1
    where a1.rn % 2 = 1
), c as (
    select id1 as id, student2 as student
    from b where b.id2 is not null
    union all
    select id2 as id, student1 as student
    from b where b.id2 is not null
)
select *
from c
union all
select id1 as id, student1 as student
from b where b.id2 is null
order by id asc
