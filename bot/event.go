package bot

// ReactionType containing the types of external events
type ReactionType int

const (
	// Comment reactions
	PositiveComment ReactionType = iota
	NegativeComment

	PostLiked
	PostDisliked

	// Social actions
	FriendAdded
	FriendRemoved
)

// BotEvent represents an incoming event
type BotEvent struct {
	Type ReactionType
}
