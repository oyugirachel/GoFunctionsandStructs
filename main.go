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
	d:=time.Now()
	c:= &Food{"Rachel","Spicy",2}
	fmt.Printf("%s ate her %dth food\n", c.name, much(c))
	fmt.Printf("%s  loves %s foods\n", c.name, c.flavour)
	fmt.Printf( "today is on %s\n", d )
	foo := map[string]interface{}{
		"Rachel": 42,
	}
	timeMap(foo)
	fmt.Println(foo)


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
func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}





