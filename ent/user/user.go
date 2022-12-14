// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldIsVerified holds the string denoting the is_verified field in the database.
	FieldIsVerified = "is_verified"
	// FieldProfileImage holds the string denoting the profile_image field in the database.
	FieldProfileImage = "profile_image"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// EdgeFavourites holds the string denoting the favourites edge name in mutations.
	EdgeFavourites = "favourites"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FavouritesTable is the table that holds the favourites relation/edge.
	FavouritesTable = "favourites"
	// FavouritesInverseTable is the table name for the Favourites entity.
	// It exists in this package in order to avoid circular dependency with the "favourites" package.
	FavouritesInverseTable = "favourites"
	// FavouritesColumn is the table column denoting the favourites relation/edge.
	FavouritesColumn = "user_favourites"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserName,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldPasswordHash,
	FieldPhoneNumber,
	FieldIsVerified,
	FieldProfileImage,
	FieldEmail,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	PasswordHashValidator func(string) error
	// PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	PhoneNumberValidator func(string) error
	// DefaultIsVerified holds the default value on creation for the "is_verified" field.
	DefaultIsVerified bool
	// DefaultProfileImage holds the default value on creation for the "profile_image" field.
	DefaultProfileImage string
)
