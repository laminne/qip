package domain

import (
	"time"

	"github.com/approvers/qip/pkg/utils/id"
)

// File ファイル
type File struct {
	ID        id.SnowFlakeID
	UserID    id.SnowFlakeID // 投稿したユーザーのID
	Host      string         // 投稿したユーザーのホスト
	MD5Hash   string         // ファイルのMD5ハッシュ値
	MimeType  string         // mimeタイプ
	FileSize  uint           // ファイルサイズ
	URL       string
	IsNSFW    bool   // 閲覧注意フラグがついているか
	BlurHash  string // ブラーハッシュ
	CreatedAt time.Time
}
