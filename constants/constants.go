package constants

import (
	"math/rand"
	"strconv"
	"time"
)

type constants struct {
	userAgents  [5]string
	widthSizes  [5]string
	heightSizes [5]string
}

func CreateConstants() *constants {
	constantsObj := new(constants)

	constantsObj.userAgents = [5]string{
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/37.0.2062.94 Chrome/37.0.2062.94 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/600.8.9 (KHTML, like Gecko) Version/8.0.8 Safari/600.8.9",
	}

	constantsObj.widthSizes = [5]string{
		"1920",
		"1440",
		"1336",
		"1280",
		"968",
	}

	constantsObj.heightSizes = [5]string{
		"1080",
		"1040",
		"768",
		"960",
		"460",
	}

	return constantsObj
}

func RandomNumber() float64 {
	rand.Seed(time.Now().Unix())

	return rand.Float64()
}

func RandomUserAgent() string {
	c := CreateConstants()

	rand.Seed(time.Now().Unix())

	return c.userAgents[rand.Intn(len(c.userAgents))]
}

func RandomScreenWidth() string {
	c := CreateConstants()

	rand.Seed(time.Now().Unix())

	return c.widthSizes[rand.Intn(len(c.widthSizes))]
}

func RandomScreenHeight() string {
	c := CreateConstants()

	rand.Seed(time.Now().Unix())

	return c.heightSizes[rand.Intn(len(c.heightSizes))]
}

func RandomPluginLength() string {
	rand.Seed(time.Now().Unix())

	min := 1
	max := 5

	return strconv.Itoa(rand.Intn(max-min+1) + min)
}
