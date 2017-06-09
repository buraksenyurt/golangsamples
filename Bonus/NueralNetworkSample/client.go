/*
	Basit bir yapay sinir ağı uygulaması
	goml/gobrain paketinden yararlanıştır
	and işlemi öğretilmiştir
*/
package main

import (
	"fmt"
	"math/rand"

	"github.com/goml/gobrain" //LiteIDE'de Build->Get ile paketi yüklemeyi unutmayalım
)

func main() {
	fmt.Println("Yapay sinir ağına VE işlemini öğretiyoruz")
	rand.Seed(0)

	//sinir ağına öğreteceğimiz VE işlemine ait örnekleme kümesi
	patterns := [][][]float64{
		{{0, 0}, {1}},
		{{0, 1}, {0}},
		{{1, 0}, {0}},
		{{1, 1}, {1}},
	}

	// FeedForward adı verilen fonksiyonellik oluşturuluyor
	//neural network function
	nnf := &gobrain.FeedForward{}
	nnf.Init(2, 2, 1)                         //2 giriş boğumu,2 gizli katman boğumu ve 1 sonuç boğumu olacak
	nnf.Train(patterns, 6000, 0.6, 0.4, true) //6000 devirlik bir öğrenme olacak. Öğrenme oranı 0.6, momentum katsayısı 0.4 ve öğrenme sırasındaki hataları da alalım mı sorusunun cevabı son parametre
	nnf.Test(patterns)                        //Testi başlat

	// Çıktıda -> önceki kısım test verisi, -> dan sonraki kısım o test için elde edilen sonuç ve :dan sonraki kısım beklenen sonuçtur.

	inputs := []float64{4, 2} //Girdiğimiz değerler için bir tahminleme yaptırıp sonucunu bastırıyoruz
	result := nnf.Update(inputs)
	fmt.Println(result)
}
