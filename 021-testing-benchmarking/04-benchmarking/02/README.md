# Remember to BET

- Benchmark
- Example
- Test

```
func BenchmarkXxx(b *testing.B)

func ExampleXxx()

func TestXxx(t *testing.T)
```

# Commands

```
go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```
