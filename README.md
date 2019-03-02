# romnum.
A Number to Roman Numerals converter written in *Go* followed by *TDD* pattern. Wrote this after getting inpired by watching [this](https://www.youtube.com/watch?v=983zk0eqYLY) video on Youtube. 

### Install
`go get -u github.com/monirz/romnum`

### Example

```go
func main() {

	result, err := romnum.Convert(2019)
	if err != nil {
		//handle error
	}

	fmt.Println(result)
	//Prints: MMXIX
}
```
