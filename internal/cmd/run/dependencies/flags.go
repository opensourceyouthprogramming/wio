package dependencies

import (
    "errors"
    "fmt"
    "regexp"
    "strings"
)

var placeholderMatch = regexp.MustCompile(`^\$(|optional)\([a-zA-Z_-][a-zA-Z0-9_]*\)$`)

// Verifies the placeholder syntax
func IsPlaceholder(flag string) bool {
    return placeholderMatch.MatchString(strings.Trim(flag, " "))
}

// matches a flag by the requested flag
func TryMatch(key, given string) (string, bool) {
    pat := regexp.MustCompile(`^` + key + `(|=|->).*$`)
    if !pat.MatchString(given) {
        return "", false
    }
    if strings.Contains(given, "->") {
        return strings.Split(given, "->")[1], true
    }
    return given, true
}

// fill placeholder flags and error if some are left unfilled
func fillPlaceholders(givenFlags, requiredFlags []string) ([]string, error) {
    var ret []string
    for _, required := range requiredFlags {
        if !IsPlaceholder(required) {
            ret = append(ret, required)
            continue
        }

        newRequest := strings.Replace(required, "$optional(", "$(", 1)

        // look for a match
        for _, given := range givenFlags {
            key := newRequest[2 : len(newRequest)-1]
            if res, match := TryMatch(key, given); match {
                ret = append(ret, res)
                goto Continue
            }
        }

        if strings.Contains(required, "$optional(") {
            continue
        }

        return nil, errors.New(fmt.Sprintf("placeholder definition \"%s\" unfilled in ", required) + "%s")

    Continue:
        continue
    }
    return ret, nil
}

// this fills global flags if they are requested
func fillDefinition(givenFlags, requiredFlags []string, optional bool) ([]string, error) {
    var ret []string
    for _, required := range requiredFlags {
        for _, given := range givenFlags {
            newRequired := strings.Replace(required, "-D", "", 1)
            newGiven := strings.Replace(given, "-D", "", 1)

            if res, match := TryMatch(newRequired, newGiven); match {
                ret = append(ret, res)
                goto Continue
            }
        }

        if !optional {
            return nil, errors.New("%s" + fmt.Sprintf(" definition \"%s\" unfilled in ", required) + "%s")
        }

    Continue:
        continue
    }

    return ret, nil
}
