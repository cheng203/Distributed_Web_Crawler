package model

import (
	"encoding/json"
)

type Profile struct {
	Name     string
	Age      int
	Height   int
	Weight   int
	Income   string
	Marriage string
	Origin   string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
