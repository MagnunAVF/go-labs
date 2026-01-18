package user

// The Bad: Uses vague names and redundant context.
// type User struct {
//     fName string // first name
//     lName string // last name
//     d     int    // days since join
// }
// func Get(u *User) {
//     // ...
// }

// The Good: Names reveal intent without comments; package name provides sufficient context.
type Account struct {
	FirstName       string
	LastName        string
	DaysSinceJoined int
}

func FetchAccount(acc *Account) {
	// ...
}
