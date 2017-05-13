/*
	Lesson_07
	Basit pointer kullanımı
*/
package main

import (
	"fmt"
)

func main() {
	var point int = 41
	fmt.Println(&point)
	// & ile değişken adresini yakalayabiliriz
	fmt.Printf("%d sayısının bellek adresi %x dir\n", point, &point)
	pntr := &point                                // point değişkeninin bellek adresini aldık
	fmt.Println(pntr, " adresindeki sayı", *pntr) //* operatörü ile pointer'ın işaret ettiği alandaki değer yakalanabilir

	newLocation := pntr //Pointer'ı başka bir pointer değişkenine atayabiliriz.
	*newLocation = 99   //newLocation ve pntr aynı bellek adresini işaret ediyorlar. Bu yüzden birisindeki değişim diğerini de etkileyecektir
	fmt.Println("Bellek adresleri pntr=>", pntr, " newLocation=>", newLocation)
	fmt.Println("Atama sonrası değerler\npntr=", *pntr, " newLocation=", *newLocation)

	lucky_number := 7
	calculate(&lucky_number)
	fmt.Println("Lucky Number = ", lucky_number)

	stk := Stock{125, 25} // Stock bir struct'tır. Value Type olarak çalışır. Ama pointer yardımıyla fonksiyonlara referans tipi olarak geçirilebilir
	increase_Stock_By_Fifty(&stk)
	fmt.Printf("Stock min %f max %f\n", stk.low, stk.high)
}

func calculate(number *int) {
	//fonksiyonlara parametre olarak pointer verebiliriz.
	//Böylece fonksiyon içerisinde değişken kopyası oluşturmak yerine
	//referans adresi taşıdığımızdan daha optimize ve bellek dostu kod üretmiş olabiliriz
	fmt.Println("Fonksiyona gelen adres ", number, "\nDeğişken değeri ", *number)
	*number += 100 //şimdi gelen pointer üzerinden değer değişikliği yaptık. Buna göre calculate'i çağırdığımız main fonksiyonundaki lucky_number içeriği de değişmiş olur

	some_points := []float32{3.4, 2.1, 1.98, -4}
	do_something(some_points) //slice'lar referans tipidir. Bu yüzden fonksiyonlara parametre olarak geçtiklerinde ilgili fonksiyonda yapılacak değişiklikler orjinal konumlarındakini de etkiler
	for _, p := range some_points {
		fmt.Println(p)
	}
}

//points bir slice ve referans türü
func do_something(points []float32) {
	for i, point := range points {
		points[i] = point + 1 //buradaki değişiklik some_points üzerinde de etkili olacaktır
	}
}

type Stock struct {
	high float64
	low  float64
}

func increase_Stock_By_Fifty(stock *Stock) {
	stock.high += 50 //calculate fonksiyonunda parametre değerini değiştirirken kullandığımız * operatörü burada kullanılmamıştır dikkat edelim
	stock.low += 50
}
