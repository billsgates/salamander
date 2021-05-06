package room

import "errors"

var (
	ErrMaxCountExceed        = errors.New("max count exceed")
	ErrNotHost               = errors.New("only host are authorized for such actions")
	ErrNotMember             = errors.New("only members are authorized for such actions")
	ErrInvalidInvitationCode = errors.New("invalid invitation code")
	ErrAlreadyJoined         = errors.New("already joined")
)
