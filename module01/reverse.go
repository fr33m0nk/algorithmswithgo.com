package module01

// import ("strings")

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
  // var res string
  // for _, r := range word {
  //   res = string(r) + res
  // }
	// return res

  // var sb strings.Builder
  // runes := []rune(word)
  // for i:= len(runes)-1; i >= 0; i-- {
  //   sb.WriteRune(runes[i])
  // }
  // return sb.String()

  runes := []rune(word)
  for i, j:= 0, len(runes) -1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}
