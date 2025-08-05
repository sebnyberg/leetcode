package p3477fruitsintobasketsii

func numOfUnplacedFruits(fruits []int, baskets []int) int {
    var res int
outer:
    for _, x := range fruits {
        for i, y := range baskets {
            if y >= x {
                baskets[i] = -1
                continue outer
            }
        }
        res++
    }
    return res
}