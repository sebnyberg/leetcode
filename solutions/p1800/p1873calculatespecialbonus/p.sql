select
       employee_id,
       IIF(employee_id % 2 = 1 and left(name, 1) <> 'M', salary, 0) as bonus
from Employees