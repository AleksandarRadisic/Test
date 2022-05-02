package domain

type User struct {
	Username         string
	following        []User
	followers        []User
	followRequests   []User
	followerRequests []User
}
