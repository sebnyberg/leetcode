with a as (
    select salary, count(salary) as groupsize
    from Employees
    group by salary
)
select employee_id,
       name,
       e.salary,
       dense_rank() over (order by e.salary asc) as team_id
from Employees e join a on a.salary = e.salary
where a.groupsize > 1
order by team_id, employee_id
