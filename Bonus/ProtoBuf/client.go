package main

import (
	"fmt"

	data "message/Southwind"

	"github.com/golang/protobuf/proto"
)

func main() {
	taverna := data.Game{Player: []*data.Player{
		{
			NickName: "Leksar",
			PlayerId: 10,
			Type:     data.PlayerType_SEMI,
			Weapons: []*data.Player_Weapon{
				{Name: "Sword", Ability: "High level sword"},
				{Name: "Machine Gun", Ability: "7.65mm"},
			},
		},
		{
			NickName: "Valira",
			PlayerId: 12,
			Type:     data.PlayerType_SEMI,
			Weapons: []*data.Player_Weapon{
				{Name: "Poison Bottle", Ability: "Dangeres green"},
			},
		},
	},
	}
	sData, err := proto.Marshal(&taverna)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(sData)
		fmt.Println(string(sData))
	}

	dsData := &data.Game{}
	err = proto.Unmarshal(sData, dsData)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, p := range dsData.Player {
			fmt.Println(p.NickName)
			for _, w := range p.Weapons {
				fmt.Printf("\t%s\t%s\n", w.Name, w.Ability)
			}
		}
	}
}
