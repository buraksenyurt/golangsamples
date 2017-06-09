/*
	Bir go kütüphanesini shared library olarak derleyip
	Ruby kodundan kullanma
	Konu detayı için http://www.buraksenyurt.com/post/golang-bir-go-paketini-ruby-den-cagirmak
*/

package main

import "C"

import (
	"math"
)

//export CircleSpace
func CircleSpace(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

func main() {}
