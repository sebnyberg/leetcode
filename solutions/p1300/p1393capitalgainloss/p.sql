with a as (
    select
        diff =
            case
                when operation = 'Buy' then -price
                else price
            end,
        stock_name
    from Stocks s
)
select
    stock_name,
    sum(diff) as capital_gain_loss
from a
group by stock_name
