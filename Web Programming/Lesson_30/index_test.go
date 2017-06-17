/*
	index.go'yu test edecek bir unit test içeriği
	index.go aslında bir web hizmeti sunuyor.
	oradaki Greetings fonksiyonunu test ediyoruz.
*/
package main

import (
	"net/http"
	"net/http/httptest"
	"testing" //test fonksiyonlarını içeren paket
)

func Test_Greetings(t *testing.T) {
	newRequest, err := http.NewRequest("GET", "/", nil) //HTTP Get'ten test talebi üretiyoruz.
	if err != nil {
		t.Fatal(err) //Eğer talep sırasında hata oluştuysa Fatal test sonucu bırakıyoruz.
	}
	response := httptest.NewRecorder() //http talebi sonucunu tutacak bir kaydedici nesne
	Greetings(response, newRequest)    //Test edilecek fonksiyon request ve response nesneleri ile çağırlıyor.
	// exptected := "<h1>Wellcome Back</h1>" //Beklediğimiz içerik
	exptected := "Wellcome Back</h1>"  //Beklediğimiz içerik(bilerek hatalı ürettik)
	incoming := response.Body.String() // Gelen cevaptaki body içeriği
	if exptected != incoming {
		t.Fatalf("/ için yapılan talebe karşılık beklenen cevap gelmedi.\nBeklenen-> %s\nGelen-> %s", exptected, incoming)
	}
}
