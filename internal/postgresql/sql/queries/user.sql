-- name: GetUserByEmail :one
SELECT is_verified,
       created_at,
       email,
       user_name,
       id,
       password_hash,
       phone_number
FROM "users"
WHERE email = @email
  AND "users"."deleted_at" IS NULL
ORDER BY "users"."id"
LIMIT 1;

-- name: GetUserById :one
SELECT is_verified,
       created_at,
       email,
       user_name,
       id,
       updated_at,
       deleted_at,
       profile_image,
       password_hash,
       phone_number
FROM "users"
WHERE id = @id
  AND "users"."deleted_at" IS NULL
ORDER BY "users"."id"
LIMIT 1;
-- name: GetUnverifiedUserById :one
SELECT is_verified,
       created_at,
       email,
       user_name,
       id,
       updated_at,
       deleted_at,
       profile_image,
       password_hash,
       phone_number
FROM "users"
WHERE email = @email
  and is_verified = false
  AND "users"."deleted_at" IS NULL
ORDER BY "users"."id"
LIMIT 1;

-- name: ListUsers :many
SELECT is_verified,
       created_at,
       email,
       user_name,
       id,
       updated_at,
       deleted_at,
       profile_image,
       password_hash,
       phone_number
FROM "users"
WHERE "users"."deleted_at" IS NULL;

-- name: UpdateUser :one
UPDATE "users"
SET user_name    = @user_name,
    email        = @email,

    phone_number = @phone_number
WHERE id = @id
  AND "users"."deleted_at" IS NULL

Returning is_verified , created_at , email , user_name , id , updated_at , deleted_at , profile_image , phone_number;


-- name: DeleteUser :exec
UPDATE "users"
SET deleted_at = current_timestamp
WHERE id = @id
  AND "users"."deleted_at" IS NULL;


-- name: UpdateUserProfileImage :one
UPDATE "users"
SET profile_image = @profile_image
WHERE email = @email
  and users.deleted_at is null
Returning is_verified, created_at, email, user_name id, updated_at, deleted_at, profile_image, phone_number;


-- name: CreateUser :one
INSERT INTO users (created_at,
                   user_name,
                   email,
                   phone_number,

                   password_hash,
                   profile_image)
VALUES (current_timestamp,
        @user_name,
        @email,
        @phone_number,

        @password_hash,
        @profile_image)
Returning is_verified, created_at, email, user_name,  id, updated_at, deleted_at, profile_image, phone_number;


-- name: UpdateUserStatus :one
UPDATE "users"
SET is_verified = true
WHERE email = @email
  and users.deleted_at is null
Returning is_verified, created_at, email, user_name, id, updated_at, deleted_at, profile_image, phone_number;


-- name: UpdateUserPassword :exec
UPDATE "users"
SET password_hash = @password_hash
WHERE id = @id
  and users.deleted_at is null;


