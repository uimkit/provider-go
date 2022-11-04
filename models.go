package uim

import (
	"time"
)

// 性别
type Gender int

const (
	GenderUnknown Gender = iota // 未知
	GenderMale                  // 男
	GenderFemale                // 女
)

// IM用户
type IMUser struct {
	UserId          string         `json:"user_id,omitempty"`          // 平台用户ID
	OpenId          string         `json:"open_id,omitempty"`          // 实际的平台用户ID，如：微信ID
	CustomId        string         `json:"custom_id,omitempty"`        // 用户自定义ID
	Username        string         `json:"username,omitempty"`         // 用户账户
	Name            string         `json:"name,omitempty"`             // 名称
	Nickname        string         `json:"nickname,omitempty"`         // 昵称
	RealName        string         `json:"real_name,omitempty"`        // 真实名字
	Mobile          string         `json:"mobile,omitempty"`           // 手机号
	Tel             string         `json:"tel,omitempty"`              // 座机电话
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
	Company         string         `json:"company,omitempty"`          // 公司
	Department      string         `json:"department,omitempty"`       // 部门
	Title           string         `json:"title,omitempty"`            // 头衔、职位
	Language        string         `json:"language,omitempty"`         // 语言
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// IM用户变更
type IMUserUpdate struct {
	UserId          string         `json:"user_id,omitempty"`          // 平台用户ID
	OpenId          *string        `json:"open_id,omitempty"`          // 实际的平台用户ID，如：微信ID
	CustomId        *string        `json:"custom_id,omitempty"`        // 用户自定义ID
	Username        *string        `json:"username,omitempty"`         // 用户账户
	Name            *string        `json:"name,omitempty"`             // 名称
	Nickname        *string        `json:"nickname,omitempty"`         // 昵称
	RealName        *string        `json:"real_name,omitempty"`        // 真实名字
	Mobile          *string        `json:"mobile,omitempty"`           // 手机号
	Tel             *string        `json:"tel,omitempty"`              // 座机电话
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
	Company         *string        `json:"company,omitempty"`          // 公司
	Department      *string        `json:"department,omitempty"`       // 部门
	Title           *string        `json:"title,omitempty"`            // 头衔、职位
	Language        *string        `json:"language,omitempty"`         // 语言
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 账号在线状态
type Presence int

const (
	PresenceInitializing       Presence = iota // 初始化中
	PresenceOnline                             // 在线
	PresenceOffline                            // 离线
	PresenceLogout                             // 登出
	PresenceDisabled                           // 禁用
	PresenceDisabledByProvider                 // 服务商封禁
)

// 账号
type IMAccount struct {
	UserId          string         `json:"user_id,omitempty"`          // 平台用户ID
	OpenId          string         `json:"open_id,omitempty"`          // 实际的平台用户ID，如：微信ID
	CustomId        string         `json:"custom_id,omitempty"`        // 用户自定义ID
	Username        string         `json:"username,omitempty"`         // 用户账户
	Name            string         `json:"name,omitempty"`             // 名称
	Nickname        string         `json:"nickname,omitempty"`         // 昵称
	RealName        string         `json:"real_name,omitempty"`        // 真实名字
	Mobile          string         `json:"mobile,omitempty"`           // 手机号
	Tel             string         `json:"tel,omitempty"`              // 座机电话
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
	Company         string         `json:"company,omitempty"`          // 公司
	Department      string         `json:"department,omitempty"`       // 部门
	Title           string         `json:"title,omitempty"`            // 头衔、职位
	Language        string         `json:"language,omitempty"`         // 语言
	Presence        Presence       `json:"presence,omitempty"`         // 状态
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
	State           string         `json:"state,omitempty"`            // 授权账号时客户传来的自定义数据，透传回去
}

// 账号变更
type IMAccountUpdate struct {
	UserId          string         `json:"user_id,omitempty"`          // 平台用户ID
	OpenId          *string        `json:"open_id,omitempty"`          // 实际的平台用户ID，如：微信ID
	CustomId        *string        `json:"custom_id,omitempty"`        // 用户自定义ID
	Username        *string        `json:"username,omitempty"`         // 用户账户
	Name            *string        `json:"name,omitempty"`             // 名称
	Nickname        *string        `json:"nickname,omitempty"`         // 昵称
	RealName        *string        `json:"real_name,omitempty"`        // 真实名字
	Mobile          *string        `json:"mobile,omitempty"`           // 手机号
	Tel             *string        `json:"tel,omitempty"`              // 座机电话
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
	Company         *string        `json:"company,omitempty"`          // 公司
	Department      *string        `json:"department,omitempty"`       // 部门
	Title           *string        `json:"title,omitempty"`            // 头衔、职位
	Language        *string        `json:"language,omitempty"`         // 语言
	Presence        *Presence      `json:"presence,omitempty"`         // 状态
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 联系人
type Contact struct {
	Account         string         `json:"account,omitempty"`          // 所属账号的平台用户ID
	UserId          string         `json:"user_id,omitempty"`          // 平台用户ID
	OpenId          string         `json:"open_id,omitempty"`          // 实际的平台用户ID，如：微信ID
	CustomId        string         `json:"custom_id,omitempty"`        // 用户自定义ID
	Username        string         `json:"username,omitempty"`         // 用户账户
	Name            string         `json:"name,omitempty"`             // 名称
	Nickname        string         `json:"nickname,omitempty"`         // 昵称
	RealName        string         `json:"real_name,omitempty"`        // 真实名字
	Mobile          string         `json:"mobile,omitempty"`           // 手机号
	Tel             string         `json:"tel,omitempty"`              // 座机电话
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
	Company         string         `json:"company,omitempty"`          // 公司
	Department      string         `json:"department,omitempty"`       // 部门
	Title           string         `json:"title,omitempty"`            // 头衔、职位
	Language        string         `json:"language,omitempty"`         // 语言
	Alias           string         `json:"alias,omitempty"`            // 备注名
	Remark          string         `json:"remark,omitempty"`           // 备注说明
	Blocked         bool           `json:"blocked,omitempty"`          // 是否拉黑
	Marked          bool           `json:"marked,omitempty"`           // 是否星标
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 粉丝
type Follower Contact

// 关注的人
type Following Contact

// 好友申请
type FriendApply struct {
	ID              string         `json:"id,omitempty"`               // 申请ID
	Account         string         `json:"account,omitempty"`          // 接收申请的账号的平台用户ID
	UserId          string         `json:"user_id,omitempty"`          // 申请人平台用户ID，非必须
	OpenId          string         `json:"open_id,omitempty"`          // 申请人实际的平台用户ID，如：微信ID
	CustomId        string         `json:"custom_id,omitempty"`        // 申请人用户自定义ID
	Username        string         `json:"username,omitempty"`         // 申请人用户账户
	Name            string         `json:"name,omitempty"`             // 申请人名称
	Nickname        string         `json:"nickname,omitempty"`         // 申请人昵称
	RealName        string         `json:"real_name,omitempty"`        // 申请人真实名字
	Mobile          string         `json:"mobile,omitempty"`           // 申请人手机号
	Tel             string         `json:"tel,omitempty"`              // 申请人座机电话
	Email           string         `json:"email,omitempty"`            // 申请人邮箱
	Avatar          string         `json:"avatar,omitempty"`           // 申请人头像URL
	QRCode          string         `json:"qrcode,omitempty"`           // 申请人二维码URL
	Gender          Gender         `json:"gender,omitempty"`           // 申请人性别
	Country         string         `json:"country,omitempty"`          // 申请人国家
	Province        string         `json:"province,omitempty"`         // 申请人省份
	City            string         `json:"city,omitempty"`             // 申请人城市
	District        string         `json:"district,omitempty"`         // 申请人区
	Address         string         `json:"address,omitempty"`          // 申请人地址
	Signature       string         `json:"signature,omitempty"`        // 申请人签名
	Birthday        *time.Time     `json:"birthday,omitempty"`         // 申请人生日
	Company         string         `json:"company,omitempty"`          // 申请人公司
	Department      string         `json:"department,omitempty"`       // 申请人部门
	Title           string         `json:"title,omitempty"`            // 申请人头衔、职位
	Language        string         `json:"language,omitempty"`         // 申请人语言
	HelloMessage    string         `json:"hello_message,omitempty"`    // 打招呼留言
	Source          string         `json:"source,omitempty"`           // 添加好友来源
	AppliedAt       *time.Time     `json:"applied_at,omitempty"`       // 申请时间
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 会话类型
type ConversationType string

const (
	ConversationTypePrivate         ConversationType = "private"          // 私聊
	ConversationTypeGroup           ConversationType = "group"            // 群聊
	ConversationTypeDiscussion      ConversationType = "discussion"       // 聊天室/讨论组
	ConversationTypeSystem          ConversationType = "system"           // 系统
	ConversationTypeCustomerService ConversationType = "customer_service" // 客服
)

// 消息@用户类型
type MentionedType int

const (
	MentionedTypeNone  MentionedType = iota // 没有@
	MentionedTypeAll                        // 所有人
	MentionedTypeUsers                      // 指定人
)

// 消息类型
type MessageType string

const (
	MessageTypeText     MessageType = "text"     // 文本消息
	MessageTypeImage    MessageType = "image"    // 图片消息
	MessageTypeVoice    MessageType = "voice"    // 语音消息
	MessageTypeVideo    MessageType = "video"    // 视频消息
	MessageTypeLink     MessageType = "link"     // 链接消息
	MessageTypeLocation MessageType = "location" // 位置消息
)

type ImageMessageBody struct {
	URL    string `json:"url,omitempty"`    // 图片URL
	Width  int    `json:"width,omitempty"`  // 宽度（像素）
	Height int    `json:"height,omitempty"` // 高度（像素）
	Size   int    `json:"size,omitempty"`   // 大小（字节）
	Ext    string `json:"ext,omitempty"`    // 类型，如：png、jpeg
	MD5    string `json:"md5,omitempty"`    // 文件内容MD5
}

type VoiceMessageBody struct {
	URL      string `json:"url,omitempty"`      // 语音URL
	Duration int    `json:"duration,omitempty"` // 时长（毫秒）
	Size     int    `json:"size,omitempty"`     // 大小（字节）
	Ext      string `json:"ext,omitempty"`      // 类型，如：mp3
	MD5      string `json:"md5,omitempty"`      // 文件内容MD5
}

type VideoMessageBody struct {
	URL      string `json:"url,omitempty"`      // 视频URL
	Duration int    `json:"duration,omitempty"` // 时长（毫秒）
	Width    int    `json:"width,omitempty"`    // 宽度（像素）
	Height   int    `json:"height,omitempty"`   // 高度（像素）
	Size     int    `json:"size,omitempty"`     // 大小（字节）
	Ext      string `json:"ext,omitempty"`      // 类型，如：mp4
	MD5      string `json:"md5,omitempty"`      // 文件内容MD5
}

// 消息参与者
type MessageParticipant struct {
	UserId string `json:"id,omitempty"`     // 平台用户ID
	Name   string `json:"name,omitempty"`   // 名称
	Avatar string `json:"avatar,omitempty"` // 头像
}

// 消息
type Message struct {
	MessageId        string                `json:"message_id,omitempty"`        // 平台消息ID
	Channel          string                `json:"channel,omitempty"`           // 消息收发地址，账号回复消息时会发送到此地址
	ConversationType ConversationType      `json:"conversation_type,omitempty"` // 会话类型
	Account          string                `json:"account,omitempty"`           // 归属账号的平台用户ID
	UserId           string                `json:"user_id,omitempty"`           // 消息发送人平台用户ID
	Type             MessageType           `json:"type,omitempty"`              // 消息类型
	Text             string                `json:"text,omitempty"`              // 文本消息
	Image            *ImageMessageBody     `json:"image,omitempty"`             // 图片消息、视频消息封面
	Thumb            *ImageMessageBody     `json:"thumb,omitempty"`             // 图片消息、视频消息封面缩略图
	Voice            *VoiceMessageBody     `json:"voice,omitempty"`             // 语音消息
	Video            *VideoMessageBody     `json:"video,omitempty"`             // 视频消息
	MentionedType    MentionedType         `json:"mentioned_type,omitempty"`    // @用户类型
	MentionedUsers   []*MessageParticipant `json:"mentioned_users"`             // @用户列表
	SentAt           *time.Time            `json:"sent_at,omitempty"`           // 发送时间
	Revoked          bool                  `json:"revoked,omitempty"`           // 是否撤回
	Metadata         map[string]any        `json:"metadata,omitempty"`          // 公开元数据
	PrivateMetadata  map[string]any        `json:"private_metadata,omitempty"`  // 私有元数据
	State            string                `json:"state,omitempty"`             // 发送消息时携带的业务自定义数据，发送后返回消息会透传给业务方
}

// 添加好友
type AddContactRequest struct {
	UserId       string `json:"user_id,omitempty"`       // 账号的平台用户ID
	Contact      string `json:"contact,omitempty"`       // 添加的好友，可以为手机号、平台ID
	HelloMessage string `json:"hello_message,omitempty"` // 打招呼消息
}

// 添加好友返回
type AddContactResponse struct {
	BaseResponse
	Success bool   `json:"success"` // 是否发起好友申请成功
	Reason  string `json:"reason"`  // 发起申请好友失败原因
}

// 群组
type Group struct {
	UserId          string         `json:"user_id,omitempty"`          // 归属的平台用户ID
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	Owner           *IMUser        `json:"owner,omitempty"`            // 群主信息
	Name            string         `json:"name,omitempty"`             // 名称
	Avatar          string         `json:"avatar,omitempty"`           // 头像URL
	Announcement    string         `json:"announcement,omitempty"`     // 群公告
	Description     string         `json:"description,omitempty"`      // 群介绍
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 群组变更
type GroupUpdate struct {
	UserId          string         `json:"user_id,omitempty"`          // 归属的平台用户ID
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	Owner           *IMUser        `json:"owner,omitempty"`            // 群主变更
	Name            *string        `json:"name,omitempty"`             // 名称
	Avatar          *string        `json:"avatar,omitempty"`           // 头像URL
	Announcement    *string        `json:"announcement,omitempty"`     // 群公告
	Description     *string        `json:"description,omitempty"`      // 群介绍
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 群组删除
type GroupDelete struct {
	UserId  string `json:"user_id,omitempty"`  // 归属的平台用户ID
	GroupId string `json:"group_id,omitempty"` // 平台群组ID
}

// 入群邀请
type GroupInvitation struct {
	ID              string         `json:"id,omitempty"`               // 入群邀请ID
	UserId          string         `json:"user_id,omitempty"`          // 收到邀请的平台用户ID
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	Inviter         *IMUser        `json:"inviter,omitempty"`          // 邀请人信息
	HelloMessage    string         `json:"hello_message,omitempty"`    // 打招呼留言
	InvitedAt       *time.Time     `json:"invited_at,omitempty"`       // 邀请时间
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 入群申请
type JoinGroupApply struct {
	ID              string         `json:"id,omitempty"`               // 入群申请ID
	UserId          string         `json:"user_id,omitempty"`          // 收到申请的平台用户ID
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	ApplyUser       *IMUser        `json:"apply_user,omitempty"`       // 申请用户信息
	HelloMessage    string         `json:"hello_message,omitempty"`    // 打招呼留言
	AppliedAt       *time.Time     `json:"applied_at,omitempty"`       // 申请时间
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 群组成员
type GroupMember struct {
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	MemberId        string         `json:"member_id,omitempty"`        // 平台群成员ID
	User            *IMUser        `json:"user,omitempty"`             // 群成员的用户信息
	IsOwner         bool           `json:"is_owner,omitempty"`         // 是否群主
	IsAdmin         bool           `json:"is_admin,omitempty"`         // 是否管理员
	Alias           string         `json:"alias,omitempty"`            // 群内备注名
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 群组成员变更
type GroupMemberUpdate struct {
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	MemberId        string         `json:"member_id,omitempty"`        // 平台群成员ID
	User            *IMUserUpdate  `json:"user,omitempty"`             // 群成员的用户信息
	IsOwner         *bool          `json:"is_owner,omitempty"`         // 是否群主
	IsAdmin         *bool          `json:"is_admin,omitempty"`         // 是否管理员
	Alias           *string        `json:"alias,omitempty"`            // 群内备注名
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 群组成员删除
type GroupMemberDelete struct {
	GroupId  string `json:"group_id,omitempty"`  // 平台群组ID
	MemberId string `json:"member_id,omitempty"` // 平台群成员ID
}

// 会话的另一方
type ConversationParty struct {
	PartyId string `json:"party_id,omitempty"` // 另一方平台ID
	Name    string `json:"name,omitempty"`     // 名称
	Avatar  string `json:"avatar,omitempty"`   // 头像
}

// 会话
type Conversation struct {
	ConversationId  string             `json:"conversation_id,omitempty"`  // 平台会话ID
	UserId          string             `json:"user_id,omitempty"`          // 归属用户平台ID
	Type            ConversationType   `json:"type,omitempty"`             // 会话类型
	Party           *ConversationParty `json:"party,omitempty"`            // 会话另一方
	Metadata        map[string]any     `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any     `json:"private_metadata,omitempty"` // 私有元数据
}

// 会话更新
type ConversationUpdate struct {
	UserId          string           `json:"user_id,omitempty"`          // 归属用户平台ID
	Type            ConversationType `json:"type,omitempty"`             // 会话类型
	PartyId         string           `json:"party_id,omitempty"`         // 另一方平台ID
	Metadata        map[string]any   `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any   `json:"private_metadata,omitempty"` // 私有元数据
}

// 消息变更
type MessageUpdate struct {
	MessageId       string         `json:"message_id,omitempty"`       // 平台消息ID
	Revoked         *bool          `json:"revoked,omitempty"`          // 是否撤回
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
	State           string         `json:"state,omitempty"`            // 发送消息时携带的业务自定义数据，发送后返回消息会透传给业务方
}

// 发送消息
type SendMessageRequest struct {
	From             string           `json:"from,omitempty"`              // 消息发送者
	To               string           `json:"to,omitempty"`                // 消息接受者
	ConversationType ConversationType `json:"conversation_type,omitempty"` // 所属的会话类型
	Seq              int              `json:"seq,omitempty"`               // 序列号，在会话中唯一且有序增长，用于确保消息顺序
	MentionedType    MentionedType    `json:"mentioned_type,omitempty"`    // @用户类型
	MentionedUserIds []string         `json:"mentioned_users"`             // @用户列表
	SentAt           *time.Time       `json:"sent_at,omitempty"`           // 发送时间
	Payload          *MessagePayload  `json:"payload,omitempty"`           // 消息内容
}

// 发送消息结果
type SendMessageResponse struct {
	BaseResponse
	Message
}

// 消息内容
type MessagePayload struct {
	Type MessageType `json:"type,omitempty"` // 消息类型
	Body any         `json:"body,omitempty"` // 消息体
}

// 自定义数据值类型
type MetafieldValueType string

const (
	MetafieldValueTypeInteger   MetafieldValueType = "integer"
	MetafieldValueTypeString    MetafieldValueType = "string"
	MetafieldValueTypeBoolean   MetafieldValueType = "boolean"
	MetafieldValueTypeDateTime  MetafieldValueType = "datetime"
	MetafieldValueTypeJsonArray MetafieldValueType = "json_array"
	MetafieldValueTypeJsonMap   MetafieldValueType = "json_map"
	MetafieldValueTypeDecimal   MetafieldValueType = "decimal"
)

// 自定义数据
type Metafield struct {
	Resource   string             `json:"resource,omitempty"`    // 归属资源类型
	ResourceId string             `json:"resource_id,omitempty"` // 归属资源ID
	Namespace  string             `json:"namespace,omitempty"`   // 命令空间
	Key        string             `json:"key,omitempty"`         // 字段名
	Type       MetafieldValueType `json:"type,omitempty"`        // 字段类型
	Value      any                `json:"value,omitempty"`       // 字段值
}

// 自定义数据变更
type MetafieldUpdate struct {
	Resource   string             `json:"resource,omitempty"`    // 归属资源类型
	ResourceId string             `json:"resource_id,omitempty"` // 归属资源ID
	Namespace  string             `json:"namespace,omitempty"`   // 命令空间
	Key        string             `json:"key,omitempty"`         // 字段名
	Type       MetafieldValueType `json:"type,omitempty"`        // 字段类型
	Value      any                `json:"value,omitempty"`       // 字段值
}

// 查询自定义数据请求
type GetMetafieldRequest struct {
	Resource   string `json:"resource,omitempty"`    // 归属资源类型
	ResourceId string `json:"resource_id,omitempty"` // 归属资源ID
	Namespace  string `json:"namespace,omitempty"`   // 命令空间
	Key        string `json:"key,omitempty"`         // 字段名
}

// 查询自定义数据结果
type GetMetafieldResponse struct {
	BaseResponse
	Metafield
}
