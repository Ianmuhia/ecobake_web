// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ecobake/ent/category"
	"ecobake/ent/schema"
	"ecobake/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescName is the schema descriptor for name field.
	categoryDescName := categoryFields[0].Descriptor()
	// category.NameValidator is a validator for the "name" field. It is called by the builders before save.
	category.NameValidator = categoryDescName.Validators[0].(func(string) error)
	// categoryDescCreatedAt is the schema descriptor for created_at field.
	categoryDescCreatedAt := categoryFields[1].Descriptor()
	// category.DefaultCreatedAt holds the default value on creation for the created_at field.
	category.DefaultCreatedAt = categoryDescCreatedAt.Default.(time.Time)
	// categoryDescIcon is the schema descriptor for icon field.
	categoryDescIcon := categoryFields[4].Descriptor()
	// category.IconValidator is a validator for the "icon" field. It is called by the builders before save.
	category.IconValidator = categoryDescIcon.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserName is the schema descriptor for user_name field.
	userDescUserName := userFields[0].Descriptor()
	// user.UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	user.UserNameValidator = userDescUserName.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescPasswordHash is the schema descriptor for password_hash field.
	userDescPasswordHash := userFields[4].Descriptor()
	// user.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	user.PasswordHashValidator = userDescPasswordHash.Validators[0].(func(string) error)
	// userDescPhoneNumber is the schema descriptor for phone_number field.
	userDescPhoneNumber := userFields[5].Descriptor()
	// user.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	user.PhoneNumberValidator = userDescPhoneNumber.Validators[0].(func(string) error)
	// userDescIsVerified is the schema descriptor for is_verified field.
	userDescIsVerified := userFields[6].Descriptor()
	// user.DefaultIsVerified holds the default value on creation for the is_verified field.
	user.DefaultIsVerified = userDescIsVerified.Default.(bool)
	// userDescProfileImage is the schema descriptor for profile_image field.
	userDescProfileImage := userFields[7].Descriptor()
	// user.DefaultProfileImage holds the default value on creation for the profile_image field.
	user.DefaultProfileImage = userDescProfileImage.Default.(string)
}