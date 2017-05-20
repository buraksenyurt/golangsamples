/*
 Lesson_15

 Kanallar ve Select ifadesi ile birlikte GoRoutine'ler için timeout kontrolü nasıl yapılır?
Pratik olarak eş zamanlı iş parçacığı ile birlikte yürütülecek ve geri sayım yapacak bir başka iş parçacığı kullanılabilir
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.Second * 10 //bu süreyi someJob'un içindekiş 5 saniyenin altına düşürüp denediğimizde timeout oluştuğunu görebiliriz
	fmt.Println("Eş zamanlı iş için zaman aşımı süresi", timeout)
	cnl := make(chan string, 1)
	go someJob(cnl)
	select {
	case result := <-cnl:
		fmt.Println(result)
	case <-time.After(timeout): //burada timeout kontrolü yapıyoruz
		fmt.Println("İş istenen sürede tamamlanamadı. :/")
	}

	// Timeout kontrolü ayrı bir fonksiyona da verebiliriz
	timeout = 2
	fmt.Println("Yeni Zaman Aşımı(Timeout) süresi", timeout)
	jobCnl := make(chan string, 1)
	counterCnl := make(chan bool, 1)
	go someJob(jobCnl)
	go counter(counterCnl, timeout) // sayacı 2 saniyeye ayarladık
	select {
	case jobMsg := <-jobCnl:
		fmt.Println(jobMsg)
	case counterMsg := <-counterCnl:
		fmt.Println("Zaman aşımı oluşma durumu", counterMsg)
	}
}

func someJob(msg chan string) {
	fmt.Println("\tEş zamanlı işler yapılıyor.")
	time.Sleep(time.Second * 5) //belli bir süre uyutma
	msg <- "İş tamamlandı"
}

func counter(info chan bool, duration time.Duration) {
	fmt.Println("Sayaç", duration, "kadar bekleyecek")
	time.Sleep(time.Second * duration)
	info <- true
}
