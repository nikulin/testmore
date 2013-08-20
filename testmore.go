// Package testmore provides primitives for easy writing testing functions.
// Provide Ok, Is, Isnt, Like, Unlike functions. All of them can optionally
// take message for note tests.
package testmore

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

var i int = 0

// func Ok checks if value is not 0, "", nil, false or any unitialized object
func Ok(t *testing.T, value interface{}, arg ...interface{}) {
	var res bool
	i++

	switch value.(type) {
	default:
		s := fmt.Sprintf("%v", value)
		res = s != "" && s != "[]" && s != "<nil>"
	case error:
		res = value != nil
	case string:
		res = len(value.(string)) > 0
	case []byte:
		res = len(value.([]byte)) > 0
	case []rune:
		res = len(value.([]rune)) > 0
	case int:
		res = value.(int) != 0
	case bool:
		res = value.(bool)
	case float32:
		res = value.(float32) != 0
	case float64:
		res = value.(float64) != 0
	}

	msg := getMsg(arg)

	if !res {
		t.Errorf("Not ok %d: %s\n%v==%s\n", i, msg, value, reflect.TypeOf(value))
	}
	t.Logf("Ok %d: %s\n", i, msg)
}

// func Is compares value versus sample and fails if they don't math
// Uses reflect.DeepEqual to compare complex objects
func Is(t *testing.T, value interface{}, sample interface{}, args ...interface{}) (res bool) {
	msg := getMsg(args)
	i++
	res = reflect.DeepEqual(sample, value)
	if !res {
		t.Errorf("Not ok %d: %s\nExpected:\t'%v'\nFound:\t'%v'\n", i, msg, sample, value)
	}
	t.Logf("Ok %d: %s\n", i, msg)
	//v  := reflect.ValueOf(i)
	return
}

// Like Is, but  func Isnt compares value versus sample and fails if they math
func Isnt(t *testing.T, value interface{}, sample interface{}, args ...interface{}) (res bool) {
	msg := getMsg(args)
	i++
	res = !reflect.DeepEqual(sample, value)
	if !res {
		t.Errorf("Not ok %d: %s\nExpecting:\t'%v' and found the same\n", i, msg, sample)
	}
	t.Logf("Ok %d: %s\n", i, msg)
	return
}

// Function Like cheks if regular expression needdle presents in stash
func Like(t *testing.T, stash interface{}, needle interface{}, arg ...interface{}) {
	i++
	re := regexp.MustCompile(needle.(string))
	res := re.Match([]byte(stash.(string)))
	msg := getMsg(arg)

	if !res {
		t.Errorf("Not ok %d: %s\n'%v'\n does not match '%v'\n", i, msg, needle, stash)
	}
	t.Logf("Ok %d: %s\n", i, msg)
}

// The inverse function for Like, fails if found needdle in stash
func Unlike(t *testing.T, stash interface{}, needle interface{}, arg ...interface{}) {
	i++
	re := regexp.MustCompile(needle.(string))
	res := re.Match([]byte(stash.(string)))
	msg := getMsg(arg)

	if res {
		t.Errorf("Not ok %d: %s\n'%v'\n match '%v'\n", i, msg, needle, stash)
	}
	t.Logf("Ok %d: %s\n", i, msg)
}

// Diag may be used for output some dignostic
func Diag(t *testing.T, msg interface{}) {
	t.Logf("# %s\n", msg)
}

func getMsg(arg ...interface{}) (msg string) {
	if (len(arg)) > 0 {
		msg = fmt.Sprintf("%s", arg[0])
		if msg == "[]" {
			msg = ""
		}
	} else {
		msg = ""
	}
	return
}
