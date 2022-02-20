# Mod

## Basics

```txt
1 mod 9 = 1
10 mod 9 = 1
a*c + b mod c = b
```

## Adding large numbers

If modulo is prime, then

```txt
a % mod + b % mod = a*b % mod
a % mod - b % mod = (a % mod + mod - b % mod) % mod
```

## Prime power

Given a prime modulo, any number to the power of that prime is equal to the base
mod p.

```txt
98214^7 % 7 = 98214 % 7
```

## Fermat's little theorem

```txt
a^(p-1) mod p = 1
```

For example:

```txt
3^6 % 7 = 1
```

## GCD

Any number is a combination of prime factors.

Taking the GCD of two numbers is equivalent to finding common prime factors
between the numbers. If no such numbers exist, the GCD is 1.

If we wish to find combinations of numbers `(a,b)` such that `a*b%k == 0`, then
we can try `gcd(a,k)*gcd(b,k)`. Why? Because `gcd(a,k)` will find common prime
factors between a and k, and vice versa for b and k, then multiplying the two
together will only have a rest of zero if a and b combined contain all prime
factors within k.
