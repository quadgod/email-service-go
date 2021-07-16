package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Email struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Provider    string             `bson:"provider" json:"provider"`
	To          string             `bson:"to" json:"to"`
	Cc          string             `bson:"cc" json:"cc"`
	From		string			   `bson:"from" json:"from"`
	Subject     string             `bson:"subject" json:"subject"`
	Body        string             `bson:"body" json:"body"`
	Type        string             `bson:"type" json:"type"`
	CreatedAt   *time.Time         `bson:"createdAt" json:"created_at"`
	LockedAt    *time.Time         `bson:"lockedAt" json:"locked_at"`
	SentAt      *time.Time         `bson:"sentAt" json:"sent_at"`
	CommittedAt *time.Time         `bson:"committedAt" json:"committed_at"`
	Attachments []string           `bson:"attachments" json:"attachments"`
	ReadyToSend bool               `bson:"readyToSend" json:"ready_to_send"`
}
