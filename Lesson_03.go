/*
	// Lesson_03
	Array, Slice, Map kullanımları
	Array'ler de boyut sabittir
	Slice'lar dinamik boyutludur. Kapasiteleri belirlenebilir.
	Map tipleri key:value veri modeline uygun şekilde kullanılır
	range fonskiyonu ile bir dizi,kesit veya harita'yı for döngüsü ile kullanabiliriz
	_ ile istersek bir fonksiyondan dönen değeri kullanmayacağımızı ifade edebiliriz
*/
package main

import (
	"encoding/json" //Basit json serileştirme için kullanacağımız paket
	"fmt"
	"reflect"
)

func main() {
	// 82 elemanlı şehirler dizisi tanımladık
	var sehirler [81]string
	sehirler[34] = "istanbul"
	sehirler[6] = "Ankara"
	sehirler[16] = "Bursa"
	sehirler[35] = "İzmir"

	fmt.Println(len(sehirler)) //len ile dizinin eleman sayısını bulabiliriz.

	//indis bazlı olacak şekilde tüm şehirleri dolaşıyoruz.
	for i := 0; i < len(sehirler); i++ {
		if sehirler[i] != "" {
			fmt.Println(sehirler[i])
		}
	}

	//bu sefer 5 elemanlı float tipinden bir dizi oluşturduk
	adaylar := [5]float32{
		34,
		55,
		90,
		10,
		88,
	}
	var toplam float32
	var eleman_sayisi = len(adaylar)
	fmt.Println(reflect.TypeOf(eleman_sayisi)) //eleman_Sayisi dizisinin tipini yazdırdık.
	for i := 0; i < eleman_sayisi; i++ {
		toplam += adaylar[i]
	}
	ortalama := toplam / float32(eleman_sayisi)
	fmt.Printf("Ortalama %f\n", ortalama) //adayların not ortalamasını hesap ediyoruz

	// range kullanımına ait örnek
	// foreach benzeri bir döngü oluşuyor
	isimler := [4]string{"jan", "claud", "van", "dam"}
	for i, isim := range isimler {
		fmt.Printf("%d\t->\t%s\n", i, isim)
	}

	var basarililar int = 0
	puanlar := [5]int{34, 55, 23, 90, 98}
	for _, puan := range puanlar {
		if puan >= 50 {
			basarililar++
		}
	}
	fmt.Printf("Başarılı %d öğrenci var\n", basarililar)

	// iki boyutlu bir dizi örneği. 3X2lik bir matris oluşturduk.
	matris := [3][2]int{{2, 3}, {6, 1}, {-9, 8}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d\t", matris[i][j])
		}
		fmt.Println()
	}

	//Örnek bir slice tanımı. string elemanlardan oluşuyor
	oyuncu_adlari := []string{"mayk", "miki", "lora", "clara", "zorro", "dam", "edriyın", "raki", "barbarossa"}
	// indis değerlerini ele almadık. _ operatörü ile.
	// sadece value'ları işliyoruz. Yani string içerikleri
	for _, value := range oyuncu_adlari {
		fmt.Println(value)
	}
	fmt.Println()
	//alt_kume := oyuncu_adlari[3:6]
	alt_kume := oyuncu_adlari[6:]
	//: işareti ile bir diziden veya kesittten alt kesitler alabiliriz.
	// Baştan itibaren, sondan itibaren veya iki rakam arasındaki kısımlardan
	alt_kume = append(alt_kume, "hera") //append ile slice sonuna eleman eklenebilir
	alt_kume = append(alt_kume, "sizar")
	for _, value := range alt_kume {
		fmt.Println(value)
	}
	var iller []string
	iller = make([]string, 5, 10) //5 eleman içeren 10a kadar genişleyebilen kesit. Eleman sayısı ve başlangıç kapasitesini belirttik
	iller[0] = "istanbul"
	iller[1] = "izmir"
	iller[2] = "ankara"
	iller[3] = "bursa"
	iller[4] = "gaziantep"
	iller = append(iller, "trabzon")
	for i := 0; i < len(iller); i++ {
		fmt.Printf("%d:%s\n", i, iller[i])
	}
	for i, il := range iller {
		fmt.Printf("%d:%s\n", i, il)
	}

	// Basit bir map tanımlaması
	// key ve value içerikleri string tipten olacaklar
	sozluk := make(map[string]string)
	sozluk["white"] = "beyaz"
	sozluk["black"] = "siyah"
	sozluk["red"] = "kirmizi"
	sozluk["blue"] = "mavi"
	for key, value := range sozluk { //hem key hem de value değerlerini alıyoruz.
		fmt.Printf("[%s:%s]\n", key, value)
	}
	// sozluk map içeriğini json formatına dönüştürüyoruz ve ekrana basıyoruz
	jsonContent, _ := json.MarshalIndent(sozluk, "", "   ")
	fmt.Println(string(jsonContent))

	//Bu sefer key içerikleri string value içerikleri int olan
	//bir map değişkeni tanımlandı
	envanvter := map[string]int{
		"Laptop":       34,
		"Desktop":      5,
		"Tablet":       12,
		"Cep Telefonu": 34,
	}
	// Envanterdeki toplam cihaz sayısını buluyoruz.
	var toplam_cihaz int = 0
	for _, value := range envanvter {
		toplam_cihaz += value
	}
	fmt.Printf("Envanterde %d cihaz var\n", toplam_cihaz)
}
