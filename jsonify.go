package traits

import "encoding/json"

type jsonify interface {
	ToJSON() (string, error)
	setJsonify(i interface{})
}

// Jsonify trait extends struct with ToJson marshalling method
type Jsonify struct {
	self interface{}
}

func (js *Jsonify) setJsonify(i interface{}) {
	js.self = i
}

// ToJSON serializes struct to json format
func (js *Jsonify) ToJSON() (string, error) {
	b, err := json.Marshal(js.self)
	return string(b), err
}
