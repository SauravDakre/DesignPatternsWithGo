package decorator

import (
	"strings"
)

type stringManipulator func(string) string

func toLower(s string) string {
	return strings.ToLower(s)
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func toCamelCase(f stringManipulator) stringManipulator {
	return func(q string) string {
		q = strings.TrimSpace(q)
		ar := strings.Split(q, " ")
		var u string
		for _, v := range ar {
			u += (" " + toUpper(string(v[0])) + toLower(string(v[1:])))
		}
		return f(strings.TrimSpace(u))
	}
}

func appendPrefix(f stringManipulator) stringManipulator {
	return func(s string) string {
		return f("Hello " + s)
	}
}

func appendCustomPrefix(prefix string) func(m stringManipulator) stringManipulator {
	return func(f stringManipulator) stringManipulator {
		return func(s string) string {
			return f(prefix + " " + s)
		}
	}
}

func identity(s string) string {
	return s
}
