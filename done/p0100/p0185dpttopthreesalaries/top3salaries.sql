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
