
with salary as (
    select employee_id,
           experience,
           salary_sum = sum(salary) over (
               partition by experience
               order by salary rows unbounded preceding
           )
    from Candidates
), senior_ids as (
    select *
    from salary
    where experience = 'Senior' and salary_sum <= 70000
), junior_ids as (
    select *
    from salary
    where experience = 'Junior'
    and salary_sum <= 70000 - (
        select iif(max(salary_sum) <= 70000, max(salary_sum), 0)
        from senior_ids
    )
), ids as (
    select *
    from junior_ids
    union all
    select *
    from senior_ids
)
select a.experience,
       accepted_candidates=count(employee_id)
from (select distinct experience from Candidates) a
left join ids b
on a.experience = b.experience
group by a.experience;
