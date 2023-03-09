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

type commonAPIErrorResponseJSON struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

var UnAuthorizedErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "UnAuthorized",
	Message: "認証情報がありません",
}
var FailedValidationErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "FailedValidation",
	Message: "バリデーションに失敗しました",
}
var NotFoundErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "NotFound",
	Message: "エンドポイントが存在しません",
}
var InternalErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "InternalError",
	Message: "内部エラーが発生しました",
}

var PostTooLongErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "PostTooLong",
	Message: "投稿が長すぎます",
}
var InvalidTargetErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "InvalidTarget",
	Message: "ターゲット先が正しく指定されていません",
}
var AlreadyWatchedErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "AlreadyWatched",
	Message: "すでにウォッチしています",
}

var InvalidRequestErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "InvalidRequest",
	Message: "リクエストボディの内容が間違っています",
}
