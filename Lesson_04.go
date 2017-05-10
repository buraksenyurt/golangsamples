/*
	Lesson_04
	Fonksiyonlar
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fonksiyonları tanıyalım")
	toplam := Topla(4, 5)
	fmt.Println(toplam)
	a, b, c, d := DortIslem(3, 2)
	fmt.Printf("%f\t%f\t%f\t%f\n", a, b, c, d)
	fmt.Println(CokluToplam(1, 2, 3, 4, 5))
	fmt.Println(CokluToplam(125, 2, 3, 4, 7, -9, 0, 12, 4, 5, 90, 2, -6))
	sayiSlice := []int{3, 7, 1, 9, 10}
	fmt.Println(SliceToplam(sayiSlice))
	sozluk := map[string]string{
		"black": "kara",
		"white": "beyaz",
		"gold":  "altın",
	}
	map_yazdir(sozluk)
	fmt.Println(Faktoryel(10))

	//iç içe fonksiyon tanımlanabilir.
	//Yani bir fonksiyon içerisinde fonksiyon tanımlayıp kullanabiliriz
	var puan int
	fmt.Println("Aldığınız puanı girer misiniz?")
	fmt.Scanln(&puan)
	//geriye bool döndüren bir iç fonksiyon tanımladık
	puan_kontrol_fonksiyon := func(d int) bool {
		if d <= 50 {
			return false
		}
		return true
	}
	fmt.Println(puan_kontrol_fonksiyon(puan))
}

// En basit haliyle bir fonksiyon tanımı
func Topla(x, y int) int {
	return x + y
}

// Bir fonksiyondan birden fazla değer döndürebiliriz
// Hatta dönüş değişkenlerini adlandırabilir ve fonskiyon içerisinde
// Bu isimlerle kullanabiliriz
func DortIslem(x, y float32) (toplam, carpim, bolum, fark float32) {
	toplam = x + y
	carpim = x * y
	bolum = x / y
	fark = x - y
	return toplam, carpim, fark, bolum
}

// Bir fonksiyona n sayıda parametrede gönderebiliriz(Variadic fonksiyon)
func CokluToplam(sayilar ...int) int {
	toplam := 0 //toplam değişkenini dönüş parametre adı olarak da tanımlayabilirdik. DortIslem fonksiyonundaki gibi
	for _, sayi := range sayilar {
		toplam += sayi
	}
	return toplam
}

// Fonksiyon parametresi olarak Array, Slice, Map'de kullanabiliriz
func SliceToplam(sayilar []int) (toplam int) {
	for _, sayi := range sayilar {
		toplam += sayi
	}
	return toplam
}

func map_yazdir(icerik map[string]string) {
	for k, v := range icerik {
		fmt.Printf("%s\t%s\n", k, v)
	}
}

// Standart recursive(Özyinelemeli) fonksiyon örneği
func Faktoryel(sayi int) int {
	if sayi == 0 || sayi == 1 {
		return 1
	} else {
		return sayi * Faktoryel(sayi-1)
	}
}
