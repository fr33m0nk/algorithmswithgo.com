package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	//Using loop

  var sum int;
  for _, val := range numbers {
    sum += val
  }
  return sum

  //Using recursion
  // if len(numbers) == 0 {
  //   return 0
  // }

  // return numbers[0] + Sum(numbers[1:])
}
