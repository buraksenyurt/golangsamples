/*
 Lesson_08
	Temel struct kullanımı
*/
package main

import (
	"fmt"
)

func main() {
	//struct' ları farklı şekillerde üretebiliriz.
	// var ile tanımlayıp niteliklerine sonradan değer atayabiliriz
	var veyron Vehicle
	veyron.id = 1
	veyron.name = "bugatti veyron"
	veyron.color = "black"

	// yine var ile tanımlayıp niteliklerine o anda değer atayabiliriz
	var gayyardo Vehicle = Vehicle{2, "lamborghini gayyardo", "gold"}

	// dinamik olarak tanımlarken atayabiliriz
	sesto := Vehicle{id: 3, name: "ferrari testo elemento", color: "red"}

	write_vehice_to_console(&veyron)
	write_vehice_to_console(&gayyardo)
	write_vehice_to_console(&sesto)

	// Vehicle tipinden bir slice tanımlıyoruz
	var some_vehicles []Vehicle
	// append fonksiyonu yardımıyla slice içerisine Vehicle örneklerini ekliyoruz
	some_vehicles = append(some_vehicles, sesto)
	some_vehicles = append(some_vehicles, gayyardo)
	some_vehicles = append(some_vehicles, veyron)
	write_all_vehicle_to_console(some_vehicles) //slice içeriğini fonksiyona parametre olarak geçiyoruz

	veyron.move(12, 10, -8) // veyron isimli yapı örneği üzerinden move metodunu kullanıyoruz
	for _, vehicle := range some_vehicles {
		vehicle.move(-3, 4, 10) // some_vehicles slice'ı içerisindeki tüm Vehicle nesneleri üzerinden move metodunu çağırıyoruz
	}
}

// fonksiyonlara struct tipini pointer üzerinden aktardık
func write_vehice_to_console(v *Vehicle) {
	fmt.Printf("(%d)-%s,%s\n", v.id, v.name, v.color)
}

func write_all_vehicle_to_console(vehicles []Vehicle) {
	for _, vehicle := range vehicles {
		write_vehice_to_console(&vehicle)
	}
}

// id, name ve color nitelikleri olan bir struct tanımladık
type Vehicle struct {
	id          int
	name, color string
}

// move isimli bir metod(kavramsal olarak fonksiyon değil metod ismini kullandığımıza dikkat edelim) oluşturduk
// Bu metod Vehicle tipinden yapılara uygulanabiliyor
func (vehicle Vehicle) move(x, y, z int) {
	fmt.Printf("%s, (%d:%d:%d) lokasyonuna gidiyor\n", vehicle.name, x, y, z)
}
