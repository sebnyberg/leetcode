This post will go through optimization steps toward a solution that is near-optimal for this problem, and optimal for the followup problem: 188.

## First solution

Form a matrix `profits[k][i]` where the value of a cell at `profits[k][i]` is the maximum profit given by `k+1` possible trades at the time `i`. 

```go
ntrades := 2
ndays := len(prices)
profits := make([][]int, ntrades)
for i := range profits {
  profits[i] = make([]int, ndays)
}
```

Filling the first row is easy - the maximum profit at day `i` is the difference between the current price and the minimum price before day `i`:

```go
minVal := prices[0]
for i := 1; i < ndays; i++ {
  profits[0][i] = max(profits[0][i-1], prices[i]-minVal)
  minVal = min(minVal, prices[i])
}
```

Given `prices=[3,3,5,0,0,3,1,4]`, `profits[k][i]` becomes:

```go
  i 0 1 2 3 4 5 6 7
k \ - - - - - - - - 
0 | 0 0 2 2 2 3 3 4
```

For `k > 1`, the profit from making a trade depends on the maximum profit of previous trades. The easiest way to take this into account is to adjust the price on a given day by the maximum profit made from a trade made **before that day**.

For example, the adjusted price for the second trade would become:

```go
[ 3  3  5  0  0  3  1  4 ] prices
[ 0  0  2  2  2  3  3  4 ] profits[0]
[ 0  3  5 -2 -2  0 -2  1 ] adjusted = prices[i] - profits[i-1]
```

It is now clear that the value of making a trade on a given day for `k > 1` is given by the recurrence relation:

* `adjusted[k][i] = prices[i] - profits[k-1][j-1]`
* `profits[k][i] = max(profits[k][i-1], prices[i]-min(adjusted[k][:i]...)`

Which gives the first solution:

```go
for k := 1; k < ntrades; k++ {
  for i := 1; i < ndays; i++ {
    minPriceBeforeToday := adjustedPrices[0]
    for j := 1; j < i; j++ {
      minPriceBeforeToday = min(minPriceBeforeToday, adjustedPrices[j])
    }
    profits[k][i] = max(profits[k][i-1], prices[i]-minPriceBeforeToday)
    adjustedPrices[i] = prices[i] - profits[k-1][i-1]
  }
}
```

## Optimization 1: Keeping track of the minimum value 


Instead of finding the minimum value of `adjusted[k][:i]` for each index i, the minimum cost of buying a stock before the day i can be kept track of, i.e.:

```go
minPrice = [ 3  3  3 -2 -2 -2 -2 -2 ]
```

Which gives:

```go
minPrices := make([]int, ndays)
minPrices[0] = prices[0]
for k := 1; k < ntrades; k++ {
  for i := 1; i < ndays; i++ {
    minPrices[i] = min(minPrices[i-1], prices[i]-profits[k-1][i-1])
    profits[k][i] = max(profits[k][i-1], prices[i]-minPrices[i])
  }
}
```

## Optimization 2: using a single minimum value

Instead of keeping a list of minimum values, keep only the minimum value for any day before the current one.

```go
var minPrice int
for k := 1; k < ntrades; k++ {
  minPrice = prices[0]
  for i := 1; i < ndays; i++ {
    // Deducting the minimum price by the profits the day / trade before,
    // we artificially add value to to the current trade
    minPrice = min(minPrice, prices[i]-profits[k-1][i-1])
    profits[k][i] = max(profits[k][i-1], prices[i]-minPrice)
  }
}
```

## Optimization 3: using a single list of profits

By updating the minimum price and current place, it's possible to use a single row of profits rather than a matrix:

```go
ntrades := 2
ndays := len(prices)
profits := make([]int, ndays)

var minPrice int
for k := 0; k < ntrades; k++ {
  minPrice = prices[0]
  for i := 1; i < ndays; i++ {
    profits[i], minPrice = max(profits[i-1], prices[i]-minPrice),
      min(minPrice, prices[i]-profits[i])
  }
}

return profits[ndays-1]
```

## Optimization 4: skip days which cannot improve the profits

Finally, there is no point in updating the list of profits for days for which it is not possible to gain from a previous trade:

```go
ntrades := 2
ndays := len(prices)
profits := make([]int, ndays)

var minPrice int
for k := 0; k < min(ntrades, ndays/2); k++ {
  minPrice = prices[0]
  for i := 1; i <= k*2; i++ {
    minPrice = min(minPrice, prices[i]-profits[i])
  }
  for i := 1 + k*2; i < ndays; i++ {
    profits[i], minPrice = max(profits[i-1], prices[i]-minPrice),
      min(minPrice, prices[i]-profits[i])
  }
}

return profits[ndays-1]
```