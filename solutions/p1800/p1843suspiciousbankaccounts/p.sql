package p1843suspiciousbankaccounts
with something as (
    select a.account_id,
           month(day) as mo,
           sum(amount) as amt
    from Transactions t
    join Accounts a on t.account_id = a.account_id
    where type = 'Creditor'
    group by a.account_id, max_income, month(day)
    having sum(amount) > max_income
)
select distinct s1.account_id
from something s1 join something s2 on s1.account_id = s2.account_id
where s1.mo = s2.mo-1
