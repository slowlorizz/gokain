package exc

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	Data map[string]any

	ExceptionType struct {
		ID       string
		DefFatal bool // If this excetion-type is Fatal by Default, can be overwritten if specified otherwise when creating a new exception
	}

	Exception struct {
		Type  ExceptionType
		ID    string
		Msg   string
		Fatal bool
		Data  Data
	}
)

var (
	UserArgsException ExceptionType = ExceptionType{
		ID:       "UserArgs",
		DefFatal: true,
	}
)

/*
Creates and Returns new Exception-Pointer
--> Takes the types default Fatality

(?)	The "data" Parameter specifies Object Data, if no data is passed, define it as <nil>
*/
func New(t ExceptionType, id string, msg string, data Data) *Exception {
	e := &Exception{Type: t, ID: id, Msg: msg, Fatal: t.DefFatal}

	if data != nil {
		e.Data = data
	}

	return e
}

/*
Creates and Returns new Exception-Pointer
--> Overwrites the Type-Default Fatality to FATAL

(?)	The "data" Parameter specifies Object Data, if no data is passed, define it as <nil>
*/
func NewFatal(t ExceptionType, id string, msg string, data Data) *Exception {
	e := New(t, id, msg, data)
	e.Fatal = true
	return e
}

/*
Creates and Returns new Exception-Pointer
--> Overwrites the Type-Default Fatality to NOT FATAL

(?)	The "data" Parameter specifies Object Data, if no data is passed, define it as <nil>
*/
func NewNonFatal(t ExceptionType, id string, msg string, data Data) *Exception {
	e := New(t, id, msg, data)
	e.Fatal = true
	return e
}

func (e *Exception) String() string {
	severity := "ERROR"

	if e.Fatal {
		severity = "FATAL"
	}

	if e.Data != nil {
		jsonBytes, err := json.MarshalIndent(e.Data, "\t", "    ")

		if err != nil {
			panic(err.Error())
		}

		return fmt.Sprintf("[%s](%s.%s): %s\n%s", severity, e.Type.ID, e.ID, e.Msg, string(jsonBytes))

	} else {
		return fmt.Sprintf("[%s](%s.%s): %s", severity, e.Type.ID, e.ID, e.Msg)
	}
}

func (e *Exception) Error() error {
	return errors.New(e.String())
}

func (e *Exception) Raise() error {
	if e.Fatal {
		panic(e.String())
	} else {
		fmt.Println(e.String())
	}

	return e.Error()
}
