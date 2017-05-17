/*
Lesson_11
En basit haliyle GoRoutine kullanımı
GoRoutine'ler eş zamanlı çalışan fonksiyonlardır
Bir fonksiyonun eş zamanlı çalışması için go anahtar kelimesi ile çağırılması yeterlidir
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	go loop1(0, 25) //Bu 3 fonksiyon GoRoutine olarak eş zamanlı çalışırlar
	go loop2()
	go loop1(-25, -5)

	// Yukarıdaki işlemler bitmeden ana uygulama sonlanabilir. Bunu şu an için önlemek
	// ve yukarıdaki işlemlerin bittiğini görmek için ekrandan tuşa basılması beklenir
	var enter string
	fmt.Println("\nÇıkmak için bir tuşa basınız")
	fmt.Scanln(&enter)
}

func loop1(min, max int) {
	for i := min; i < max; i++ {
		fmt.Printf("%d", i)
		time.Sleep(time.Microsecond * 500)
	}
}
func loop2() {
	letters := "abcçdefgğhijklmnoöprsştuüvyz"
	for _, c := range letters {
		fmt.Printf("%s", string(c))
		time.Sleep(time.Millisecond * 550)
	}
}
