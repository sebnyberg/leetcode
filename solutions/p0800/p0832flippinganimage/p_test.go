package p0832flippinganimage

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image[0])
	for i := range image {
		for l, r := 0, n-1; l <= r; l, r = l+1, r-1 {
			image[i][l], image[i][r] = image[i][r], image[i][l]
			image[i][l] = 1 - image[i][l]
			if l != r {
				image[i][r] = 1 - image[i][r]
			}
		}
	}
	return image
}
