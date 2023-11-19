package main

import (
	"fmt"
	"os"
)

type Friend struct {
	id         uint
	name       string
	address    string
	occupation string
	reason     string
}

var friends []Friend = []Friend{
	{
		id:         1,
		name:       "Vormes",
		address:    "Purwakarta",
		occupation: "Software Developer",
		reason:     "Karena Golang populer!",
	},
	{
		id:         2,
		name:       "Ali",
		address:    "Semarang",
		occupation: "Web Developer",
		reason:     "Karena Golang seru!",
	},
	{
		id:         3,
		name:       "Daus",
		address:    "Semarang",
		occupation: "Hardware Engineer",
		reason:     "Karena Golang sangat dicari!",
	},
	{
		id:         4,
		name:       "Chia",
		address:    "Bekasi",
		occupation: "Data Scientist",
		reason:     "Karena Golang powerfull!",
	},
}

func main() {
	var arg = os.Args

	if len(arg) < 2 {
		fmt.Println("Masukkan nama!")
		return
	}

	findMyFriend(arg[1])
}

func findMyFriend(name string) {
	found := false
	for _, eachFriend := range friends {
		if name == eachFriend.name {
			fmt.Println("ID:", eachFriend.id)
			fmt.Println("Nama:", eachFriend.name)
			fmt.Println("Alamat:", eachFriend.address)
			fmt.Println("Pekerjaan:", eachFriend.occupation)
			fmt.Println("Alasan:", eachFriend.reason)
			found = true
		}
	}
	if !found {
		fmt.Println(name, "tidak ditemukan!")
	}
}
