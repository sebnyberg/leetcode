# Testing the solution

Below is an example of how "fill in the blanks" type of exercises can be tested.

If there are test case failures for very large inputs, then it's necessary to put test data as separate input files and load + transform them into typed cases.

```go
func TestBank(t *testing.T) {
	type any interface{}
	type actionName string
	const (
		actionWithdraw = "withdraw"
		actionDeposit  = "deposit"
		actionTransfer = "transfer"
	)
	type testCaseAction struct {
		name actionName
		args []any
		want bool
	}
	type testCase struct {
		name    string
		balance []int64
		actions []testCaseAction
	}
	testCases := []testCase{
		{
			"example",
			[]int64{10, 100, 20, 50, 30},
			[]testCaseAction{
				{actionWithdraw, []any{3, int64(10)}, true},
				{actionTransfer, []any{5, 1, int64(20)}, true},
				{actionDeposit, []any{5, int64(20)}, true},
				{actionTransfer, []any{3, 4, int64(15)}, false},
				{actionWithdraw, []any{10, int64(50)}, false},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf(tc.name, tc.balance), func(t *testing.T) {
			b := Constructor(tc.balance)
			for _, act := range tc.actions {
				var got bool
				switch act.name {
				case actionWithdraw:
					got = b.Withdraw(act.args[0].(int), act.args[1].(int64))
				case actionTransfer:
					got = b.Transfer(act.args[0].(int), act.args[1].(int), act.args[2].(int64))
				case actionDeposit:
					got = b.Deposit(act.args[0].(int), act.args[1].(int64))
				}
				require.Equal(t, act.want, got)
			}
		})
	}
}
```

# Solution

```go
type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	b := Bank{
		balance: make([]int64, len(balance)+1),
	}
	copy(b.balance[1:], balance)
	b.balance[0] = math.MinInt32
	return b
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.validAccount(account1) ||
		!this.validAccount(account2) ||
		this.balance[account1] < money {
		return false
	}
	this.balance[account1] -= money
	this.balance[account2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.validAccount(account) {
		return false
	}
	this.balance[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.validAccount(account) || this.balance[account] < money {
		return false
	}
	this.balance[account] -= money
	return true
}

func (this *Bank) validAccount(id int) bool {
	return id >= 1 && id <= len(this.balance)
}
```