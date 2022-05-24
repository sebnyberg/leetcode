package p2264largest3samedigitnumberinstring

var want = map[string]bool{
	"000": true,
	"111": true,
	"222": true,
	"333": true,
	"444": true,
	"555": true,
	"666": true,
	"777": true,
	"888": true,
	"999": true,
}

func largestGoodInteger(num string) string {
	res := ""
	for i := 0; i < len(num)-2; i++ {
		if want[num[i:i+3]] && num[i:i+3] > res {
			res = num[i : i+3]
		}
	}
	return res
}
