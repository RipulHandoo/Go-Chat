-- name: FollowUser :one

INSERT INTO user_followers (follower_id, following_id)
VALUES ($2, $1)
ON CONFLICT (following_id, follower_id)
DO NOTHING
RETURNING *;

-- name: UnfollowUser :one

DELETE FROM user_followers
WHERE
    follower_id = $1
    AND following_id = $2
RETURNING *;