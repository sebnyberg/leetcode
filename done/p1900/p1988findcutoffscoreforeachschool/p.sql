select top 1 with ties
    school_id,
    coalesce(score, -1) as score
from Schools
    left join Exam on student_count <= capacity
order by row_number() over (partition by school_id order by student_count desc, score asc)
