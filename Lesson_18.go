/*
Lesson_18
strings paketinin kullanımı
Temel metinsel işlemler
Burada geçenler dışındaki String işlemleri için
https://golang.org/pkg/strings/
adresine bakılabilir
*/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	paragraph := `Line 1 : If there was a problem walking to the file or directory named by path, 
	Line 2 : the incoming error will describe the problem and the function can decide 
	Line 3 : how to handle that error. If an error is returned, processing stops. 
	Line 4 : The sole exception is when the function returns the special value SkipDir.
	Line 5:end of code !` // " yerine ` kullanarak birden fazla satırdan oluşan metinler belirtebiliriz

	containsSample("an")
	containsSample("ist")
	countSample(paragraph)
	fieldsSample(paragraph, 7)
	fieldsFuncSample(paragraph)

	sampleText := "bir bilmecem var çocuklar! Haydi sor sor, çay'da kahvaltı da yenir :) Acaba nedir nedir?"
	hasPrefixSample(sampleText, "bir")
	hasPrefixSample(sampleText, "peki")

	//IndexAny ile bir metinde ilgili veri setindeki karakterlerden ilkinin görüldüğü indis değeri bulunur
	indexNo := strings.IndexAny(sampleText, ":,?!;.")
	fmt.Println(indexNo) //26 döndürür. Çünkü ikinci parametrede aranan içerikteki bilgilerden ilk görülen yerin indisi döner

	// Kullanışlı fonksiyonlardan biris de Join
	// Örneğin
	values := []string{"data source", "tcp", "connection", "log", "function"}
	newValue := strings.Join(values, "|") // values isimli string dizideki elemanların arasına | işareti koyarak birleştirir
	fmt.Println(newValue)

	// Join benzeri eğlenceli fonksiyonlardan birisi de Repeat
	// Bir ifadenin n sayıda tekrarı için kullanılıyor.
	fmt.Printf("Hey Ney%s\n", strings.Repeat("Na", 3))

	//Map kullanımı. Bu fonksiyon ile bir metnin tüm karakterlerinin belli bir kurala göre değiştirilmesi sağlanabilir.
	//Şifreleme işlemlerinde mesela. Popüler örnek ROT13 şifreleme algoritması
	rot13 := func(c rune) rune { //rune tipinin anlamlandığı bir yer. int32 gibi olan rune aslında karakterin sayısal karşılığını veriyor.
		return c + 13
	}
	encryptedText := strings.Map(rot13, "ordu sağ kanattan sabah şafakla harekete geçecek")
	fmt.Println(encryptedText)
	decryptedText := strings.Map(func(c rune) rune {
		return c - 13
	}, encryptedText)
	fmt.Println(decryptedText)

	//Split fonksiyonu ile bir metni belli bir karaktere göre ayırmamız mümkündür
	product := "Pro Go Lang|Book|35,50|550|ISBN:345676"
	columns := strings.Split(product, "|")
	for _, column := range columns {
		fmt.Println(column)
	}

	// Title, ToLower, ToUpper ile harf çevirimleri
	motto := "bEniM hala UMUDUM vaAArr"
	fmt.Println(motto)
	fmt.Println(strings.Title(motto))   //sadece başharfleri büyüğe çevirdi
	fmt.Println(strings.ToUpper(motto)) // tüm harfleri büyüye çevirdi
	fmt.Println(strings.ToLower(motto)) // tüm harfleri küçüğe çevirdi
}

// Bir metnin başında belirtilen ifadenin olup olmadığını söyler
// HasSuffix kullanımı da benzerdir. Sonunda aranan ifadenin olup olmadığını söyler
func hasPrefixSample(text string, searching string) {
	if strings.HasPrefix(text, searching) {
		fmt.Printf("\n'%s' in başında '%s' VAR!\n", text, searching)
	} else {
		fmt.Printf("\n'%s' in başında '%s' YOK!\n", text, searching)
	}
}

/*FieldsFunc örneği. İkinci parametreye dikkat. Normalde orada bir fonksiyon tanımı var.
Parametre olarak geçen fonksiyon karakterin sayı olmama koşuluna bakara bir sonuç döndürüyor.
Böylece metin içerisinde geçen sayısal değerleri elde ediyoruz.*/
func fieldsFuncSample(paragraph string) {
	words := strings.FieldsFunc(paragraph, func(c rune) bool { //rune int32'nin alias'ıdır. Karakter değerlerini tam sayı değerlerinden ayırt ederken kullanırız.
		return !unicode.IsNumber(c)
	})
	fmt.Printf("%q", words)
}

// fields fonksiyon örneği. Metinde geçen kelimelerden belli bir değerden uzun olanlarını listelemek
func fieldsSample(paragraph string, charCount int) {
	fmt.Println("Harf sayısı -", charCount, "- dan fazla olanlar")
	words := strings.Fields(paragraph)
	for _, word := range words {
		if len(word) >= charCount {
			fmt.Println(word)
		}
	}
}

// bir paragraftaki kelime sayısını bulma örneğinde Count fonksiyonu kullanımı
func countSample(paragraph string) {
	space_count := strings.Count(paragraph, " ")
	line_count := strings.Count(paragraph, "\n")
	fmt.Printf("Metinde %d boşluk ve %d satır bulundu.\nBuna göre toplam kelime sayısı %d\n", space_count, line_count, space_count+line_count-1)
}

// contains ile metin için metin arama örnek kullanımı
func containsSample(part string) {
	words := []string{
		"istanbul",
		"mavi",
		"yeşil",
		"antalya",
		"pist",
		"ankara",
		"izmir",
		"antartika",
		"arktika",
		"anason",
		"mandalina",
		"banana",
	}
	fmt.Println("'", part, "' geçen kelimeler")
	for _, word := range words {
		if strings.Contains(word, part) {
			fmt.Println(word)
		}
	}
}
