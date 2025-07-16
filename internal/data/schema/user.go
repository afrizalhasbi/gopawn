package schema

type User struct {
	Uuid    string
	Name    string
	Created string
	Updated string
	Elo     int16
	Games   int32
}
