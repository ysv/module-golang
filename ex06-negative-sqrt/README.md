# Negative Sqrt with Errors handling

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers

Create a new type
```go
type ErrNegativeSqrt float64
```

and make it an error by giving it a
```go
func (e ErrNegativeSqrt) Error() string
```
