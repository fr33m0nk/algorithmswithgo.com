package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
  const charset = "0123456789ABCDEF"
	var res string
  var rem int
  for dec > 0 {
    // gets the base representation from charset using the remainder and resets dec to be quotient for new base
    rem = dec % base
    res = string(charset[rem]) + res
    dec = dec / base    
  }
  return res
}
