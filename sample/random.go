package sample

import (
	"math/rand"
	"time"

	"main.go/pb"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY

	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Apple Silicon", "AMD", "Intel")
}

func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet(
			"RTX 3090",
			"RTX 3080",
			"RTX 3070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
	)
}

func randomCPUName(cpuBrand string) string {
	if cpuBrand == "Intel" {
		return randomStringFromSet(
			"Core i5",
			"Core i7",
			"Xeon",
			"Core i9",
		)
	}
	if cpuBrand == "AMD" {
		return randomStringFromSet(
			"Ryzen 7",
			"Ryzen 9",
			"Ryzen 5",
		)
	}

	return randomStringFromSet(
		"M1",
		"M1 PRO",
		"M1 PRO MAX",
		"M2",
		"M2 PRO",
		"M2 PRO MAX",
	)

}

func randomScreenResolution() *pb.Screen_Resolution {

	heigth := randomInt(1000, 4320)
	width := heigth * 16 / 9

	screenResolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Heigth: uint32(heigth),
	}
	return screenResolution
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "ThinkPad", "MSi", "ASUS")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Pro", "Macbook Air")
	case "ThinkPad":
		return randomStringFromSet("x230", "x240")
	case "MSi":
		return randomStringFromSet("Carbon")
	default:
		return randomStringFromSet("Zen Book", "GameBOOK5000NAGIBATOR")
	}

}

func randomStringFromSet(stringsSet ...string) string {
	n := len(stringsSet)
	return stringsSet[rand.Intn(n)]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomId() string {
	return uuid.New().String()
}
