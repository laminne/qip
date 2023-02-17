package router

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

var unAuthorizedErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "UnAuthorized",
	Message: "認証情報がありません",
}
var failedValidationErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "FailedValidation",
	Message: "バリデーションに失敗しました",
}
var notFoundErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "NotFound",
	Message: "エンドポイントが存在しません",
}
var internalErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "InternalError",
	Message: "内部エラーが発生しました",
}

var postTooLongErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "PostTooLong",
	Message: "投稿が長すぎます",
}
var invalidTargetErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "InvalidTarget",
	Message: "ターゲット先が正しく指定されていません",
}
var alreadyWatchedErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "AlreadyWatched",
	Message: "すでにウォッチしています",
}
