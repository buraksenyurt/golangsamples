/*
Lesson_16
Timer kullanımı
belli bir süre sonra başlatılacak işler veya belirli zaman peryitolarında başlatılacak işler için kullanabiliriz
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 4)

	fmt.Println("timer nesnesi tanımlandı. Kod akışına devam ediyor.")
	fmt.Println(time.Now())
	now := <-timer.C //C ile timer'ın NewTimer'a parametre olarak gelen süre sonrasındaki zaman elde edilir
	fmt.Println("Timer ile belirtilen süre doldu.")
	fmt.Println(now)

	// Bu seferki Timer, süresi dolduğu için Expire durumuna düşecek
	// Bir Timer'ı expire olmadan önce durdurmak istediğimiz senaryolarda ele alabiliriz
	timer = time.NewTimer(time.Second)
	go func() {
		<-timer.C
		fmt.Println("İkinci timer süresi geçti") // time.Second nedeniyle func içerisinde timer.C yakalanamadan Stop metoduna düşülür
	}()
	stop := timer.Stop()
	if stop {
		fmt.Println("Timer durduruldu")
	}

	// ticker ile zamanlanmış görevler hazırlayabiliriz.
	tickTime := time.NewTicker(time.Second * 2) // iki saniyede bir zaman döndürecek Ticker tanımlandı
	go func() {
		fmt.Println("İş yapıyorum...")
		for t := range tickTime.C { // C ile yukarıdaki tickTime'ın o anki süresi ele yakalandı
			fmt.Println(t)
		}
	}()
	//time.Sleep(time.Second * 12) // main thread 12 saniye duracak. Bu süre boyunca 2 saniyede bir for t:=range bloğu çalışacaktır
	//tickTime.Stop()              //Ticker durduruldu

	// Yukarıdaki kullanımdan farklı olarak şimdi kullanıcı Enter tuşuna basana kadar for t:=range bloğu çalışacaktır
	var enter string
	fmt.Println("Çıkmak için Enter tuşuna basınız")
	fmt.Scanln(&enter)
	tickTime.Stop()
}
