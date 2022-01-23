package types

// Thread - UI component to display to the users
// unique key Id
// unique key UrlUuid

type Thread struct {
	Id uint64
	CreatedTime uint64     		 // Milliseconds since Epoch
	LastUpdatedTime uint64		 // Milliseconds since Epoch
	Title string
	CreatorId int64       		 // SQL references to users table
	Likes int64
	Dislikes int64
	Clicks uint64
	Watchers uint64
	UrlUuid string          	 // URL Path for the frontend to know how to get this thread
	ImageUrl string        		 // Main image of the thread
	IsPublic bool          		 // does/can this show up on the front page naturally?
	IsReported bool        		 // when this is true the thread id will go into the reported SQL table
	Description string           // User defined Description when creating the thread
	Befs map[string]interface{}  // Back End Functional Stuff - a JSON to hedge against the future

}

// A thread will have people that can contribute to the discussion.
//This acts as the many-to-many relationships between the threads and the users.

type AccessLevel int64
const (
	Undefined AccessLevel = iota
	Admin
	Moderator
	Creator
	Commenter
	Revoked
	Blocked
	Viewer
)

type Contribute struct {
	Id uint64
	CreatedTime uint64       // milliseconds since epoch
	LastUpdatedTime uint64   // milliseconds since epoch
	ContributorId int64      // reference to the Users table
	ThreadId uint64          // reference to the Threads table
	Access AccessLevel       // to define how they can interact with the thread
}

// Topic - many topics to one thread relationship.
// Something similar to subreddits
// examples include Biology, Chemistry, Physics
type Topic struct {
	Id uint64
	CreatedTime uint64     		 // Milliseconds since Epoch
	LastUpdatedTime uint64		 // Milliseconds since Epoch
	ThreadId uint64              // reference to thread
	Topic string 				 // threads will have keywords and topics
}

// Hashtags - many hashtags to one thread relationship.
// Hashtags are hashtags of course, but while a thread might only be posted to one
// or two very specific topics, they can have many hashtags that can also describe the thread
// #python #dataanalysis are both applicable on chemistry, physics, or ML threads but wont be posted to those topics
type Hashtags struct {
	Id uint64
	CreatedTime uint64     		 // Milliseconds since Epoch
	LastUpdatedTime uint64		 // Milliseconds since Epoch
	ThreadId uint64              // reference to thread
	Hashtag string 				 // threads will have keywords and topics
}

// WatcherStatus - on thread update
//(every, say, comments rounded to nearest 10^n,
// meaning they would get emailed every 10, 20, 30, ... then 100, 200, ... comments)
type WatcherStatus int64
const (
	Notification WatcherStatus = iota
	Email
	Summary      // After the thread has died down, notify top X comments
	EmailSummary // After the thread has died down, notify and Email top X comments
)

// Watchers - Many watchers to one thread relationship
// unique key (ThreadId, WatcherId)
type Watchers struct {
	Id uint64
	CreatedTime uint64     		 // Milliseconds since Epoch
	LastUpdatedTime uint64		 // Milliseconds since Epoch
	ThreadId uint64              // reference to thread
	WatcherId int64				 // SQL references to users table
	Status WatcherStatus
}

