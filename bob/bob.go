package bob

import "strings"

const testVersion = 2

func Hey(greeting string) string {
  var response string

  greeting = strings.TrimSpace(greeting)

  switch {
  case greeting == "":
    response = "Fine. Be that way!"
  case greeting == strings.ToUpper(greeting) && greeting != strings.ToLower(greeting):
    response = "Whoa, chill out!"
  case strings.HasSuffix(greeting, "?"):
    response = "Sure."
  default:
    response = "Whatever."
  }

  return response
}
