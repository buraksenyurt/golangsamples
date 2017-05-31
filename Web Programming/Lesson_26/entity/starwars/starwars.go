//Bu dosyayi GOPATH'in ya da GOROOT'un kuruldugu yere gore
// src altinda
// -entity
// -entity/starwars
// -entity/starwars/starwars.go
// seklinde olusturup derlemeliyiz.
package starwars

type model struct {
	Id       int
	Title    string
	Price    float32
	Category category
}

type category struct {
	Id   int
	Name string
}
