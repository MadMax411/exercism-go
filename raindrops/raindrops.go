package raindrops

import "fmt"

const testVersion = 2

func Convert(number int) string {
  rain_speak := ""

  rainSpeakDictionary := map[int]string {
    3: "Pling",
    5: "Plang",
    7: "Plong",
  }

  for i := 3; i < 8; i +=2 {
    if number % i == 0 {
      rain_speak += rainSpeakDictionary[i]
    }
  }

  if rain_speak == "" {
    rain_speak = fmt.Sprintf("%v", number)
  }

  return rain_speak
}
