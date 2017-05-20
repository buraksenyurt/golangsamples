/* Lesson_14
Çok sayıda kanal kullanımı ve select ile senkronizasyonun sağlanması
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	m1 := make(chan string) //kanallar tanımlanıyor
	m2 := make(chan string)
	m3 := make(chan string)

	go jobA(m1, "A..C")
	go jobB(m2, "C..M")
	go jobC(m3, "M..Z")

	for i := 0; i < 3; i++ {
		select { // işlemlerden hangisinin bittiğini anlamak için select ifadesi kullandık
		case messageA := <-m1: //m1 kanalından işin bittiğine dair bilgi geldiyse
			fmt.Println(messageA)
		case messageB := <-m2:
			fmt.Println(messageB) //m2 kanalından işin bittiğine dair bilgi geldiyse
		case messageC := <-m3:
			fmt.Println(messageC) //m3 kanalından işin bittiğine dair bilgi geldiyse
		}
	}
	fmt.Println(time.Now())
}

func jobA(msg chan string, sets string) {
	fmt.Println(sets, "için işlemler yapılacak")
	time.Sleep(time.Second * 10) // belli bir süre uyutalım
	msg <- "A görevi tamamlandı"
}
func jobB(msg chan string, sets string) {
	fmt.Println(sets, "için işlemler yapılacak")
	time.Sleep(time.Second * 3)
	msg <- "B görevi tamamlandı"
}
func jobC(msg chan string, sets string) {
	fmt.Println(sets, "için işlemler yapılacak")
	time.Sleep(time.Second * 6)
	msg <- "C görevi tamamlandı"
}
