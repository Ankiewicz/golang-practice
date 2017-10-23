package main

import "fmt"

func main() int {
  nums := [3]int{3,2,3}
  target := 6

  for i := 0; i < len(nums); i++ {
    for j := i + 1; j < len(nums); j++ {
        var total int
        total = nums[i] + nums[j]
        if total == target {
          return[2]int{i, j}
          fmt.Println(successVar)
          return  successVar
        }
    }
  }
}
