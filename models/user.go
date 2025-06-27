package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type ContactInvite struct {
	ID         uuid.UUID `json:"id" db:"id"`
	SenderID   uuid.UUID `json:"sender_id" db:"sender_id"`
	ReceiverID uuid.UUID `json:"receiver_id" db:"receiver_id"`
	Status     string    `json:"status" db:"status"` // pending, accepted, rejected
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type Friendship struct {
	ID        uuid.UUID `json:"id" db:"id"`
	User1ID   uuid.UUID `json:"user1_id" db:"user1_id"`
	User2ID   uuid.UUID `json:"user2_id" db:"user2_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Group struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedBy uuid.UUID `json:"created_by" db:"created_by"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type GroupMember struct {
	ID       uuid.UUID `json:"id" db:"id"`
	GroupID  uuid.UUID `json:"group_id" db:"group_id"`
	UserID   uuid.UUID `json:"user_id" db:"user_id"`
	JoinedAt time.Time `json:"joined_at" db:"joined_at"`
}

type GroupInvite struct {
	ID              uuid.UUID `json:"id" db:"id"`
	GroupID         uuid.UUID `json:"group_id" db:"group_id"`
	InvitedUserID   uuid.UUID `json:"invited_user_id" db:"invited_user_id"`
	InvitedByUserID uuid.UUID `json:"invited_by_user_id" db:"invited_by_user_id"`
	Status          string    `json:"status" db:"status"` // pending, accepted, rejected
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
