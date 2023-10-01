function createCounter(n: number): () => number {
  let i = n;
  return function () {
    const res = i;
    i++;
    return res;
  };
}

/**
 * const counter = createCounter(10)
 * counter() // 10
 * counter() // 11
 * counter() // 12
 */
