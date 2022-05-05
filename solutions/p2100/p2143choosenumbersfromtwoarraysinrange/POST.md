
This problem evaluates contiguous ranges from an array, which makes it a subarray-type problem. For subarray problems, it's good to consider:

* Queues. N/A for this problem.
* Prefix aggregates (sum). N/A for this problem.
* Aggregate (sum) counts. Applicable.
* Transformations of elements e.g. making some of the negative, taking the absolute value etc. Applicable.

Approach:

* Subtract from one array and add from the other, a balanced range must then sum to zero.
* Keep track of all possible prior sums and their counts using a map.
* Add each element on its own to the current sum, this 'starts' a new sum range in the current position

# Map-based solution ~300ms

<iframe src="https://leetcode.com/playground/HszNpwNR/shared" frameBorder="0" width="640" height="530"></iframe>

# Optimized solution (array) ~20ms

Use an array instead of a map.

Since numbers are negative, we must find a proper offset.

The maximum valid sum starts with 50 elements of `[100, 100, ...]`, then ends with 50 `[-100, -100, ...]` for a maximum sum of 5000 halfway through the array. Similarly, the minimum valid sum starts with 50 negative elements, then 50 positive for a minimum of -5000 halfway through the array. This gives us an offset of 5100.

The range of valid sums changes as the index changes. For early indices, the range is determined by the maximum / minimum possible sum. For higher indices, the range is determined by whether the remaining elements could pull back the sum to be balanced (zero).

<iframe src="https://leetcode.com/playground/LhUt2zDa/shared" frameBorder="0" width="640" height="1120"></iframe>