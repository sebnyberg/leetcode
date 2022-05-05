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
)
select employee_id
from junior_ids
union all
select employee_id
from senior_ids
