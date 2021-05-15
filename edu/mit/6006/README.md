# MIT 6.006 Introduction to Algorithms

https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-006-introduction-to-algorithms-fall-2011/index.htm

## Unit 1: Introduction

### Algorithmic thinking, asymptotic complexity, peak finding

[Lecture](https://www.youtube.com/watch?v=HtSuA80QTyo)
[Recitation](https://www.youtube.com/watch?v=P7frcB_-g4w&)

Summary: introduction to algorithms in general. 

Time complexity:

* Upper-bound: $O(f(n)) = g(n)$, where $g(n)$ is the highest possible value $f(n)$ can take as $n \rightarrow \infty$
* Lower-bound: $\Omega(f(n)) = g(n)$ where $g(n)$ is lowest possible value of $f(n)$ as $n \rightarrow \infty$
* Average: $\Theta(f(n)) = g(n)$ where $g(n)$ is the average or most likely value of $f(n)$ as $n \rightarrow \infty$

Running time: $T(n)$

Peak finding using divide and conquer for the problem: "There is a single peak (or no peak) in a list of numbers, find it (if it exists)"
