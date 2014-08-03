package log

import (
	"errors"
	"os"
	"testing"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO: real tests :)
func TestLog(t *testing.T) {
	l := New(os.Stderr, DEBUG, "")
	l.Debug("something happened")
	l.Info("hello %s", "Tobi")
	l.Error("boom something exploded")

	l.SetPrefix("myapp")
	l.Info("something")

	Debug("something")
	Fatal("hello %s %s", "tobi", "ferret")
	Error("stan smith\n")
	Warning("roger")

	err := errors.New("francine")
	Errorif(nil)
	l.Errorif(nil)

	Errorif(err)
	l.Errorif(err)

	Panicif(nil)
	Fatalif(nil)

	l.Panicif(nil)
	l.Fatalif(nil)
}
