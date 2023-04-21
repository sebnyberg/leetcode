-- For a triangle to be valid, the largest of its sides must be shorter than
-- half of its total length
select x, y, z,
       case when x >= y + z
                or y >= x + z
                or z >= x + y
            then 'No'
            else 'Yes'
        end as triangle
from Triangle
