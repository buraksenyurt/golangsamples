/*
 Lesson 09
 Embedded type kullanımı
 gömülü türlerden yararlanarak çoklu türetme özelliğini kullanabiliriz
*/
package main

import (
	"fmt"
)

func main() {
	var zidane FootballPlayer
	zidane.self = Player{id: 10, nickName: "Zinadine Zidane"}
	zidane.position = "Midfield"
	zidane.abilities = []Ability{
		Ability{name: "shoot", power: 92},
		Ability{name: "high pass", power: 84},
	}
	zidane.abilities[1].useAbility()
	zidane.self.saySomething("What can I do sometimes. This is football.")
	zidane.abilities[0].useAbility()

	var tayson Boxer
	tayson.self = Player{id: 88, nickName: "Bulldog"}
	tayson.knockdownCount = 32
	tayson.abilities = []Ability{
		Ability{name: "defense", power: 76}, //virgül koymayınca derleme hatası verir ;)
	}
	tayson.self.saySomething("I will win this game")
	tayson.abilities[0].useAbility()
}

// oyuncuların ortak niteliklerini barındıran bir struct
type Player struct {
	id       int
	nickName string
}

// player yapısına monte edilmiş saySomething metodu
// oyuncunun bir şeyler söylemesi için kullanılabilecek bir metod
func (player *Player) saySomething(message string) {
	fmt.Printf("%s says that '%s'\n", player.nickName, message)
}

// oyuncuların farklı yeteneklerini tanımlayacak olan Ability isimli yapı
type Ability struct {
	name  string
	power int
}

// Ability yapısına monte edilmiş olan useAbility isimli bir metod
// oyuncunun bir yeteneğini kullandırmak için
func (ability *Ability) useAbility() {
	fmt.Printf("[%s] yeteneği kullanılıyor. Güç %d\n", ability.name, ability.power)
}

// Player ve Ability yapılarını gömülü tip olarak kullanan ve
// futbolcuları tanımlayan yapı
type FootballPlayer struct {
	position  string
	self      Player
	abilities []Ability
}

// farklı bir oyuncu tipi
type Boxer struct {
	knockdownCount int
	self           Player
	abilities      []Ability
}
