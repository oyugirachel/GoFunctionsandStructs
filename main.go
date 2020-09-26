package main

import (
	"fmt"
	"time"
)

type Food struct {
	name, flavour string
	quantity int

}
func much(b *Food) int{
	b.quantity++
	return b.quantity
}


func main() {
	fmt.Println(add(234,789))
	fmt.Println(subtract(789,600))
	region, continent := location("Nairobi")
	fmt.Printf("Rachel lives in %s, %s\n", region, continent)
	d:=time.Date(2020, time.September,
		26, 21, 34, 01, 0, time.UTC)
	c:= &Food{"Rachel","Spicy",2}
	fmt.Printf("%s ate her %dth food\n", c.name, much(c))
	fmt.Printf("%s  loves %s foods\n", c.name, c.flavour)
	fmt.Printf( "today is on %s", d )


}
func add(x int , y int) int{
	return x+y
}
func subtract(p int, f int) int{
	return p-f
}
func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Mombasa", "Kisumu", "Nairobi":
		region, continent = "Nairobi", "Kenya"
	case "Nakuru", "Kitale":
		region, continent = "Nakuru", "Central Kenya"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}





