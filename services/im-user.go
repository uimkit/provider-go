package services

import (
	"time"
)

// IM用户
type IMUser struct {
	ID              string         `json:"id,omitempty"`               // 用户唯一ID，规则：provider|user_id，如：wechat|wx_id
	ConnectionId    string         `json:"connection_id,omitempty"`    // 默认的服务连接
	UserId          string         `json:"user_id,omitempty"`          // 平台用户ID，如：微信ID
	CustomId        string         `json:"custom_id,omitempty"`        // 用户自定义ID，如：微信号
	Username        string         `json:"username,omitempty"`         // 用户账户
	Name            string         `json:"name,omitempty"`             // 名称
	Mobile          string         `json:"mobile,omitempty"`           // 手机号
	Email           string         `json:"email,omitempty"`            // 邮箱
	Avatar          string         `json:"avatar,omitempty"`           // 头像URL
	QRCode          string         `json:"qrcode,omitempty"`           // 二维码URL
	Gender          Gender         `json:"gender,omitempty"`           // 性别
	Country         string         `json:"country,omitempty"`          // 国家
	Province        string         `json:"province,omitempty"`         // 省份
	City            string         `json:"city,omitempty"`             // 城市
	District        string         `json:"district,omitempty"`         // 区
	Address         string         `json:"address,omitempty"`          // 地址
	Signature       string         `json:"signature,omitempty"`        // 签名
	Birthday        *time.Time     `json:"birthday,omitempty"`         // 生日
	Language        string         `json:"language,omitempty"`         // 语言
	Identities      []*IMIdentity  `json:"identities,omitempty"`       // 用户在各服务连接的身份信息
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
	CreatedAt       *time.Time     `json:"created_at,omitempty"`       // 创建时间
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`       // 最后更新时间
}

// 创建IM用户
type CreateIMUser struct {
	ConnectionId    string         `json:"connection_id,omitempty"`    // 默认的服务连接
	UserId          string         `json:"user_id,omitempty"`          // 平台用户ID，如：微信ID
	CustomId        string         `json:"custom_id,omitempty"`        // 用户自定义ID，如：微信号
	Username        string         `json:"username,omitempty"`         // 用户账户
	Name            string         `json:"name,omitempty"`             // 名称
	Mobile          string         `json:"mobile,omitempty"`           // 手机号
	Email           string         `json:"email,omitempty"`            // 邮箱
	Avatar          string         `json:"avatar,omitempty"`           // 头像URL
	QRCode          string         `json:"qrcode,omitempty"`           // 二维码URL
	Gender          Gender         `json:"gender,omitempty"`           // 性别
	Country         string         `json:"country,omitempty"`          // 国家
	Province        string         `json:"province,omitempty"`         // 省份
	City            string         `json:"city,omitempty"`             // 城市
	District        string         `json:"district,omitempty"`         // 区
	Address         string         `json:"address,omitempty"`          // 地址
	Signature       string         `json:"signature,omitempty"`        // 签名
	Birthday        *time.Time     `json:"birthday,omitempty"`         // 生日
	Language        string         `json:"language,omitempty"`         // 语言
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 修改用户信息
type UpdateIMUser struct {
	ID              string         `json:"id,omitempty"`               // 用户唯一ID，规则：provider|user_id，如：wechat|wx_id
	Username        *string        `json:"username,omitempty"`         // 用户账户
	Name            *string        `json:"name,omitempty"`             // 名称
	Mobile          *string        `json:"mobile,omitempty"`           // 手机号
	Email           *string        `json:"email,omitempty"`            // 邮箱
	Avatar          *string        `json:"avatar,omitempty"`           // 头像URL
	QRCode          *string        `json:"qrcode,omitempty"`           // 二维码URL
	Gender          *Gender        `json:"gender,omitempty"`           // 性别
	Country         *string        `json:"country,omitempty"`          // 国家
	Province        *string        `json:"province,omitempty"`         // 省份
	City            *string        `json:"city,omitempty"`             // 城市
	District        *string        `json:"district,omitempty"`         // 区
	Address         *string        `json:"address,omitempty"`          // 地址
	Signature       *string        `json:"signature,omitempty"`        // 签名
	Birthday        *time.Time     `json:"birthday,omitempty"`         // 生日
	Language        *string        `json:"language,omitempty"`         // 语言
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// IM身份信息。同一个IMUser，在每一个IMConnection，都会有一个对应的IMIdentity。
type IMIdentity struct {
	ID           string     `json:"id,omitempty"`            // 身份ID
	UserId       string     `json:"user_id,omitempty"`       // 平台用户ID，用于识别身份是否属于同一个IMUser，如：微信ID
	OpenId       string     `json:"open_id,omitempty"`       // IMConnection分配的用户ID
	Provider     string     `json:"provider,omitempty"`      // 服务提供者，如：微信
	ConnectionId string     `json:"connection_id,omitempty"` // 身份属于的服务连接
	CreatedAt    *time.Time `json:"created_at,omitempty"`    // 创建时间
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`    // 最后更新时间
}

type CreateIMIdentity struct {
	UserId       string `json:"user_id,omitempty"`       // 平台用户ID，用于识别身份是否属于同一个IMUser，如：微信ID
	OpenId       string `json:"open_id,omitempty"`       // 服务连接分配的用户ID
	Provider     string `json:"provider,omitempty"`      // 服务提供者，如：微信
	ConnectionId string `json:"connection_id,omitempty"` // 身份属于的服务连接
}

type UpdateIMIdentity struct {
	ID           string  `json:"id,omitempty"`            // 身份ID
	UserId       *string `json:"user_id,omitempty"`       // 平台用户ID，用于识别身份是否属于同一个IMUser，如：微信ID
	OpenId       *string `json:"open_id,omitempty"`       // 服务连接分配的用户ID
	Provider     *string `json:"provider,omitempty"`      // 服务提供者，如：微信
	ConnectionId *string `json:"connection_id,omitempty"` // 身份属于的服务连接
}
