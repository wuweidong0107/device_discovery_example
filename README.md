# Run
```
$ go run examples/pinger/main.go 
Broadcasting to 239.0.0.0:9999
```

```
$ go run examples/listener/main.go
Listening to 239.0.0.0:9999
2023/08/14 22:02:36 main.go:35: 12 bytes read from 192.168.50.110:33096
2023/08/14 22:02:36 main.go:36: 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 0a              |hello world.|
```