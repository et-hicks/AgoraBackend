package types


// TODO: Create this

type ReportedType int64
const (
	ReportedComment ReportedType = iota
	ReportedUser
	ReportedThread
)

type Report struct {
	Id uint
	CreatedTime uint64     		 // Milliseconds since Epoch
	LastUpdatedTime uint64		 // Milliseconds since Epoch
	ReporterId int64			 // references users table
	ReportedId int64			 // references users, comments, or threads table

}
