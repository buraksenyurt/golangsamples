/*
 Lesson_12
 Channels
 kanallar yardımıyla GoRoutine'ler arası veya GoRoutine'ler ile çağıran yerler arası
 bilgi transferi yapma şansına sahip oluruz
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 1nci örnek en ilkel haliyle kanal kullanımı

	transitInfo := make(chan string) //içerisinden string veri taşınabilecek bir kanal tanımı yapıldı
	go func() {                      //iç fonksiyon tanımladık ve bunu go ile bir GoRoutine haline getirdik
		fmt.Println("Burası mapping işlemlerini yapar")
		time.Sleep(time.Millisecond * 1000)
		transitInfo <- "mapping sonucu uygun 3 donör bulundu" // kanala string bir veri içeriği bırakıyoruz
	}()
	incomingInfo := <-transitInfo // GoRoutine içerisinden kanala yazılan bilgi <- operatörü ile değişkene atanıyor
	fmt.Printf("Map isimli goroutine'den '%s' bilgisi döndü\n", incomingInfo)
	close(transitInfo) // kanalı kapatıyoruz

	// 2nci örnek
	// oluşturduğumuz kanalı Listener ve Listener içerisindeki iç fonksiyon da ortaklaşa kullanmaktayız
	message := make(chan string)
	go Listener(message)
	input := <-message
	fmt.Printf("Listener dedi ki [%s]\n", input)

	close(message)

	// 3ncü örnek
	// Dilersek kanallardaki içerik sayısını kısıtlayabilir ve tampon bazlı kullanılmalarını sağlayabiliriz
	// Sadece 5 eleman alabilecek bir Channel için
	parts := make(chan string, 5)
	parts <- "ayakkabılar"
	parts <- "ceketler"
	parts <- "pantalonlar"
	parts <- "bluzlar"
	parts <- "çoraplar"

	// döngü içerisinde 5 GoRoutine çağırılıyor ve her birisinde kanaldaki bilgi yakalanıyor
	for i := 0; i < 5; i++ {
		go func(p chan string) {
			value := <-p
			fmt.Printf("\t[%s] parçalar işlenecek\n", value)
		}(parts)
	}
	close(parts)

	// 4ncü örnek
	// kanallarda yön verilebilir ve tek yönlü çalışacak hale getirilebilir
	// Yani sadece mesaj alma veya mesaj gönderme şeklinde kanallar tanımlanabilir
	soundChannel := make(chan string, 1) //sadece 1 eleman taşıyacak bir kanal tanımlandı
	microphone(soundChannel, "HOLA!")
	soundBox(soundChannel, 10, 1500)

	// 5nci örnek
	// GoRoutine'lerin senkronizasyonu için kanallardan yararlanılabilir
	statusChannel := make(chan bool)
	go worker(statusChannel)
	<-statusChannel //Burada kanaldan bilgi alınacaya kadar ana iş parçacığını bloklamış olduk.

	var enter string
	fmt.Printf("\nÇıkmak için Enter\n")
	fmt.Scanln(&enter)
}

func worker(completed chan bool) {
	fmt.Println("Bir takım işlemler yapılıyor")
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("İşlemler tamamlandı")
	completed <- true //burada kanala işlemin bittiğine dair bir mesaj bırakıyoruz
}

func soundBox(sound <-chan string, volumeLevel int, duration int) { //<-chan nedeniyle kanaldan sadece veri alabilir(kanaldan veri okuyabilir ama yazamaz)
	fmt.Printf("Sound is %s.\nLevel = %d\nDuration = %d\n", <-sound, volumeLevel, duration)
}
func microphone(sound chan<- string, message string) { // chan<- nedeniyle sadece veri gönderebilir(kanala veri yazabilir ama okuyamaz)
	sound <- message
}

func Listener(msg chan string) {
	msg <- "pong"
	go func(chn chan string) {
		chn <- "ping"
	}(msg)
	output := <-msg
	fmt.Printf("\tiç fonksiyon dedi ki [%s]\n", output)
}
