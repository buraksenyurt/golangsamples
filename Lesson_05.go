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
