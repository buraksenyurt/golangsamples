/*
	Lesson_10
	interface kullanımı
	Bir interface yardımıyla metod bazlı sözleşmeler(contract) sunabiliriz
*/
package main

import (
	"fmt"
)

func main() {
	vosvos := Car{id: 10, name: "Orange Vos Vos"}
	cesna := Plane{id: 900, owner: "Chesna", maxAltitude: 3000}
	pejo205 := Car{id: 12, name: "Pejo 205 GTI"}

	startVehicleEngine(vosvos)
	startVehicleEngine(cesna)
	startVehicleEngine(pejo205)

	// allTeam isimli slice içerisinde Car ve Plane tipli yapı örnekleri bulunuyor
	allTeam := []VehicleContract{
		vosvos,
		Car{id: 2, name: "Little Dori FSI"},
		cesna,
		pejo205,
		Plane{id: 16, owner: "Loched Martin", maxAltitude: 35000},
	}
	// tüm üyelerinde dolaşıp o anki yapı örneği kimse onun için startEngine metodunun çalıştırılması sağlanabilir
	for _, member := range allTeam {
		startVehicleEngine(member)
	}
}

// bu interface sözleşmesi startEngine isimli bir metod sunuyor
// yapılarda da bu metod tanımlanırsa
// startVehicleEngine metoduna parametre olarak ilgili yapılar da gönderilebilir
// nitekim ilgili yapılar sözleşmenin belirttiği metodu uygulamışlardır
type VehicleContract interface {
	startEngine()
}

// Metod parametre olarak VehicleContract arayüzünü alır
// Car ve Plane yapıları arayüzde belirtilen startEngine metodunu uyguladıklarından
// vehicle değişkeni içeride çok biçimli(Polymorphic) yapı gösterebilir
func startVehicleEngine(vehicle VehicleContract) {
	vehicle.startEngine()
}

type Car struct {
	id   int
	name string
}

// Car yapılarına uygulanabilen startEngine metodu aynı zamanda interface sözleşmesinde de yer alır
func (v Car) startEngine() {
	fmt.Printf("'%s' motoru çalıştırıyor\n", v.name)
}

type Plane struct {
	id          int
	owner       string
	maxAltitude int
}

// Plane tipine tanımlanan startEngine metodu yine interface sözleşmesinde yer alan metoddur
func (p Plane) startEngine() {
	fmt.Printf("[%s] uçuş öncesi motorlar çalıştırılıyor. Hedef yükseklik %d\n", p.owner, p.maxAltitude)
}
