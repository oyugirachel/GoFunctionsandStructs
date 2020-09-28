package main

import (
	"fmt"
	"math"
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
type geometry interface{
	area() float64
	perim() float64
}
type circle struct{
	radius float64
}

type rect struct{
	width , length float64

}


func (r rect) area() float64{
	return r.width*r.length
}
func(c circle) area() float64{
	return math.Pi*c.radius*c.radius
}
func(r rect)  perim() float64{
	return 2*r.width +2*r.length
}
func(c circle) perim() float64{
	return 2*math.Pi*c.radius
}
func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	fmt.Println(add(234,789))
	fmt.Println(subtract(789,600))
	region, continent := location("Nairobi")
	fmt.Printf("Rachel lives in %s, %s\n", region, continent)
	d:=time.Now()
    r:=rect{7,8}
    c:=circle{45}
    measure(r)
    measure(c)


	h:= &Food{"Rachel","Spicy",2}
	fmt.Printf("%s ate her %dth food\n", h.name, much(h))
	fmt.Printf("%s  loves %s foods\n", h.name, h.flavour)
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





