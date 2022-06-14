package services

import (
	"time"
)

// 用户动态
type Moment struct {
	ID              string         `json:"id,omitempty"`               // 动态唯一ID
	MomentId        string         `json:"moment_id,omitempty"`        // 平台动态ID
	User            *IMUser        `json:"user,omitempty"`             // 发布人用户信息
	Content         *MomentContent `json:"content,omitempty"`          // 动态内容
	PublishedAt     *time.Time     `json:"published_at,omitempty"`     // 发布时间
	IsPrivate       bool           `json:"is_private,omitempty"`       // 是否私密动态
	Tags            []string       `json:"tags,omitempty"`             // 标签
	Location        *Location      `json:"location,omitempty"`         // 位置信息
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
	CreatedAt       *time.Time     `json:"created_at,omitempty"`       // 创建时间
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`       // 最后更新时间
}

// 发布动态
type PostMoment struct {
	MomentId        string         `json:"moment_id,omitempty"`        // 平台动态ID
	UserId          string         `json:"user_id,omitempty"`          // 发布用户ID
	Content         *MomentContent `json:"content,omitempty"`          // 动态内容
	PublishedAt     *time.Time     `json:"published_at,omitempty"`     // 发布时间
	IsPrivate       bool           `json:"is_private,omitempty"`       // 是否私密动态
	Tags            []string       `json:"tags,omitempty"`             // 标签
	Location        *Location      `json:"location,omitempty"`         // 位置信息
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 更新动态
type UpdateMoment struct {
	ID              string         `json:"id,omitempty"`               // 动态唯一ID
	Content         *MomentContent `json:"content,omitempty"`          // 动态内容
	IsPrivate       *bool          `json:"is_private,omitempty"`       // 是否私密动态
	Tags            []string       `json:"tags,omitempty"`             // 标签
	Location        *Location      `json:"location,omitempty"`         // 位置信息
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 动态内容
type MomentContent struct {
	Type MomentType `json:"type,omitempty"` // 类型
	Body any        `json:"body,omitempty"` // 内容
}

// 文本动态
type TextMomentBody struct {
	Content string `json:"content,omitempty"` // 文本
}

// 图文动态
type ThumbMomentBody struct {
	URL    string `json:"url,omitempty"`    // 图片URL
	Width  int    `json:"width,omitempty"`  // 宽度（像素）
	Height int    `json:"height,omitempty"` // 高度（像素）
	Ext    string `json:"ext,omitempty"`    // 类型，如：png、jpeg
}

type ImageMomentBody struct {
	URL    string           `json:"url,omitempty"`    // 图片URL
	Width  int              `json:"width,omitempty"`  // 宽度（像素）
	Height int              `json:"height,omitempty"` // 高度（像素）
	Size   int              `json:"size,omitempty"`   // 大小（字节）
	Ext    string           `json:"ext,omitempty"`    // 类型，如：png、jpeg
	MD5    string           `json:"md5,omitempty"`    // 文件内容MD5
	Thumb  *ThumbMomentBody `json:"thumb,omitempty"`  // 缩略图
}

type ImageListMomentBody struct {
	Images []*ImageMomentBody `json:"images,omitempty"` // 图片
	Text   string             `json:"text,omitempty"`   // 文本
}

// 视频动态
type VideoMomentBody struct {
	Text     string           `json:"text,omitempty"`     // 文本
	URL      string           `json:"url,omitempty"`      // 视频URL
	Duration int              `json:"duration,omitempty"` // 时长（毫秒）
	Width    int              `json:"width,omitempty"`    // 宽度（像素）
	Height   int              `json:"height,omitempty"`   // 高度（像素）
	Size     int              `json:"size,omitempty"`     // 大小（字节）
	Ext      string           `json:"ext,omitempty"`      // 类型，如：mp4
	MD5      string           `json:"md5,omitempty"`      // 文件内容MD5
	Thumb    *ThumbMomentBody `json:"thumb,omitempty"`    // 缩略图
}

// 位置信息
type Location struct {
	Latitude   float64 `json:"latitude,omitempty"`    // 纬度
	Longitude  float64 `json:"longitude,omitempty"`   // 经度
	Altitude   float64 `json:"altitude,omitempty"`    // 海拔
	Accuracy   float64 `json:"accuracy,omitempty"`    // 精度
	City       string  `json:"city,omitempty"`        // 城市
	PlaceName  string  `json:"place_name,omitempty"`  // 位置名称
	PoiName    string  `json:"poi_name,omitempty"`    // 位置描述
	PoiAddress string  `json:"poi_address,omitempty"` // 位置详细地址
}
