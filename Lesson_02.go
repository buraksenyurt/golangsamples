/*
	Lesson_02
	Temel for kullanımı
	if ve switch case kullanımı
*/
package main

import (
	"fmt"
)

// Go dilinde sadece for döngüsü varmış biliyor muydunuz? :)
func main() {
	//1 ile 100 arasındaki sayılardan iki ile bölünebilenlerin toplamı
	toplam := 0
	for i := 0; i < 100; i++ {
		if i%2 == 0 { //Mod operatörü ile kalan hesaplanır
			toplam += i
		}
	}
	fmt.Println("2 ile bölünebilenlerin toplamı=", toplam)

	// While döngüsünün GO dilinde for ile yapılışı
	// 3 ile bölünebilen sayıların toplamı hesap ediliyor
	i := 0
	toplam = 0
	for i < 100 {
		i++
		if i%3 == 0 {
			toplam += i
		}
	}
	fmt.Println("3 ile bölünebilenlerin toplamı", toplam)

	//sonsuz döngü kurmak istersek bu şekilde
	//for {

	//}

	var t1, t2, t3 int

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			t2++
		} else if i%3 == 0 {
			t3++
		} else {
			t1++
		}
	}
	fmt.Printf("1 ile 100 arasında 2 ile bölünebilen %d sayı\n3 ile bölünebilen %d sayı var. Diğer kalan %d\n", t1, t2, t3)

	// switch case kullanımı
	// sinav_notu değerine göre ekrana bir bilgi yazılıyor
	sinav_notu := 46
	switch {
	case sinav_notu >= 0 && sinav_notu < 45:
		fmt.Println("üzgünüm ama sınıfta kaldın")
	case sinav_notu >= 45 && sinav_notu < 50:
		fmt.Println("Himmm bir kannat notu kullansam iyi olacak")
	case sinav_notu >= 50 && sinav_notu < 75:
		fmt.Println("Yeterli bir not görünüyor")
	case sinav_notu >= 75 && sinav_notu <= 100:
		fmt.Println("Güzelll")
	default: //Yukarıdaki koşullardan hiçbirisine uyulmadıysa
		fmt.Println("Sanki geçerli bir aralıkta değil gibisin")
	}
}
