/*
	Kod bir web sayfasındaki HTML elementlerinde arama yapmak için kullanılıyor.
	Örnekte www.buraksenyurt.com adresindeki post-title css sınıfını kullanan
	satırlardaki a elementi yakalanıp text içeriği döndürülüyor.
	Bir başka deyişle makale başlıklarını yakalıyoruz.
*/
package main

import (
	"fmt"
	"log"

	q "github.com/PuerkitoBio/goquery" //q ile bir alias tanımlamış olduk
)

func main() {
	// Web dokümanını aç
	doc, err := q.NewDocument("http://www.buraksenyurt.com")
	if err != nil {
		log.Fatal(err)
	}

	// post-title css'inin uygulandığı her bir element için iç fonksiyonu çalıştır
	doc.Find(".post-title").Each(func(
		i int, s *q.Selection) {
		title := s.Find("a").Text() //Selection üzerinden a elementini ara
		fmt.Printf("%d\t%s\n", i, title)
	})
}
