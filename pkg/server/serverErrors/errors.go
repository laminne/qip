package serverErrors

// Qip API Error Definition
/*
Qip API エラー一覧

一般エラー:
UnAuthorized/認証情報がありません
FailedValidation/バリデーションに失敗
NotFound/APIが存在しません
InternalError/内部エラー

エンドポイント固有のエラー
PostTooLong/投稿が長すぎます
InvalidTarget/(マージやウォッチ、リアクション)ターゲット先が正しく指定されていません
AlreadyWatched/すでにウォッチしています
*/

type CommonAPIErrorResponseJSON struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

var NotFoundErrorResponseJSON = CommonAPIErrorResponseJSON{
	Type:    "NotFound",
	Message: "エンドポイントが存在しません",
}
var InternalErrorResponseJSON = CommonAPIErrorResponseJSON{
	Type:    "InternalError",
	Message: "内部エラーが発生しました",
}
var InvalidRequestErrorResponseJSON = CommonAPIErrorResponseJSON{
	Type:    "InvalidRequest",
	Message: "リクエストボディの内容が間違っています",
}
var UnAuthorizedRequestErrorResponseJSON = CommonAPIErrorResponseJSON{
	Type:    "UnAuthorized",
	Message: "ログインしていません",
}

// Post

var PostNotFoundErrorResponseJSON = CommonAPIErrorResponseJSON{
	Type:    "PostNotFound",
	Message: "投稿が存在しません",
}

// User

var UserNotFoundErrorResponseJSON = CommonAPIErrorResponseJSON{
	Type:    "UserNotFound",
	Message: "ユーザーが存在しません",
}
