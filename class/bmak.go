package class

import (
	"fmt"
	"strconv"
	"time"

	"nike.bot/akamai_reverse_eng/constants"
)

type bmak struct {
	Ver        string
	Session_id string
	UserAgent  string
	Z1         int
	Y1         int
	StartTime  int64
	ProductSub string
	Brave      string
	Language   string
	Product    string
	Plen       string
	Xagg       string
	Pen        string
	Wen        string
	Den        string
	D3         string
	Ins        string
	Cns        string
	InformInfo string
	Kact       string
	Mact       string
	Tact       string
	Doact      string
	Dmact      string
	Pact       string
	Vcact      string
	Url        string
}

func BmakConstructor() *bmak {
	bmakObj := new(bmak)
	bmakObj.Ver = "1.68"
	bmakObj.Y1 = 2016
	bmakObj.Session_id = "default_session"
	bmakObj.UserAgent = constants.RandomUserAgent()
	bmakObj.StartTime = now()                     //equivalent as Date.now() from javascript
	bmakObj.ProductSub = "20030107"               //safari and chrome default build number
	bmakObj.Brave = "0"                           //brave navigator default to 0
	bmakObj.Language = "pt-BR"                    //navigator language default to portugueses
	bmakObj.Product = "Gecko"                     //all browsers default to Gecko
	bmakObj.Plen = constants.RandomPluginLength() //number of plugins on browser, random between 1 and 5
	bmakObj.Xagg = bmakObj.Bc()
	bmakObj.Pen = "0" //doesnt use phantomJs
	bmakObj.Wen = "0" //doesnt use webdriver
	bmakObj.Den = "0" //doesnt use dom automation
	bmakObj.D3 = "0"
	bmakObj.InformInfo = "0,0,0,0,1197,113,0;0,-1,0,0,2075,2075,0;0,1,0,0,2835,2835,0;0,-1,0,0,2963,2963,0;0,-1,0,0,2189,2189,0;0,-1,0,0,2906,2906,0;0,-1,0,0,2686,2686,0;0,-1,0,0,2903,2903,0;0,-1,0,0,2074,2074,0;0,-1,0,0,2599,2599,0;0,-1,0,0,2424,2424,0;0,-1,0,0,2949,2949,0;0,-1,0,0,2401,2401,0;0,-1,0,0,2364,2364,0;0,-1,0,0,2402,2402,0;0,-1,0,0,700,399,0;0,-1,0,0,789,488,0;"
	bmakObj.Cns = "0;0;0;0;0;0;0;0;0;0;0;1;0;0;0;0;0;0;0;0;0;0;0;0;0;0;0;0;1;1;"
	bmakObj.Ins = "0;0;0;0;0;0;0;0;0;0;0;1;0;0;0;0;0;0;0;0;0;0;0;0;0;0;0;0;1;1;"
	bmakObj.Kact = ""
	bmakObj.Mact = ""
	bmakObj.Tact = ""
	bmakObj.Doact = ""
	bmakObj.Dmact = ""
	bmakObj.Pact = ""
	bmakObj.Vcact = ""
	bmakObj.Url = "https://unite.nike.com.br/oauth.html?client_id=QLegGiUU042XMAUWE4qWL3fPUIrpQTnq&redirect_uri=https%3A%2F%2Fwww.nike.com.br%2Fapi%2Fv2%2Fauth%2Fnike-unite%2Fset&response_type=code&locale=pt_BR"

	return bmakObj
}

func (b bmak) Ab() int {
	a := 0
	for e := 0; e < len(b.UserAgent); e++ {
		var n = int(b.UserAgent[e])
		if n < 128 {
			a += n
		}
	}

	return a
}

func (b bmak) Bc() string {
	t := 1
	a := 1
	e := 0
	n := 0
	o := 1
	m := 1
	r := 1
	i := 0
	c := 1
	b2 := 1
	d := 0
	s := 1
	k := 1
	l := 1

	return strconv.Itoa(t + (a << 1) + (e << 2) + (n << 3) + (o << 4) + (m << 5) + (r << 6) + (i << 7) + (k << 8) + (l << 9) + (c << 10) + (b2 << 11) + (d << 12) + (s << 13))
}

func (b bmak) Bd() string {
	return ",cpen:0,i1:0,dm:0,cwen:0,non:1,opc:0,fc:0,sc:0,wrc:1,isc:0,vib:1,bat:1,x11:0,x12:1"
}

func (b bmak) Gd() string {
	var t = b.UserAgent
	var a = strconv.Itoa(b.Ab()) //converts value to string
	var e = strconv.Itoa(int(b.StartTime / 2))
	var n = constants.RandomScreenWidth()
	var o = constants.RandomScreenHeight()
	var m = n
	var r = o
	var i = o
	var c = n
	var b2 = n

	b.Z1 = int(int(b.StartTime) / (b.Y1 * b.Y1))

	var d = constants.RandomNumber()
	var s = strconv.Itoa(int((1e3 * d / 2)))
	var k = strconv.FormatFloat(d, 'E', -1, 64)

	k = k[0:11] + s

	fmt.Println("rasda", b.D3)

	return t + ",uaend," + b.Xagg + "," + b.ProductSub + "," + b.Language + "," + b.Product + "," + b.Plen + "," + b.Pen + "," + b.Wen + "," + b.Den + "," + strconv.Itoa(b.Z1) + "," + b.D3 + "," + n + "," + o + "," + m + "," + r + "," + c + "," + i + "," + b2 + "," + b.Bd() + "," + a + "," + k + "," + e + "," + b.Brave + ",loc:"
}

func (b bmak) X2() int64 {
	return now()
}

func now() int64 {
	return (time.Now().UTC().UnixNano() / 1e6)
}

func (b *bmak) To() {
	b.D3 = strconv.Itoa(int(b.X2() % 1e7))
}

func (b bmak) getFormInfo() string {
	return "0,-1,0,0,2075,2075,0;0,1,0,0,2835,2835,0;0,-1,0,0,2963,2963,0;0,-1,0,0,2189,2189,0;0,-1,0,0,2906,2906,0;0,-1,0,0,2686,2686,0;0,-1,0,0,2903,2903,0;0,-1,0,0,2074,2074,0;0,-1,0,0,2599,2599,0;0,-1,0,0,2424,2424,0;0,-1,0,0,2949,2949,0;0,-1,0,0,2401,2401,0;0,-1,0,0,2364,2364,0;0,-1,0,0,2402,2402,0;0,-1,0,0,700,399,0;0,-1,0,0,789,488,0;"
}

func (b bmak) updateTime() string {
	return strconv.Itoa(int(now()) - int(b.StartTime))
}

func (b bmak) GetSensorData() string {
	return b.Ver +
		"-1,2,-94,-100," +
		b.Gd() + "-1,2,-94,-101," +
		"do_en" + "dm_en" + "t_en" +
		"-1,2,-94,-105," +
		b.InformInfo +
		"-1,2,-94,-102," +
		b.getFormInfo() +
		"-1,2,-94,-108," +
		b.Kact +
		"-1,2,-94,-110," +
		b.Mact +
		"-1,2,-94,-117," +
		b.Tact +
		"-1,2,-94,-111," +
		b.Doact +
		"-1,2,-94,-109," +
		b.Dmact +
		"-1,2,-94,-114," +
		b.Pact +
		"-1,2,-94,-103," +
		b.Vcact +
		"-1,2,-94,-112," +
		b.Url +
		"-1,2,-94,-115," +
		v +
		"-1,2,-94,-106," +
		d
}
