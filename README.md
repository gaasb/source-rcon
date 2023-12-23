# Source RCON
### Requirement
```text
golang 1.21+
```
## To use
```go
func main() {
	client := rcon.NewClient("localhost", 27015, "password", rcon.WithDeadline(time.Second*3))
	defer client.Close()

	if err := client.Auth(); err != nil {
		log.Fatal(err)
	}

	if response, err := client.Execute("say hello"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(response)
	}
}
```
