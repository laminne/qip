package types

import "time"

type PostJSONLD struct {
	ID        string      `json:"id"`    // /notes/:id/activity
	Actor     string      `json:"actor"` // /users/:id
	Type      string      `json:"type"`  // 本人: Create / RN: Announce
	Published time.Time   `json:"published"`
	Object    interface{} `json:"object"` // NoteObject or string
	To        []string    `json:"to"`     // 公開: https://www.w3.org/ns/activitystreams#Public
	Cc        []string    `json:"cc"`     // /users/:id/followers
}

type PostObject struct {
	ID             string `json:"id"`               // /notes/:id/activity
	Type           string `json:"type"`             // Note
	AttributedTo   string `json:"attributedTo"`     // /users/:id
	Content        string `json:"content"`          // 本文 (HTMLになる
	MisskeyContent string `json:"_misskey_content"` // mfmがそのまま入る
	Source         struct {
		Content   string `json:"content"`   // misskeyContentと一緒
		MediaType string `json:"mediaType"` // text/x.misskeymarkdown <- mfm
	} `json:"source"`
	Published  time.Time `json:"published"`
	To         []string  `json:"to"` // NoteJSONLDと一緒？
	Cc         []string  `json:"cc"`
	InReplyTo  *string   `json:"inReplyTo"` // リプライ先のURL /notes/:id
	Attachment []struct {
		Type      string  `json:"type"`      // Document
		MediaType string  `json:"mediaType"` // mimeタイプ
		Url       string  `json:"url"`       // ファイルへのリンク
		Name      *string `json:"name"`
	} `json:"attachment"`
	Sensitive bool          `json:"sensitive"` // CWかどうか？
	Tag       []interface{} `json:"tag"`       // 多分ハッシュタグ
}
