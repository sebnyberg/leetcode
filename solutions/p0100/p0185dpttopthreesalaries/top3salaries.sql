/* Write your T-SQL query statement below */
-- Below is a better solution:
-- select a.Department, a.Employee , a.Salary 
-- from (
--     select d.Name as Department, 
--     e.Name as Employee, 
--     e.Salary, 
--     DENSE_RANK() OVER(Partition by d.Id order by Salary desc) as rnk
--     from Employee as e join Department as d 
--     on e.DepartmentId = d.Id
-- ) as a
-- where a.rnk <=3

-- Here was my attempt, roughly 20% slower than the above optimal
with GroupedSalaries as (
    select DepartmentId, Salary
    from Employee e
    group by e.DepartmentId, e.Salary
), RankedSalaries as (
    select *, ROW_NUMBER()
    over (
        PARTITION BY gs.DepartmentId
        ORDER BY gs.Salary DESC
    ) as RowNo
    FROM GroupedSalaries gs
), TopSalaries as (
    select *
    from RankedSalaries
    where RowNo <= 3
)
select
       d.Name as Department,
       e.Name as Employee,
       ts.Salary as Salary
from Employee e
join TopSalaries ts
on e.DepartmentId = ts.DepartmentId
    and e.Salary = ts.Salary
join Department d on d.Id = ts.DepartmentId
order by ts.DepartmentId asc, ts.Salary desc
