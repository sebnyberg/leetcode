package p2353designafoodratingsystem

import (
	"container/heap"
	"testing"
)

func TestA(t *testing.T) {
	c := Constructor(
		[]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
		[]string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
		[]int{9, 12, 8, 15, 14, 7},
	)
	c.ChangeRating("sushi", 16)
	c.ChangeRating("ramen", 16)
}

type food struct {
	name    string
	cuisine string
	rating  int

	heapIdx int // keep track of food item's position in the heap
}

type FoodRatings struct {
	foods    map[string]*food
	cuisines map[string]*foodHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	// The idea is to use pointers and heap.Fix to manage updates to the
	// heap and ratings of a given food.
	var r FoodRatings
	r.cuisines = make(map[string]*foodHeap)
	r.foods = make(map[string]*food)
	for i := range foods {
		if _, exists := r.cuisines[cuisines[i]]; !exists {
			r.cuisines[cuisines[i]] = &foodHeap{}
		}
		f := &food{
			name:    foods[i],
			cuisine: cuisines[i],
			rating:  ratings[i],
		}
		heap.Push(r.cuisines[cuisines[i]], f)
		r.foods[foods[i]] = f
	}
	return r
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	f := this.foods[food]
	f.rating = newRating
	heap.Fix(this.cuisines[f.cuisine], f.heapIdx)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return (*this.cuisines[cuisine])[0].name
}

type foodHeap []*food

func (h foodHeap) Len() int { return len(h) }
func (h foodHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx = i
	h[j].heapIdx = j
}
func (h foodHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].name < h[j].name
	}
	return h[i].rating > h[j].rating
}
func (h *foodHeap) Push(x interface{}) {
	a := x.(*food)
	a.heapIdx = len(*h)
	*h = append(*h, a)
}
func (h *foodHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
