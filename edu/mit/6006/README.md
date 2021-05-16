# MIT 6.006 Introduction to Algorithms

[Course Website](https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-006-introduction-to-algorithms-fall-2011/index.htm)

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

### Models of computation 

[Lecture](https://www.youtube.com/watch?v=Zc54gFhdpLA)
[Recitation](https://www.youtube.com/watch?v=QFcyt8fgQMU)

Summary: 

Model of computation specifies:
* What operations an algorithm is allowed
* Cost (time, space) of each op

#### Random Access Machine (RAM)

Modeled by a big array from 1 .. words in one stick of RAM.

A word is a cell in memory, and it has width $w$.

#### Pointer Machine

- OOP
- Dynamically allocated objects
- Object has O(1) fields
- Field or pointer or null = word (e.g. int) 