package ports

import (
	"github.com/tarm/serial"
)

//Writer .
type Writer struct {
	Name string
	Sw   *serial.Port
}

//NewWriter .
func NewWriter(name string) (Writer, error) {
	w := &serial.Config{Name: name, Baud: 115200}
	sw, err := serial.OpenPort(w)
	if err != nil {
		return Writer{}, err
	}
	return Writer{name, sw}, nil
}

//GetName .
func (w *Writer) GetName() string {
	return w.Name
}

//GetSerialPort .
func (w *Writer) GetSerialPort() *serial.Port {
	return w.Sw
}

//Writing .
func (w *Writer) Writing(messages []string) {
	(w.Sw).Write([]byte(unite(putLimiter(messages))))
}

func putLimiter(m []string) []string {
	var messages []string
	for _, message := range m {
		messages = append(messages, "B1"+message+"O75")
	}
	return messages
}
func unite(m []string) string {
	var messages string
	for _, message := range m {
		messages += message
	}
	return messages
}
