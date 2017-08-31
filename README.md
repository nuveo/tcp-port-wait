# tcp-port-wait
Wait container is accepting TCP connections.

## Example

```
echo "Waiting ip:port connecting"

# FIXME： timeout is not included
sh tcp-port-wait.sh 127.0.0.1 5432

echo "Done tcp-port-wait"
```

# Using tcpPortWait

## Example

```
var p Port
p.Timeout = time.Duration(10) * time.Second
timeout, err := p.Check("127.0.0.1:5432")
if err != nil {
	fmt.Printf(err)
}
if timeout {
	fmt.Printf("timeout")
}
```
