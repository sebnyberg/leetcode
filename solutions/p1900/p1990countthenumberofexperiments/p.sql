with groups as (
    select *
    from (values ('IOS'),('Android'),('Web')) a(platform)
    cross join (values ('Programming'),('Sports'),('Reading')) b(experiment_name)
), counts as (
    select platform, experiment_name, count(*) as num_experiments
    from Experiments
    group by platform, experiment_name
)
select
		g.platform,
		g.experiment_name,
		iif(c.num_experiments > 0, c.num_experiments, 0) as num_experiments
from groups g
left join counts c
    on g.platform = c.platform
    and g.experiment_name = c.experiment_name
