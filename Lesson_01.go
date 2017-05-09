/*
	// Lesson_01
	Merhaba Dünya
	Temel değişken tanımlamaları
*/
package main

// Uygulamada kullanacağımız paket bildirimleri
import (
	"fmt"
	"reflect"
)

//birden fazla sabit tanımı
const (
	izmir    = 35
	istanbul = 34
)

// Uygulamanın giriş noktası
func main() {
	fmt.Println("Merhaba. Benim adım Gopher!\nŞimdi temel tipleri görelim.")
	fmt.Println("Şu an\t", izmir, "\tplakalı bir ilden yazıyorum")
	var x, y int //önce değişken adı sonra tipi
	x = 4
	y = x + 8
	z := x + y //dinamik değişken ataması
	fmt.Println("Sayıların toplamı", z)

	var cap float32 = 3.89
	var alan = 3.14 * ((cap / 2) * (cap / 2))
	fmt.Println("Dairenin alanı", alan)

	var ad string = "jan\tclaud\n"
	var soyad string = "van\tdam"
	fmt.Println(ad + soyad)

	var isim, parola, email bool
	isim = true
	parola = false
	email = true
	fmt.Println(isim && parola && email) // Mantıksal ve kullanımı
	fmt.Println(isim || parola || email) // Mantıksal veya kullanımı
	fmt.Println("İstanbul" == "Ankara")  //Eşitlik kullanımı

	const pi float32 = 3.1415
	fmt.Println(pi)

	birDeger := false

	//bir değişkenin tipini anlamak için reflect kullanılabilir
	fmt.Println(reflect.TypeOf(pi), reflect.TypeOf(birDeger))
}
