/*
 Lesson_06
 panic recover defer kullanımı
*/
package main

import (
	"fmt"
)

func main() {
	// main içerisinden çıkılırkan bu iç fonksiyon devreye girecektir
	//eğer bir hata varsa içeride yakalanıp ekrana basılır
	defer func() {
		err := recover() //bir hata varsa yakalayalım
		if err != nil {
			fmt.Printf("Ana program fonksiyonunda hata oluştu\n\tError:%s\n", err)
		}
	}()

	fmt.Println("Uygulama başlıyor")
	sonuc := baglan("net.tcp://localhost:9023/services/data")
	fmt.Println(sonuc) //panic nedeniyle return sonuc çalışmayacağından sonuc boş dönecektir

	sayilar := make([]int, 5) //5 elamanlı tanımladığımız slice
	sayilar[6] = 100          //olmayan bir indise eleman atamaya çalışıyoruz. Built-in runtime hatası oluşacaktır
	sayilar[1] = 90
	fmt.Println("Ana program sonu") //main'deki hata nedeniyle burası hiç çalışmayacak
	//slice'a hata olmayacak şekilde değer atayıp tekrar deneyin
}

func baglan(conStr string) string {
	defer func() { // defer ile hata oluştuğunda devreye girecek son operasyonu işaret edebiliriz
		err := recover() //hatayı yakalayalım
		if err != nil {  //eğer bir hata varsa
			fmt.Printf("Servis bağlantı hatası\n\tError:%s\n", err) //err mesajı panic ile ürettiğimiz mesaj olacaktır
			//loglama yapılabilir
		}
	}()

	//normal kod satırları
	fmt.Println("Bağlantı açılıyor")
	panic("Bağlantı yapılırken hata oluştu") //sembolik olarak bir hata fırlatıldı
	return "durum bilgisi"
}
