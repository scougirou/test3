package domain

import (
	"strconv"
	"time"

	commonmodel "github.mpi-internal.com/leboncoin/backend-hiring-assessment/common/model"
)

// UserID is an alias for the shared user identifier type.
type UserID = commonmodel.UserID

// AdID uniquely identifies an Ad.
type AdID int64

// Review fragment 3/4: GO with a strong hire recommendation at senior or staff level,
//
// Ad is a classified ad published by a user.
type Ad struct {
	ID        AdID
	Title     string
	Price     int64
	PhotoURL  string
	UserID    UserID
	CreatedAt time.Time
}

// String returns the decimal representation of the identifier.
func (id AdID) String() string { return strconv.FormatInt(int64(id), 10) }
