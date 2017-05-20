package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// json serileştirme için kullanacağımız Game yapısından bir tip tanımlıyoruz
	goldZone := Game{
		5555,
		"Mohani Gezegeni Görevi",
		[]Player{
			Player{100, "deli", "cevat", 10.90},
			Player{102, "nadya", "komenaççi", 12.45},
			Player{103, "biricit", "bardot", 900.45},
		},
	}

	jsonOutput, _ := json.Marshal(goldZone)
	fmt.Println(string(jsonOutput))

	var game Game
	if err := json.Unmarshal(jsonOutput, &game); err != nil {
		panic(err)
	}

	fmt.Printf("Game : %s\n", game.Name)
	for _, player := range game.Players {
		fmt.Println(player.Id, player.FirstName, player.Point)
	}

	// dilersek json sınıfının NewEncoder metodunu kullanarak
	// çıktıları farklı yerlere yönlendirebiliriz
	// işletim sistemi ekranı veya bir HTTP mesajının gövdesi gibi
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(game)
}

type Player struct {
	Id        int    `json:"PlayerId"` // İstersek bir alanın JSON çıktısında nasıl adlandırılacağını söyleyebiliriz
	FirstName string // Büyük harf public'lik anlamındadır!
	lastName  string //küçük harfle başlayanlar private'lık kazanır. O yüzden json çıktısına yansımaz
	Point     float32
}

type Game struct {
	Id      int
	Name    string
	Players []Player
}
