package greetings

import (
	"regexp"
	"testing"
)

// funciones de tipo test reciben como parametro reciben un objeto de tipo testing
func TestHello(t *testing.T) {
	name := "Lizeth"
	//creamos una expresión regular
	want := regexp.MustCompile(`\b` + name + `\b`) //  busca una coincidencia exacta en el nombre con esta expresión regular
	msg, err := Hello("Lizeth")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") =  %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello(") = %q, %v, quiere "", error`, msg, err)
	}

}
