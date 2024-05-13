package main
import "fmt"

func findMaxK(numbers []int, k int) int {
  var nameslice = numbers[:]
  var max int
  var index int
  
  for i := 1; i <= k; i++ {
    max = nameslice[0]
    index = 0
    for j := 0; j < len(nameslice); j++ {
      if nameslice[j] > max {
        max = nameslice[j]
        index = j
      }
    }
    
    nameslice[index] = 0
  }
  
  return max
}

func main() {
  var k int
  fmt.Println("Enter k:")
  fmt.Scanf("%d\n", &k)
  
  var numbers = [...]int{4, 5, 16, 6, 1, 34}
  fmt.Println("Array: ", numbers)
  max := findMaxK(numbers[:], k)
  fmt.Println("Max: ", max)
}