# Strong password checker

Observations:

* When len(password) < 6, inserts are required
* When len(password) > 20, removals are required
* Inserts and replacements can be used to meet lower/upper/digit constraints
* When eliminating repeating groups, replacements are more valuable than inserts
* When 6 <= len(password) <= 20, replacements are most valuable 
* Repeating groups should be eliminated as effectively as possible while inserting / removing characters

The observation that replacements are more effective comes from the fact that "aaaaa" can be fixed with a single replacement, whereas it would require two inserts. Replacements can reduce the repeat count by 3, inserts by 2, removals by 1:

| Repeats | Replacements | Inserts | Removals 
| --- | --- | --- | --- | 
| 3 | 1 | 1 | 1 |
| 4 | 1 | 1 | 2 |
| 5 | 1 | 2 | 3 |
| 6 | 2 | 2 | 4 |
| 7 | 2 | 3 | 5 |
| 8 | 2 | 3 | 6 |

Since replacements are more effective than inserts and removals, inserts or removals should only be used when necessary, i.e. when len(password) < 6 or > 20. And when insert or removal is required, the goal is to unlock efficient use of replacements afterwards. Basically, insert/remove should "land on" 5, 8, 11, ... `counts[i]%3 == 2`.

This gives the following prioritization for removal of characters while n > 20

1. Remove once in groups of 6, 9, 12, ..., i.e. `counts[i]%3 == 0`
2. (n>21) Remove twice in groups of 7, 10, 13, ..., i.e.  `counts[i]%3 == 1`
3. Repeated removal from group until <= 2
4. Remove without reducing repeated group

And for inserts, i.e. n < 6:

1. Insert once in groups of 7, 10, 13, i.e.  `counts[i]%3 == 1`
2. Insert in any group
3. Insert without reducing repeated group

Once 6 <= n <= 20, use replacements to fix any repeating groups.

Finally, count number of inserts/replacements to determine if more replacements are needed to match constraints.
