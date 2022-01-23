package types

type AccountType int64
const (
	Contributor AccountType = iota  // thread creator
	Commentary						// can make comments but not threads
	Login 							// login and just view stuff
	Employee						// Employees of Agora
)

// User - we will deal with phone code or password on the site as is, not needing a whole column dedicated to it
// unique key ID
// unique key email
// unique key Username
// unique key PhoneNumber
type User struct {
	Id              uint64
	CreatedTime     uint64 // Milliseconds since Epoch
	LastUpdatedTime uint64 // Milliseconds since Epoch
	FirstName		string
	LastName		string
	Username		string
	Email			string
	Password		string
	PhoneNumber		uint64
	PhoneCode		uint64
}
