type person struct {
    name string
    age int
}

var P person 

P.name = "Astaxie"  
P.age = 25 
fmt.Printf("The person's name is %s", P.name)