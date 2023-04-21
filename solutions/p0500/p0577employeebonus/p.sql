select e.name, b.bonus
from Employee e
left join Bonus b
on e.empId = b.empId
where b.bonus is NULL or bonus < 1000
