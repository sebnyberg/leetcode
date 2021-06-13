package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntervalTree(t *testing.T) {
	var tree IntervalTree
	tree.Insert(17, 19)
	tree.Insert(15, 20)
	tree.Insert(10, 30)
	tree.Insert(5, 20)
	tree.Insert(1, 2)
	tree.Insert(12, 15)
	tree.Insert(30, 40)
	tree.Insert(60, 70)

	for _, tc := range []struct {
		start, end int
		want       bool
	}{
		{3, 6, true},
		{3, 3, false},
		{39, 41, true},
		{41, 45, false},
		{45, 65, true},
	} {
		t.Run(fmt.Sprintf("[%v,%v)", tc.start, tc.end), func(t *testing.T) {
			require.Equal(t, tc.want, tree.IntersectsInterval(tc.start, tc.end))
		})
	}
}

func TestIntervalTreeLarge(t *testing.T) {
	for _, tc := range []struct {
		intervals   [][]int
		noIntersect []bool
	}{
		{
			[][]int{
				{69, 70}, {3, 4}, {39, 40}, {35, 36}, {3, 4}, {55, 56}, {61, 62}, {97, 98}, {79, 80}, {76, 77}, {46, 47}, {78, 79}, {47, 48}, {38, 39}, {83, 84}, {90, 91}, {90, 91}, {49, 50}, {49, 50}, {77, 78}, {23, 24}, {89, 90}, {8, 9}, {3, 4}, {2, 3}, {48, 49}, {96, 97}, {4, 5}, {54, 55}, {30, 31}, {97, 98}, {65, 66}, {93, 94}, {49, 50}, {24, 25}, {17, 18}, {53, 54}, {45, 46}, {53, 54}, {32, 33}, {37, 38}, {5, 6}, {50, 51}, {48, 49}, {14, 15}, {91, 92}, {79, 80}, {73, 74}, {28, 29}, {31, 32}, {98, 99}, {37, 38}, {19, 20}, {49, 50}, {54, 55}, {37, 38}, {98, 99}, {12, 13}, {24, 25}, {46, 47}, {74, 75}, {87, 88}, {64, 65}, {61, 62}, {68, 69}, {28, 29}, {43, 44}, {89, 90}, {64, 65}, {72, 73}, {69, 70}, {88, 89}, {68, 69}, {28, 29}, {20, 21}, {64, 65}, {17, 18}, {40, 41}, {88, 89}, {22, 23}, {8, 9}, {33, 34}, {13, 14}, {19, 20}, {53, 54}, {99, 100}, {24, 25}, {82, 83}, {77, 78}, {90, 91}, {72, 73}, {33, 34}, {73, 74}, {0, 1}, {25, 26}, {69, 70}, {73, 74}, {12, 13}, {33, 34}, {47, 48}, {26, 27}, {77, 78}, {95, 96}, {28, 29}, {77, 78}, {28, 29}, {87, 88}, {16, 17}, {42, 43}, {51, 52}, {44, 45}, {63, 64}, {24, 25}, {18, 19}, {0, 1}, {45, 46}, {65, 66}, {21, 22}, {37, 38}, {77, 78}, {97, 98}, {24, 25}, {83, 84}, {20, 21}, {29, 30}, {66, 67}, {29, 30}, {37, 38}, {63, 64}, {15, 16}, {85, 86}, {61, 62}, {0, 1}, {23, 24}, {96, 97}, {91, 92}, {90, 91}, {80, 81}, {18, 19}, {69, 70}, {3, 4}, {59, 60}, {21, 22}, {75, 76}, {54, 55}, {65, 66}, {34, 35}, {19, 20}, {79, 80}, {6, 7}, {24, 25}, {29, 30}, {35, 36}, {9, 10}, {0, 1}, {73, 74}, {65, 66}, {78, 79}, {32, 33}, {58, 59}, {25, 26}, {3, 4}, {78, 79}, {92, 93}, {37, 38}, {91, 92}, {5, 6}, {79, 80}, {94, 95}, {78, 79}, {38, 39}, {16, 17}, {81, 82}, {34, 35}, {16, 17}, {33, 34}, {42, 43}, {34, 35}, {89, 90}, {88, 89}, {33, 34}, {68, 69}, {92, 93}, {73, 74}, {64, 65}, {91, 92}, {44, 45}, {13, 14}, {97, 98}, {64, 65}, {31, 32}, {91, 92}, {1, 2}, {57, 58}, {21, 22}, {38, 39}, {70, 71}, {84, 85}, {50, 51}, {58, 59},
			},
			[]bool{
				true, true, true, true, false, true, true, true, true, true, true, true, true, true, true, true, false, true, false, true, true, true, true, false, true, true, true, true, true, true, false, true, true, false, true, true, true, true, false, true, true, true, true, false, true, true, false, true, true, true, true, false, true, false, false, false, false, true, false, false, true, true, true, false, true, false, true, false, false, true, false, true, false, false, true, false, false, true, false, true, false, true, true, false, false, true, false, true, false, false, false, false, false, true, true, false, false, false, false, false, true, false, true, false, false, false, false, true, true, true, true, true, false, true, false, false, false, true, false, false, false, false, false, false, true, true, false, false, false, true, true, false, false, false, false, false, false, true, false, false, false, true, false, true, false, false, true, false, false, true, false, false, false, true, false, false, false, false, false, true, false, false, false, true, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, true, true, false, false,
			},
		},
		{
			[][]int{
				{12, 20}, {21, 29}, {11, 20}, {12, 17}, {84, 90}, {60, 68}, {88, 94}, {23, 32}, {88, 94}, {15, 20}, {77, 83}, {34, 42}, {44, 53}, {35, 40}, {24, 31}, {48, 55}, {0, 6}, {6, 13}, {58, 63}, {15, 23},
			},
			[]bool{
				true, true, false, false, true, true, false, false, false, false, true, true, true, false, false, false, true, false, false, false,
			},
		},
		{
			[][]int{
				{40, 49}, {40, 49}, {49, 50}, {49, 50}, {27, 34}, {23, 30}, {39, 46}, {8, 15}, {3, 9}, {2, 8}, {48, 50}, {46, 50}, {4, 12}, {4, 10}, {30, 36}, {47, 50}, {15, 23}, {43, 50}, {49, 50}, {24, 33}, {17, 26}, {3, 11}, {45, 50}, {3, 8}, {32, 40}, {37, 43}, {5, 13}, {0, 9}, {48, 50}, {14, 22},
			},
			[]bool{
				true, false, true, false, true, false, false, true, false, true, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false,
			},
		},
		{
			[][]int{
				{20, 29}, {13, 22}, {44, 50}, {1, 7}, {2, 10}, {14, 20}, {19, 25}, {36, 42}, {45, 50}, {47, 50}, {39, 45}, {44, 50}, {16, 25}, {45, 50}, {45, 50}, {12, 20}, {21, 29}, {11, 20}, {12, 17}, {34, 40}, {10, 18}, {38, 44}, {23, 32}, {38, 44}, {15, 20}, {27, 33}, {34, 42}, {44, 50}, {35, 40}, {24, 31},
			},
			[]bool{
				true, false, true, true, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
			},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.intervals), func(t *testing.T) {
			var tree IntervalTree
			for i, interval := range tc.intervals {
				nointersect := !tree.IntersectsInterval(interval[0], interval[1])
				require.Equal(t, tc.noIntersect[i], nointersect, i)
				if nointersect {
					tree.Insert(interval[0], interval[1])
				}
			}
		})
	}
}
