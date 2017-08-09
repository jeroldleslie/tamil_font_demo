package commons

import (
	"labix.org/v2/mgo"
)

func GetDBSession() *mgo.Session {
	s, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	//defer s.Close()

	s.SetMode(mgo.Monotonic, true)
	return s
}
