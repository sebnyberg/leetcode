create table if not exists Transactions (account_id int, day date, type nvarchar(255), amount int)
truncate table Transactions
insert into Transactions (account_id, day, type, amount) values ('1', '2021-11-07', 'Deposit', '2000')
insert into Transactions (account_id, day, type, amount) values ('1', '2021-11-09', 'Withdraw', '1000')
insert into Transactions (account_id, day, type, amount) values ('1', '2021-11-11', 'Deposit', '3000')
insert into Transactions (account_id, day, type, amount) values ('2', '2021-12-07', 'Deposit', '7000')
insert into Transactions (account_id, day, type, amount) values ('2', '2021-12-12', 'Withdraw', '7000')

with diff as (
    select
        *,
        IIF(type = 'Withdraw', -amount, amount) as delta
    from Transactions
)
select
    account_id,
    day,
    sum(delta) over (partition by account_id order by day) as balance
from diff
order by account_id, day