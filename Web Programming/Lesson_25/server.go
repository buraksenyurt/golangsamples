/*
	Lesson 25
	HttpRouter isimli yönlendirici yardımıyla bir web sayfası geliştiriyoruz.

*/
package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter" // Kodu LiteIDE'de bu şekilde yazdıktan sonra Debug menüsünden Get dersek, github'daki ilgili paket otomatik olarak sisteme yüklenir
)

func main() {
	router := httprouter.New()              // Önce Router nesnesini örnekliyoruz.
	router.GET("/", Index)                  // root adrese gelecek olan HTTP Get taleplerini Index isimli fonksiyona yönlendiriyoruz.
	router.GET("/planets", GetPlanets)      // root adresten /planets şeklinde gelecek talepleri GetPlanets fonksiyonuna yönlendiriyoruz.
	router.GET("/planets/:name", GetCities) //bu sefer planets arkasına bir de parametre aldık. Bu parametre tahmin edileceği gibi URL'den gelecek
	http.ListenAndServe(":4568", router)    // 4568 nolu porttan yapacağımız dinlemelerde, gelen tüm talepleri httpRouter nesne örneğinin ele alacağını belirtiyoruz.
}

// Bu sefer bir gezegendeki şehirleri listeyeceğiz
// params üzerinden gezegen adını yakalayıp slice'dan onun şehirlerini bulacağız
func GetCities(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	planetName := params.ByName("name")
	planets := LoadSomeData()
	fmt.Fprintf(response, `<html><head><title>%s</title></head><body><h1>%s</h1>`, planetName, planetName)
	for _, planet := range planets {
		if strings.ToLower(planet.Name) == strings.ToLower(planetName) {
			for _, city := range planet.Cities {
				fmt.Fprintf(response, `<p><b>%s</b>-%s</p>`, city.Name, city.Affiliation)
			}
			break
		}
	}
	fmt.Fprintf(response, `</body></html>`)
}

func GetPlanets(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	planets := LoadSomeData() //Star Wars'dan bir kaç gezegen ve şehir adı yüklenir
	fmt.Fprintf(response, `<html><head><title>Planets</title></head><body><h1>Planets</h1>`)
	for _, planet := range planets {
		fmt.Fprintf(response, `<p><a href='planets/%s'>%s-%s (%d)</p>`, planet.Name, planet.Name, planet.Sector, planet.Population)
	}
	fmt.Fprintf(response, `</body></html>`)
}

// builtin kullandığımız http.HandleFunc'dan farklı olarak 3ncü bir parametremiz var.
// httpRouter.Params ile talep ile birlikte gelen parametreleri yakalayabiliriz
func Index(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(response, `<html>
		<body>
		<head>
			<title>Star Wars Planets</title>
		</head>
		<body>
			<h1>Star Wars Planets</h1>
			<a href="/planets">Planets</a><br/>
			<p>Planet list updates every morning with new planets</p>
		</body>
		</html>`)
	//HTML içeriğini bir dosyadan okutursak daha güzel olur. Ama dinamik olarak da üretilebilir
}

// Örnek verileri yüklemek için kullanacağımız fonksiyon
// Geriye bir slice döndürecek
func LoadSomeData() []Planet {
	var planets []Planet

	planets = append(planets, Planet{Name: "Naboo", Sector: "Chommel", Population: 4500000,
		Cities: []City{
			City{Id: 1, Name: "Theed", Affiliation: "Galactic Empire"},
			City{Id: 2, Name: "Umberbool City", Affiliation: "Gungan Grand Army"},
			City{Id: 3, Name: "Spinnaker", Affiliation: "Galactic Empire"},
			City{Id: 4, Name: "Otoh Gunga", Affiliation: "Trade Federation"},
		}})

	planets = append(planets, Planet{Name: "Coruscant", Sector: "Corusca", Population: 1000000000,
		Cities: []City{
			City{Id: 1, Name: "Galactic City", Affiliation: "Rebellian"},
		}})

	planets = append(planets, Planet{Name: "Mustafar", Sector: "Atravis", Population: 20000,
		Cities: []City{
			City{Id: 1, Name: "Fralideja", Affiliation: "Rise of Empire"},
		}})

	return planets
}

type Planet struct {
	Name       string
	Sector     string
	Population int64
	Cities     []City
}

type City struct {
	Id          int
	Name        string
	Affiliation string
}
