/*
Lesson 26
	Bu sefer JSON çıktı veren REST tadından bir servis yazıyoruz
	sistemde paket olarak oluşturulan entity/starwars içerisindeki
	Category ve Model yapılarının çıktıları JSON formatında basılıyor
	Router için httpRouter paketini kullanıyoruz
*/
package main

import (
	"encoding/json"
	"entity/starwars" // bu paket C:\Go Works\Samples\src\entity\starwars altında yer alıyor
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New() //router nesnesi oluşturuluyor

	// static sayfaları bir sub-path'e yönlendiriyoruz. Eğer / kullanırsak *filepath bildirimi diğer Get bildirimlerini de kapsayacağından derleme hatası alırız.
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.GET("/category", GetCategories)                 // /category taleplerini GetCategories fonksiyonu karşılayacak
	router.GET("/category/:name", GetModelsByCategoryName) // /category/fighter gibi talepleri de GetModelsByCategoryName karşılayacak

	http.ListenAndServe(":4569", router) // 4569 nolu porttan dinleme yapıyoruz. İstekleri Router nesnemize yönlendiriyoruz
}

func GetCategories(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	c, _ := loadDataSet()                                     // örnek veri setini yüklüyoruz
	cJson, _ := json.Marshal(c)                               // kategorileri json formatına dönüştürüyoruz
	response.Header().Set("Content-Type", "application/json") // çıktının json olarak yorumlanması gerektiğini ifade ediyoruz
	response.WriteHeader(200)                                 // HTTP 200 OK bildirimini ekliyoruz
	fmt.Fprintf(response, "%s", cJson)                        //içeriği basıyoruz
}

func GetModelsByCategoryName(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, models := loadDataSet()
	var result []starwars.Model // istenen kategoriye bağlı modelleri tutacağımız slice
	for _, m := range models {
		if strings.ToLower(m.Category.Name) == strings.ToLower(params.ByName("name")) { //kategori adına uyan bir modelse ekliyoruz
			result = append(result, m)
		}
	}
	cJson, _ := json.Marshal(result) // bulunan sonuç kümesini json formatına çeviriyoruz
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	fmt.Fprintf(response, "%s", cJson)
}

func loadDataSet() (categories []starwars.Category, models []starwars.Model) {
	fighter := starwars.Category{Id: 1, Name: "Fighter"}
	cruiser := starwars.Category{Id: 2, Name: "Cruiser"}

	vwing := starwars.Model{Id: 1, Title: "V-Wing Fighter", Price: 45.50, Category: fighter}
	n1 := starwars.Model{Id: 2, Title: "Naboo N-1 Starfighter", Price: 250.45, Category: fighter}
	republicCruiser := starwars.Model{Id: 3, Title: "Republic Cruiser", Price: 450.00, Category: cruiser}
	attackCruiser := starwars.Model{Id: 4, Title: "Republic Attack Cruiser", Price: 950.00, Category: cruiser}
	eta2 := starwars.Model{Id: 5, Title: "ETA-2 Jedi Starfighter", Price: 650.50, Category: fighter}
	delta7 := starwars.Model{Id: 6, Title: "Delta-7 Jedi Starfighter", Price: 250.35, Category: fighter}
	bwing := starwars.Model{Id: 7, Title: "B-Wing", Price: 195.50, Category: fighter}
	ywing := starwars.Model{Id: 8, Title: "Y-Wing", Price: 45.50, Category: fighter}
	monCalamari := starwars.Model{Id: 9, Title: "Mon Calamari Star Crusier", Price: 1500.00, Category: cruiser}

	categories = append(categories, fighter, cruiser)
	models = append(models, vwing, n1, republicCruiser, attackCruiser, eta2, delta7, bwing, ywing, monCalamari)

	return categories, models
}
