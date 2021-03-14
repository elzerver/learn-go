# Structs


Structs is a data structure wich allows us to compose different values of different types.

First define a type 

```
type person struct {
    first string
    last string
}
```

```
p1 := person {}  <- Composite literal>
```
```
p1 := person{
    first: "Peter",
    last: "Parker",
}
```