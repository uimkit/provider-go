package uim

import "time"

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

// IM账号
type IMAccount struct {
	User            *IMUser        `json:"user,omitempty"`             // 用户信息
	Presence        Presence       `json:"presence,omitempty"`         // 状态
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 账号变更
type IMAccountUpdate struct {
	User            *IMUserUpdate  `json:"user,omitempty"`             // 用户信息变更
	Presence        *Presence      `json:"presence,omitempty"`         // 状态
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 新账号
type NewIMAccount struct {
	Account *IMAccount `json:"account,omitempty"` // 账号
	State   string     `json:"state,omitempty"`   // 授权账号时客户传来的自定义数据，透传回去
}

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
	CustomId        string         `json:"custom_id,omitempty"`        // 用户自定义ID
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

// IM用户变更
type IMUserUpdate struct {
	UserId          string         `json:"user_id,omitempty"`          // 平台用户ID
	CustomId        *string        `json:"custom_id,omitempty"`        // 用户自定义ID
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

// 好友申请
type FriendApply struct {
	ID              string         `json:"id,omitempty"`               // 申请ID
	UserId          string         `json:"user_id,omitempty"`          // 接收申请的平台用户ID
	ApplyUser       *IMUser        `json:"apply_user,omitempty"`       // 申请人信息
	HelloMessage    string         `json:"hello_message,omitempty"`    // 打招呼留言
	AppliedAt       *time.Time     `json:"applied_at,omitempty"`       // 申请时间
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 联系人
type Contact struct {
	UserId          string         `json:"user_id,omitempty"`          // 归属的平台用户ID
	ContactUser     *IMUser        `json:"contact_user,omitempty"`     // 联系人的用户信息
	Alias           string         `json:"alias,omitempty"`            // 备注名
	Remark          string         `json:"remark,omitempty"`           // 备注说明
	Tags            []string       `json:"tags,omitempty"`             // 标签
	Blocked         bool           `json:"blocked,omitempty"`          // 是否拉黑
	Marked          bool           `json:"marked,omitempty"`           // 是否星标
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 联系人变更
type ContactUpdate struct {
	UserId          string         `json:"user_id,omitempty"`          // 归属的平台用户ID
	ContactUser     *IMUserUpdate  `json:"contact_user,omitempty"`     // 联系人的用户信息
	Alias           *string        `json:"alias,omitempty"`            // 备注名
	Remark          *string        `json:"remark,omitempty"`           // 备注说明
	Tags            []string       `json:"tags,omitempty"`             // 标签
	Blocked         *bool          `json:"blocked,omitempty"`          // 是否拉黑
	Marked          *bool          `json:"marked,omitempty"`           // 是否星标
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 联系人删除
type ContactDeleted struct {
	UserId        string `json:"user_id,omitempty"`         // 归属的平台用户ID
	ContactUserId string `json:"contact_user_id,omitempty"` // 联系人的用户ID
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

// 会话类型
type ConversationType int

const (
	ConversationTypePrivate         ConversationType = iota + 1 // 私聊
	ConversationTypeGroup                                       // 群聊
	ConversationTypeDiscussion                                  // 聊天室/讨论组
	ConversationTypeSystem                                      // 系统
	ConversationTypeCustomerService                             // 客服
)

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

// 消息@用户类型
type MentionedType int

const (
	MentionedTypeNone     MentionedType = iota // 没有@
	MentionedTypeAll                           // 所有人
	MentionedTypeSpecific                      // 指定人
)

// 消息参与方
type MessageParticipant struct {
	ID     string `json:"id,omitempty"`     // 平台ID
	Name   string `json:"name,omitempty"`   // 名称
	Avatar string `json:"avatar,omitempty"` // 头像
}

// 消息
type Message struct {
	MessageId        string              `json:"message_id,omitempty"`        // 平台消息ID
	UserId           string              `json:"user_id,omitempty"`           // 归属用户平台ID
	From             *MessageParticipant `json:"from,omitempty"`              // 消息发送者
	To               *MessageParticipant `json:"to,omitempty"`                // 消息接受者
	ConversationType ConversationType    `json:"conversation_type,omitempty"` // 会话类型
	MentionedType    MentionedType       `json:"mentioned_type,omitempty"`    // @用户类型
	MentionedUsers   []*IMUser           `json:"mentioned_users"`             // @用户信息
	SentAt           *time.Time          `json:"sent_at,omitempty"`           // 发送时间
	Payload          *MessagePayload     `json:"payload,omitempty"`           // 消息内容
	Revoked          bool                `json:"revoked,omitempty"`           // 是否撤回
	Metadata         map[string]any      `json:"metadata,omitempty"`          // 公开元数据
	PrivateMetadata  map[string]any      `json:"private_metadata,omitempty"`  // 私有元数据
	State            string              `json:"state,omitempty"`             // 发送消息时携带的业务自定义数据，发送后返回消息会透传给业务方
}

// 消息变更
type MessageUpdate struct {
	MessageId       string         `json:"message_id,omitempty"`       // 平台消息ID
	Revoked         *bool          `json:"revoked,omitempty"`          // 是否撤回
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
	State           string         `json:"state,omitempty"`            // 发送消息时携带的业务自定义数据，发送后返回消息会透传给业务方
}

// 消息类型
type MessageType int

const (
	MessageTypeUndefined MessageType = iota // 未定义消息
	MessageTypeText                         // 文本消息
	MessageTypeImage                        // 图片消息
	MessageTypeVoice                        // 语音消息
	MessageTypeVideo                        // 视频消息
	MessageTypeLink                         // 链接消息
	MessageTypeLocation                     // 位置消息
)

// 消息内容
type MessagePayload struct {
	Type MessageType `json:"type,omitempty"` // 消息类型
	Body any         `json:"body,omitempty"` // 消息体
}

// 文本消息内容
type TextMessageBody struct {
	Content string `json:"content,omitempty"` // 文本内容
}

// 图片消息内容
type ThumbMessageBody struct {
	URL    string `json:"url,omitempty"`    // 图片URL
	Width  int    `json:"width,omitempty"`  // 宽度（像素）
	Height int    `json:"height,omitempty"` // 高度（像素）
	Ext    string `json:"ext,omitempty"`    // 类型，如：png、jpeg
}

type ImageMessageBody struct {
	URL    string            `json:"url,omitempty"`    // 图片URL
	Width  int               `json:"width,omitempty"`  // 宽度（像素）
	Height int               `json:"height,omitempty"` // 高度（像素）
	Size   int               `json:"size,omitempty"`   // 大小（字节）
	Ext    string            `json:"ext,omitempty"`    // 类型，如：png、jpeg
	MD5    string            `json:"md5,omitempty"`    // 文件内容MD5
	Thumb  *ThumbMessageBody `json:"thumb,omitempty"`  // 缩略图
}

// 语音消息内容
type VoiceMessageBody struct {
	URL      string `json:"url,omitempty"`      // 语音URL
	Duration int    `json:"duration,omitempty"` // 时长（毫秒）
	Size     int    `json:"size,omitempty"`     // 大小（字节）
	Ext      string `json:"ext,omitempty"`      // 类型，如：mp3
	MD5      string `json:"md5,omitempty"`      // 文件内容MD5
}

// 视频消息内容
type VideoMessageBody struct {
	URL      string            `json:"url,omitempty"`      // 视频URL
	Duration int               `json:"duration,omitempty"` // 时长（毫秒）
	Width    int               `json:"width,omitempty"`    // 宽度（像素）
	Height   int               `json:"height,omitempty"`   // 高度（像素）
	Size     int               `json:"size,omitempty"`     // 大小（字节）
	Ext      string            `json:"ext,omitempty"`      // 类型，如：mp4
	MD5      string            `json:"md5,omitempty"`      // 文件内容MD5
	Thumb    *ThumbMessageBody `json:"thumb,omitempty"`    // 缩略图
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

// ============ Deprecated

// 更新账号用户资料
type UpdateIMUser struct {
	UserId    string     `json:"user_id,omitempty"`   // 平台用户ID，如：微信ID
	Username  *string    `json:"username,omitempty"`  // 用户账户
	Name      *string    `json:"name,omitempty"`      // 名称
	Mobile    *string    `json:"mobile,omitempty"`    // 手机号
	Email     *string    `json:"email,omitempty"`     // 邮箱
	Avatar    *string    `json:"avatar,omitempty"`    // 头像URL
	QRCode    *string    `json:"qrcode,omitempty"`    // 二维码URL
	Gender    *Gender    `json:"gender,omitempty"`    // 性别
	Country   *string    `json:"country,omitempty"`   // 国家
	Province  *string    `json:"province,omitempty"`  // 省份
	City      *string    `json:"city,omitempty"`      // 城市
	District  *string    `json:"district,omitempty"`  // 区
	Address   *string    `json:"address,omitempty"`   // 地址
	Signature *string    `json:"signature,omitempty"` // 签名
	Birthday  *time.Time `json:"birthday,omitempty"`  // 生日
	Language  *string    `json:"language,omitempty"`  // 语言
}

// 查询账号列表
type ListIMAccounts struct {
	QueryId string `json:"query_id,omitempty"` // 查询ID，用于匹配异步请求
	Limit   int    `json:"limit,omitempty"`    // 查询数量
	After   string `json:"after,omitempty"`    // 上次查询返回 next，重新开始查询为空
}

// 账号列表
type IMAccountList struct {
	QueryId string       `json:"query_id,omitempty"` // 请求的查询ID，用于匹配异步请求
	Data    []*IMAccount `json:"data,omitempty"`     // 数据列表
	Next    string       `json:"next,omitempty"`     // 游标，用于下次查询请求的 after 参数
}

// 申请添加好友
type NewFriendApply struct {
	UserId       string   `json:"user_id,omitempty"`       // 发起申请的平台用户ID，如：微信ID
	Contacts     []string `json:"contacts,omitempty"`      // 好友列表
	HelloMessage string   `json:"hello_message,omitempty"` // 打招呼留言
}

// 通过好友申请
type ApproveFriendApply struct {
	UserId  string `json:"user_id,omitempty"`  // 操作账号的平台用户ID，如：微信ID
	ApplyId string `json:"apply_id,omitempty"` // 申请ID
}

// 更新联系人
type UpdateContact struct {
	UserId        string   `json:"user_id,omitempty"`         // 操作账号的平台用户ID，如：微信ID
	ContactUserId string   `json:"contact_user_id,omitempty"` // 联系人的平台用户ID
	Alias         *string  `json:"alias,omitempty"`           // 备注名
	Remark        *string  `json:"remark,omitempty"`          // 备注说明
	Tags          []string `json:"tags,omitempty"`            // 标签
	Blocked       *bool    `json:"blocked,omitempty"`         // 是否拉黑
	Marked        *bool    `json:"marked,omitempty"`          // 是否星标
}

// 查询联系人列表
type ListContacts struct {
	QueryId string `json:"query_id,omitempty"` // 查询ID，用于匹配异步请求
	UserId  string `json:"user_id,omitempty"`  // 联系人归属账号的平台用户ID，如：微信ID
	Limit   int    `json:"limit,omitempty"`    // 查询数量
	After   string `json:"after,omitempty"`    // 上次查询返回 next，重新开始查询为空
}

// 联系人列表
type ContactList struct {
	QueryId string     `json:"query_id,omitempty"` // 请求的查询ID，用于匹配异步请求
	Data    []*Contact `json:"data,omitempty"`     // 数据列表
	Next    string     `json:"next,omitempty"`     // 游标，用于下次查询请求的 after 参数
}

// 创建群组
type NewGroup struct {
	UserId       string   `json:"user_id,omitempty"`      // 操作账号的平台用户ID，如：微信ID
	Name         string   `json:"name,omitempty"`         // 名称
	Avatar       string   `json:"avatar,omitempty"`       // 头像URL
	Announcement string   `json:"announcement,omitempty"` // 群公告
	Description  string   `json:"description,omitempty"`  // 群介绍
	UserIds      []string `json:"user_ids,omitempty"`     // 初始邀请入群的用户列表
}

// 修改群组资料
type UpdateGroup struct {
	UserId       string  `json:"user_id,omitempty"`      // 操作账号的平台用户ID，如：微信ID
	GroupId      string  `json:"group_id,omitempty"`     // 平台群组ID
	Name         *string `json:"name,omitempty"`         // 名称
	Avatar       *string `json:"avatar,omitempty"`       // 头像URL
	Announcement *string `json:"announcement,omitempty"` // 群公告
	Description  *string `json:"description,omitempty"`  // 群介绍
}

// 查询群组列表
type ListGroups struct {
	QueryId string `json:"query_id,omitempty"` // 查询ID，用于匹配异步请求
	UserId  string `json:"user_id,omitempty"`  // 群组归属账号的平台用户ID，如：微信ID
	Limit   int    `json:"limit,omitempty"`    // 查询数量
	After   string `json:"after,omitempty"`    // 上次查询返回 next，重新开始查询为空
}

// 群组列表
type GroupList struct {
	QueryId string   `json:"query_id,omitempty"` // 请求的查询ID，用于匹配异步请求
	Data    []*Group `json:"data,omitempty"`     // 数据列表
	Next    string   `json:"next,omitempty"`     // 游标，用于下次查询请求的 after 参数
}

// 邀请入群
type InviteToGroup struct {
	UserId        string   `json:"user_id,omitempty"`         // 操作账号的平台用户ID，如：微信ID
	InviteUserIds []string `json:"invite_user_ids,omitempty"` // 邀请入群的用户列表
	GroupId       string   `json:"group_id,omitempty"`        // 平台群组ID
	HelloMessage  string   `json:"hello_message,omitempty"`   // 打招呼留言
}

// 接受入群邀请
type AcceptGroupInvitation struct {
	UserId       string `json:"user_id,omitempty"`       // 操作账号的平台用户ID，如：微信ID
	GroupId      string `json:"group_id,omitempty"`      // 平台群组ID
	InvitationId string `json:"invitation_id,omitempty"` // 邀请ID
}

// 申请加入群组
type NewJoinGroupApply struct {
	UserId       string `json:"user_id,omitempty"`       // 操作账号的平台用户ID，如：微信ID
	GroupId      string `json:"group_id,omitempty"`      // 平台群组ID
	HelloMessage string `json:"hello_message,omitempty"` // 打招呼留言
}

// 通过入群申请
type ApproveJoinGroupApply struct {
	UserId  string `json:"user_id,omitempty"`  // 操作账号的平台用户ID，如：微信ID
	GroupId string `json:"group_id,omitempty"` // 平台群组ID
	ApplyId string `json:"apply_id,omitempty"` // 申请ID
}

// 查询群成员列表
type ListGroupMembers struct {
	QueryId string `json:"query_id,omitempty"` // 查询ID，用于匹配异步请求
	GroupId string `json:"group_id,omitempty"` // 平台群组ID
	Limit   int    `json:"limit,omitempty"`    // 查询数量
	After   string `json:"after,omitempty"`    // 上次查询返回 next，重新开始查询为空
}

// 群成员列表
type GroupMemberList struct {
	QueryId string         `json:"query_id,omitempty"` // 请求的查询ID，用于匹配异步请求
	Data    []*GroupMember `json:"data,omitempty"`     // 数据列表
	Next    string         `json:"next,omitempty"`     // 游标，用于下次查询请求的 after 参数
}

// 发送消息
type SendMessage struct {
	From             string           `json:"from,omitempty"`              // 消息发送者
	To               string           `json:"to,omitempty"`                // 消息接受者
	ConversationType ConversationType `json:"conversation_type,omitempty"` // 所属的会话类型
	Seq              int              `json:"seq,omitempty"`               // 序列号，在会话中唯一且有序增长，用于确保消息顺序
	MentionedType    MentionedType    `json:"mentioned_type,omitempty"`    // @用户类型
	MentionedUserIds []string         `json:"mentioned_users"`             // @用户列表
	SentAt           *time.Time       `json:"sent_at,omitempty"`           // 发送时间
	Payload          *MessagePayload  `json:"payload,omitempty"`           // 消息内容
}

// 评论类型
type CommentType int

const (
	CommentTypeUndefined CommentType = iota // 未定义评论
	CommentTypeText                         // 文本评论
)

// 动态类型
type MomentType int

const (
	MomentTypeUndefined MomentType = iota // 未定义动态
	MomentTypeText                        // 文本动态
	MomentTypeImage                       // 图文动态
	MomentTypeVideo                       // 视频动态
	MomentTypeLink                        // 链接动态
	MomentTypeWebcast                     // 直播
	MomentTypeFeed                        // 视频号
)

type Moment struct {
}

type MomentUpdate struct {
}

type MomentDelete struct{}

type MomentComment struct{}

type MomentCommentUpdate struct {
}

type MomentCommentDelete struct {
}

type MomentLike struct {
}

type MomentLikeDelete struct{}
