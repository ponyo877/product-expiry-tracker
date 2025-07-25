// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package generated_sql

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Outbox struct {
	ID        int32
	EventType string
	Payload   json.RawMessage
	Processed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PrimeCheck struct {
	ID         int32
	UserID     int32
	NumberText string
	TraceID    sql.NullString
	MessageID  sql.NullString
	IsPrime    sql.NullBool
	Status     sql.NullString
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type User struct {
	ID        int32
	AuthToken string
	CreatedAt time.Time
	UpdatedAt time.Time
}
