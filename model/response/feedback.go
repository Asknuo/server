package response

import "sever/database"

type FeedbackInfo struct {
	List  []database.Feedback `json:"list"`
	Total int64               `json:"total"`
}
