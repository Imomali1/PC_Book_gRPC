package sample

import (
	"github.com/Imomali1/gRPC/pcbook/pb"
	"math/rand"
)

func RandKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func RandCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func RandCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"XEON E-2286M",
			"Core i9-1010HK",
			"Core i7-8328H",
			"Core i5-1012F",
			"Core i3-5535W",
		)
	} else {
		return randomStringFromSet(
			"Ryzen 7 PRO 2700U",
			"Ryzen 5 PRO 3113E",
			"Ryzen 3 PRO 3242W",
		)
	}
}

func randomStringFromSet(s ...string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	return s[rand.Intn(n)]
}

func RandBool() bool {
	return rand.Intn(2) == 1
}

func RandInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
