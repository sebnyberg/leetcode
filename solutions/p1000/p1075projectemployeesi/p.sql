/* Write your T-SQL query statement below */
select p.project_id,
    round(
        sum(cast(e.experience_years as decimal(10,2)))
        /count(*),
        2
    ) as average_years
from Employee e
join Project p
on e.employee_id = p.employee_id
group by p.project_id
order by p.project_id
