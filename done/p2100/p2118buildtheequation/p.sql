select concat(trim(string_agg(
    case
        when power = 0 and factor >= 0 then concat('+', factor)
        when power = 0 and factor < 0 then cast(factor as char)
        when power = 1 and factor >= 0 then concat('+', factor, 'X')
        when power = 1 and factor < 0 then concat(factor, 'X')
        when factor >= 0 then concat('+', factor, 'X^', power)
        when factor < 0 then concat(factor, 'X^', power)
    end
    , '') within group (order by power desc)), '=0') as equation
from Terms
