package main

import (
	"database/sql"
	"encoding/json"   //JSON Serileştirmeler için
	"entity/starwars" //Entity paketimiz
	"fmt"
	"log" //gerekli yerlerde log atmak için
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter" // Yönlendirme işlemleri için kullandığımız paket
	_ "github.com/mattn/go-sqlite3"       // SQLite veritabanı işlemlerimizi kolaylaştıran paket
)

func main() {
	router := httprouter.New()

	// yönlendiriciyi oluşturduk
	router.GET("/", home)                                        //root
	router.GET("/categories", getCategories)                     //Tüm kategoriler
	router.GET("/categories/:categoryId", getModelsByCategoryId) //belli bir kateogori id'ye bağlı modeller
	router.GET("/models/:firstLetter", getModelsByFirstLetter)   // Baş harfine göre modeller
	router.POST("/newCategory", createCategory)                  // Yeni bir kategori eklemek için kullanılan fonksiyon

	http.ListenAndServe(":4571", router) //localhost:4571 adresinden host ediyoruz
}

func createCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	category := starwars.Category{}
	json.NewDecoder(r.Body).Decode(&category) //Body üstünden gelen JSON içeriğini ayrıştır ve category nesnesi al
	log.Printf("Insert request. %d,%s\n", category.Id, category.Name)
	conn, _ := sql.Open("sqlite3", "starwars.sdb")                                       // DB Bağlantısını aç
	defer conn.Close()                                                                   // fonksiyondan çıkarken bağlantıyı kapat
	_, err := conn.Exec("Insert into Category values (?,?)", category.Id, category.Name) //Insert sorgusunu çalıştır. ? ile parametreler veriliyor
	if err == nil {
		render(w, category) // render fonksiyonu ile category içeriğini bastır
	} else {
		log.Println(err.Error())
		//404 basılabilir
	}
}

func getCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conn, _ := sql.Open("sqlite3", "starwars.sdb") //bağlantıyı aç
	defer conn.Close()
	rows, _ := conn.Query("Select Id,Name from Category order by Name") //select sorgu sonuçlarını rows değişkenine al
	defer rows.Close()                                                  //çıkarken bağlantıyı kapat
	categories := make([]*starwars.Category, 0)
	for rows.Next() { //tüm sonuç kümesinde dön
		category := new(starwars.Category)
		rows.Scan(&category.Id, &category.Name)   //Sorgu sonuçlarındaki alanları category nesnesindeki alanlar ile eşleştir
		categories = append(categories, category) //slice'a ekle
	}
	render(w, categories) //slice içeriğini JSON formatından istemciye basacak fonksiyonu çağır
}

func getModelsByCategoryId(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	conn, _ := sql.Open("sqlite3", "starwars.sdb")
	defer conn.Close()
	id, _ := strconv.Atoi(params.ByName("categoryId"))
	cRow := conn.QueryRow("Select * from Category Where Id=?", id)
	ctgry := new(starwars.Category)
	if cRow != nil {
		cRow.Scan(&ctgry.Id, &ctgry.Name)
		rows, _ := conn.Query("Select Id,Title,ListPrice from Model where CategoryId=?", id)
		defer rows.Close()
		models := make([]*starwars.Model, 0)
		for rows.Next() {
			model := new(starwars.Model)
			model.Category = starwars.Category{Id: ctgry.Id, Name: ctgry.Name}
			rows.Scan(&model.Id, &model.Title, &model.Price)
			models = append(models, model)
		}
		render(w, models)
	}
}

func getModelsByFirstLetter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	conn, _ := sql.Open("sqlite3", "starwars.sdb")
	defer conn.Close()
	// Like kullanımındaki 'A%' için önce Sprintf ile sorgu metni oluşturuldu
	statement := fmt.Sprintf("Select Id,Title,ListPrice from Model where Title like '%s%%'", params.ByName("firstLetter"))
	rows, _ := conn.Query(statement) // Sorguyu çalıştır ve sonuçları rows setine aktar
	defer rows.Close()
	models := make([]*starwars.Model, 0)
	for rows.Next() {
		model := new(starwars.Model)
		rows.Scan(&model.Id, &model.Title, &model.Price)
		models = append(models, model)
	}
	render(w, models)
}

func home(rWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprintf(rWriter, "Star Wars universe!")
}

// İkinci parametre göre bu fonksiyona categories, models, category gibi önceki fonksiyonlarda kullandığımız tüm nesneleri aktarabiliriz
func render(w http.ResponseWriter, d data) {
	// Header içeriğini JSON çıktı verecek şekilde işaretle
	w.Header().Set("Content-Type", "application/json")
	// İstmeciye HTTP 200 OK mesajını ver
	w.WriteHeader(200)
	// gelen d içeriğini JSON olarak serileştir
	jContent, _ := json.Marshal(d)
	// Serileşen içeriği ResponseWriter üzerinden istemciye bas
	fmt.Fprintf(w, "%s", jContent)
}

type data interface {
}
