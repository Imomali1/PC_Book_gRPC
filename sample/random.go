package sample

import (
	"github.com/Imomali1/gRPC/pcbook/pb"
	"github.com/google/uuid"
	"math/rand"
)

func randKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randCPUBrand() string {
	return randStringFromSet("Intel", "AMD")
}

func randCPUName(brand string) string {
	if brand == "Intel" {
		return randStringFromSet(
			"XEON E-2286M",
			"Core i9-1010HK",
			"Core i7-8328H",
			"Core i5-1012F",
			"Core i3-5535W",
		)
	} else {
		return randStringFromSet(
			"Ryzen 7 PRO 2700U",
			"Ryzen 5 PRO 3113E",
			"Ryzen 3 PRO 3242W",
		)
	}
}

func randGPUBrand() string {
	return randStringFromSet("NVIDIA", "AMD")
}

func randGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	} else {
		return randStringFromSet(
			"RX 590",
			"RX 580",
			"RX 5700-XT",
			"RX Vega-56",
		)
	}
}

func randScreenResolution() *pb.Screen_Resolution {
	height := randInt(1080, 4320)
	width := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}

	return resolution
}

func randScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randID() string {
	return uuid.New().String()
}

func randLaptopBrand() string {
	return randStringFromSet("HP", "Acer", "Lenovo", "Apple")
}

func randLaptopName(brand string) string {
	switch brand {
	case "HP":
		return randStringFromSet(
			"DragonFly G3",
			"NB-9007",
			"Pavilion Plus",
			"Victus 15",
			"Omen 16",
		)
	case "Acer":
		return randStringFromSet(
			"Predator Triton 500",
			"Spin 5",
			"Chromebook C933T",
			"Swift 3x",
			"Aspire 5",
		)
	case "Lenovo":
		return randStringFromSet(
			"Legion Slim",
			"V14 G2",
			"IdeaPad",
			"ThinkPad",
			"Yoga 7",
		)
	case "Apple":
		return randStringFromSet(
			"MacBook Pro",
			"MacBook Air",
			"iMac",
			"M1 Pro",
		)
	}
	return ""
}

func randStringFromSet(s ...string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	return s[rand.Intn(n)]
}

func randBool() bool {
	return rand.Intn(2) == 1
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
