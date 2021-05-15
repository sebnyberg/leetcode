# 1 - Algorithmic Thinking, Peak Finding

## Peak finding

A position b is a peak if and only if b >= a and b >= c

Worst-case complexity is O(n)

Assuming that there is only one peak (if there is one), then divide & concuqer can be used to find the peak.

### 2D Peak finding

a is a 2d-peak iff a >= b, a >= d, a >= c, a >= e, where b,c,d,e are the manhattan neighbours of a.

Greedy ascent (find neighbour greater than current and move in that direction) - O(n*m) or O(n^2) if m = n.

#### Divide & conquer

Pick middle column. j := m/2
Find global maximum on col j at (i,j)
Compare (i, j-1), (i, j), (i, j+1) 
Pick left cols if (i, j-1) > (i, j) (similar for right)
When there is a single column, find the global max <- done

T(n, m) = T(n, m/2) + O(n)
T(n, 1) = O(n)
T(n, m) = O(n) + ... + O(n) = O(n*log2(m))

Omega(x) = lower-bound
O(x) = upper-bound
Theta(x) = asymptotic bound