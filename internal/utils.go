package internal

import (
	"regexp"
	"strconv"
)

var maths Maths

func TimeDiff(time1, time2 float64) float64 {
  return maths.Fixhour(time2 - time1)
}

func ConvertStringToFloat(str string) float64 {
  re := regexp.MustCompile(`[-+]?\d*\.?\d+`)
  if len(re.FindAllString(str, -1)) <= 0 {
    return 0
  }
  val, err := strconv.ParseFloat(re.FindAllString(str, -1)[0], 64)
  if err != nil {
    return 0
  }
  return val
}

func HasMin(arg string) bool {
  re := regexp.MustCompile(`min`)
  if re.MatchString(arg) {
    return true
  }
  return false
}
