package main

import (
	"fmt"
	"strconv"

	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	//AddAndReadHash()
	//AddLudwig()
	//aragorn := Card{NickName: "Aragorn", Greetings: "Well Met!", Price: 9, Attack: 10, Defense: 12, Owner: "Luktar"}
	//AddCard(aragorn, "card:45")
	card := GetCard("card:100")
	card.ToString()
}

//func AddAndReadHash() {
//	conn, err := redis.Dial("tcp", "localhost:6379")
//	if err != nil {
//		fmt.Println(err.Error())
//	} else {
//		defer conn.Close()
//		response := conn.Cmd("HMSET", "card:93", "nickName", "murlock", "greetings", "I'am ready, I'am not ready", "price", 5, "attack", 4, "defense", 4, "owner", "shammon")
//		if response.Err != nil {
//			fmt.Println(response.Err)
//		}
//		fmt.Println(response.String())
//		read, _ := conn.Cmd("HGETALL", "card:93").Map()
//		for k, v := range read {
//			fmt.Printf("%s\t%s\n", k, v)
//		}
//	}
//}

//func AddLudwig() {
//	conn, err := redis.Dial("tcp", "localhost:6379")
//	if err != nil {
//		fmt.Println(err.Error())
//	} else {
//		defer conn.Close()
//		pong := conn.Cmd("ping")
//		fmt.Println(pong.String())

//		response := conn.Cmd("set", "players:ludwig", "{\"nick\":ludwig,\"genre\":classic,\"SongCount\":98}")
//		if response.Err != nil {
//			fmt.Println(response.Err)
//		}
//		fmt.Println(response.String())
//	}
//}

func AddCard(card Card, id string) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer conn.Close()

		response := conn.Cmd("HMSET", id, "nickName", card.NickName, "greetings", card.Greetings, "price", card.Price, "attack", card.Attack, "defense", card.Defense, "owner", card.Owner)
		if response.Err != nil {
			fmt.Println(response.Err)
		}
		fmt.Println(response.String())
	}
}

func GetCard(id string) *Card {
	card := new(Card)
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer conn.Close()
		response, _ := conn.Cmd("HGETALL", id).Map()
		if response != nil {
			card.NickName = response["nickName"]
			card.Greetings = response["greetings"]
			card.Owner = response["owner"]
			card.Attack, _ = strconv.Atoi(response["attack"])
			card.Price, _ = strconv.Atoi(response["price"])
			card.Defense, _ = strconv.Atoi(response["defense"])
		}
	}
	return card
}

func (card *Card) ToString() {
	fmt.Printf("Nickname:%s\n", card.NickName)
	fmt.Printf("Greetings:%s\n", card.Greetings)
	fmt.Printf("Owner:%s\n", card.Owner)
	fmt.Printf("Price:%d\n", card.Price)
	fmt.Printf("Attack:%d\n", card.Price)
	fmt.Printf("Defense:%d\n", card.Defense)
}

type Card struct {
	NickName  string
	Greetings string
	Price     int
	Attack    int
	Defense   int
	Owner     string
}
