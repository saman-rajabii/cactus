# cactus

Cactus is a parallel function executor. It's useful when you are going to execute a function on an array elements.

As you see in the example below, it benefits from generic mode of Golang. The couple of types that are passed to the beginning of the Parallel function are the type of array elements which is equal to input of function and type of function's response which is equal to the final result that pushed to the passed channel.

Input description:

- Array of input items
- Function that executes on each element of array
- Channel for accessing the outputs
- Concurrency input option in order to defining how many functions execute concurrently.

How to use:

```go
...

func do(in int, channel chan string) {
	time.Sleep(5 * time.Second)
	channel <- strconv.Itoa(in)
}

func main() {
	ch := make(chan string)

	cactus.Parallel[int, string]([]int {1,2,3,4,5,6,7,8,9}, do, ch, 3)

func main() {
	ch := make(chan string)

	cactus.Parallel[int, string]([]int {1,2,3,4,5,6,7,8,9}, do, ch, 3)

  for val := range ch {
		fmt.Println(val)
	}
}

```
