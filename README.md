go test -bench=. -benchmem .

go test -bench=. -memprofile=mem.out -cpuprofile=cpu.out -x .
either
    go tool pprof util.test.exe cpu.out
or
    go tool pprof -http :8082 util.test.exe cpu.out