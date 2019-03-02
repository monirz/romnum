# romnum.
Roman Numeral Kata in Go. 

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
