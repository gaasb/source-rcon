# Source RCON
### Requirement
```text
golang 1.21+
```
## To use
```go
func main() {
	client := rcon.NewClient("public_ip", 27015, "password", rcon.WithDeadline(time.Second*3))
	defer client.Close()

	if err := client.Auth(); err != nil {
		log.Fatal(err)
	}

	if response, err := client.Execute("echo hello"); err != nil {
		log.Fatal(err)
	} else {
		log.Println(response)
	}
}
```
