## Struct and method
In Golang, a struct is a composite data type that groups together variables, known as fields, under a single name. It allows you to create a custom data type by combining different types of values. Each field in a struct can have a different data type.
```go
// Defining a struct with a method
type Circle struct {
    Radius float64
}

// Method associated with the Circle struct
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    // Creating an instance of the Circle struct
    circle1 := Circle{Radius: 5.0}

    // Calling the Area method
    area := circle1.Area()
    fmt.Println("Area of the circle:", area)
}
```
Structs can also be used to represent more complex data structures, and they are often employed in Go for defining data models, such as those used in database interactions or API responses. Additionally, methods can be associated with structs to provide behavior associated with the data.
