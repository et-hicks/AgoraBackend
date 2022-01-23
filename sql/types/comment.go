package types


type Comment struct {
	Id              uint64
	CreatedTime     uint64 // Milliseconds since Epoch
	LastUpdatedTime uint64 // Milliseconds since Epoch
	ThreadId		uint64
	CreatorId		uint64
	ParentCommentId uint64
	Comment 		string
	IsEdited		bool
	IsReported 		bool
	Likes			uint64
	Dislikes		uint64
	Befs			map[string]interface{}  // Back End Functional Stuff - a JSON to hedge against the future

}
