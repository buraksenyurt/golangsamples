/*
	Lesson_21
	net/http ve io/util paketine ait bir kaç örnek
*/
package main

import (
	"fmt"
	"io/ioutil" //io stream'lerini rahat okumak için kullandık
	"net/http"  //Http tabanlı istemci sunucu kabiliyetlerini sunar. Web sunucusu yapmak, HTTP Get,Post,Put,Delete gibi talepler yapmak ve daha pek çok operasyon sunar.
	"os"        //web içeriklerini yazacağımız dosyanın oluşumu için kullandık
)

func main() {
	// getWebPageContent("https://golang.org/pkg/net/http/", "net.http.pkg.html")

	// ListenAndServer fonksiyonuna verdiğimiz ilk parametre ile 4567 nolu porttan
	// Lesson_21.go'nun olduğu klasördeki docs dizinini kullanıma açtık
	// kodu çalıştırdıktan sonra makinenizden http://localhost:4567 ile içeriği görmeyi deneyin
	// (Tabii docs isimli klasörü açıp içerisine bir kaç dosya koymayı unutmayalım)
	err := http.ListenAndServe(":4567", http.FileServer(http.Dir("docs")))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getWebPageContent(page string, fileName string) {
	response, err := http.Get(page) //parametre olarak gelen sayfaya Get talebi yolluyoruz
	if err == nil {
		defer response.Body.Close() //defer ettiğimiz şey Body'nin kapatılması
		// Header'dan Status, StatusCode, Content Length gibi bilgileri alıyoruz
		fmt.Printf("Response Status\t:\t%s\nResponse Code\t:\t%d\nContent Length\t:\t%d\n", response.Status, response.StatusCode, response.ContentLength)

		//Normalde Header bir map olarak geriye dönmektedir
		fmt.Println(response.Header)
		// map üzerinden dolaşarak da Header bilgilerini yakalayabiliriz.
		for k, v := range response.Header {
			fmt.Println(k)
			for _, vv := range v {
				fmt.Println("\t", vv)
			}
		}

		if response.StatusCode == 200 { // Durum kodumuz HTTP 200 OK ise
			content, _ := ioutil.ReadAll(response.Body) //Gövdeyi okuyoruz. ReadAll metodu içeriğin byte array olarak elde edilmesini sağlıyor
			f, _ := os.Create(fileName)                 //içeriği kaydedeceğimiz dosyayı oluşturuyoruz
			f.WriteString(string(content))              //içeriğini yazıyoruz
			f.Close()                                   // dosyayı kapatıyoruz
			fmt.Println("İçerik kaydedildi")
		}

	} else {
		fmt.Printf("Bir hata oluştu dostum yaa :|\n\t%s\n", err.Error())
	}
}
