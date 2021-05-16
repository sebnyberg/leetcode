\documentclass{article}
\usepackage{blindtext}
\usepackage[T1]{fontenc}
\usepackage[utf8]{inputenc}

\title{Sections and Chapters}
\author{Gubert Farnsworth}
\date{\today}

\begin{document}

\maketitle

\section{Introduction}

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

Machines have one word per cell in its memory.

* Random Access Machine is a big chunk o' memory
* Pointer Machine uses objects instead which can point to other objects

#### Document distance problem

Given two documents (d1, d2), figure out a distance $d(D_1, D_2)$.

A document is a sequence of words. A word is a string of alphanumerical characters.

Idea: shared words. A vector of words - D[w] = # occurrences of w in D. The goal is to find a measure of similarity based on the vectors. One example would be the dot-product, but that causes unnormalized scores. Another is cosine-similarity, which is given by

$$
\newcommand{\d1}{\textrm{D}_1}
\textrm{d}'(\textrm{D}_1, \textrm{D}_2) = \frac{\textrm{D}_1}{\left|\textrm{D}_1\right|}
$$

\end{document}