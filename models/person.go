package models

// Persons map teste
var Persons = make(map[string]Person)

func init() {
	Persons["1"] = Person{ID: "1", Firstname: "Thiago", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}}
	Persons["2"] = Person{ID: "2", Firstname: "Rodrigo", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}}
	Persons["3"] = Person{ID: "3", Firstname: "Emad", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}}
	Persons["4"] = Person{ID: "4", Firstname: "Ubira", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}}
	Persons["5"] = Person{ID: "5", Firstname: "Valeria", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}}
}

// Person teste
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
