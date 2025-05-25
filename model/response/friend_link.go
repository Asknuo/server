package response

import "sever/database"

type FriendLinkInfo struct {
	List  []database.FriendLink `json:"list"`
	Total int64                 `json:"total"`
}
