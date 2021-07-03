# Matrix

## Intuition

Consider:

* If the goal is to find a region in the matrix, try to merge rows into a single row to allow for 1d operations.
* Similarily, combining all rows into a single long row by "tying the next row to the end of the current", it becomes easy to use DSU.