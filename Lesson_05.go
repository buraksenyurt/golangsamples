/*
 Lesson_05
 defer kullanımı
 bir fonksiyon tamamlanmadan hemen önce çalıştırılmasını istediğimiz fonksiyonlar için kullanırız
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("main başlangıcı")
	sonuc := dosya_sifrele("bilgiler.dat")
	fmt.Println(sonuc)
	fmt.Println("main bitişi")
	fmt.Println(dosya_ayristir("urunler.json"))	
}

func dosya_sifrele(dosya string) bool {
	defer dosyayi_kapat(dosya) // sonra bu çalışır
	defer bellegi_temizle()    // fonksiyondan çıkılırken önce bu
	fmt.Println("Şifreleme operasyonu yapılıyor")
	return true
}

func dosyayi_kapat(dosya string) {
	fmt.Println("kalan veriler dosyaya yazilip kapatiliyor")
}
func bellegi_temizle() {
	fmt.Println("Ön bellek temizleniyor")
}
func dosya_ayristir(dosya string)string{
	//defer operasyonları iç fonksiyon olarak da tanımlanabilir
	defer func(){
		fmt.Printf("%s için gerekli kapatma operasyonları yapılacaktır\n",dosya)
	}()
	fmt.Println("Dosya açılıyor...")
	fmt.Println("Ayrıştırma işlemi yapılıyor")
	return "operasyon başarılı"
}

//defer genellikle panic oluşma hallerinde de ele alınır.
//panic durumlarında defer ile işaret edilen fonksiyon veya fonksiyon bloğu
//otomatik olarak çalışır. try...catch...finally'deki finally gibi işlem görür
// diyebiliriz.
