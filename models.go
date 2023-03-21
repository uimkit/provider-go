package uim

import (
	"time"
)

// 游标查询方向
type CursorDirection string

const (
	CursorDirectionBefore CursorDirection = "before"
	CursorDirectionAfter  CursorDirection = "after"
)

// 游标查询请求
type CursorQuery struct {
	Cursor    string          `json:"cursor,omitempty"`    // 游标
	Limit     int32           `json:"limit,omitempty"`     // 查询数量
	Direction CursorDirection `json:"direction,omitempty"` // 查询方向
}

// 游标查询结果扩展信息
type CursorExtra struct {
	Limit       int32 `json:"limit,omitempty"` // 平台可能对单词查询数量有限制，这里返回最终实际的查询数量
	HasPrevious bool  `json:"has_previous"`    // 是否有更多数据
	HasNext     bool  `json:"has_next"`        // 是否有更多数据
}

// 游标查询结果条目
type CursorItem[T any] struct {
	Cursor string `json:"cursor,omitempty"` // 对应的游标
	Item   T      `json:"item,omitempty"`   // 对应的数据
}

// 游标查询结果
type CursorPage[T any] struct {
	Extra CursorExtra     `json:"extra,omitempty"` // 扩展信息
	Items []CursorItem[T] `json:"items,omitempty"` // 结果集
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
	OpenId          string         `json:"open_id,omitempty"`          // 实际的平台用户ID，如：微信ID
	CustomId        string         `json:"custom_id,omitempty"`        // 用户自定义ID
	Username        string         `json:"username,omitempty"`         // 用户账户
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
	PresenceInactive     Presence = iota // 未激活
	PresenceActive                       // 在线
	PresenceDisconnected                 // 掉线
	PresenceDisabled                     // 停用
	PresenceBanned                       // 封号
)

// 账号
type IMAccount struct {
	IMUser                         // 账号用户信息
	Presence        Presence       `json:"presence,omitempty"`         // 状态
	State           string         `json:"state,omitempty"`            // 授权账号时客户传来的自定义数据，透传回去
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 账号变更
type IMAccountUpdate struct {
	IMUserUpdate                   // 更新账号用户信息
	Presence        *Presence      `json:"presence,omitempty"`         // 状态
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 联系人
type Contact struct {
	IMUser                         // 联系人用户信息
	Account         string         `json:"account,omitempty"`          // 所属账号的平台用户ID
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
	IMUser                         // 申请人用户信息
	ID              string         `json:"id,omitempty"`               // 申请ID
	Account         string         `json:"account,omitempty"`          // 接收申请的账号的平台用户ID
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
	MessageTypeAudio    MessageType = "audio"    // 语音消息
	MessageTypeVideo    MessageType = "video"    // 视频消息
	MessageTypeLink     MessageType = "link"     // 链接消息
	MessageTypeLocation MessageType = "location" // 位置消息
)

type ImageInfo struct {
	URL    string `json:"url,omitempty"`    // 图片URL
	Width  int    `json:"width,omitempty"`  // 宽度（像素）
	Height int    `json:"height,omitempty"` // 高度（像素）
}

type ImageMessageBody struct {
	Size   int          `json:"size,omitempty"`   // 大小（字节）
	Format string       `json:"format,omitempty"` // 类型，如：png、jpeg
	MD5    string       `json:"md5,omitempty"`    // 文件内容MD5
	Infos  []*ImageInfo `json:"infos,omitempty"`  // 图片信息，索引0是原图，1是中图，2是小图
}

type AudioMessageBody struct {
	URL      string `json:"url,omitempty"`      // 语音URL
	Duration int    `json:"duration,omitempty"` // 时长（毫秒）
	Size     int    `json:"size,omitempty"`     // 大小（字节）
	Format   string `json:"format,omitempty"`   // 类型，如：mp3
	MD5      string `json:"md5,omitempty"`      // 文件内容MD5
}

type VideoMessageBody struct {
	URL      string `json:"url,omitempty"`      // 视频URL
	Duration int    `json:"duration,omitempty"` // 时长（毫秒）
	Width    int    `json:"width,omitempty"`    // 宽度（像素）
	Height   int    `json:"height,omitempty"`   // 高度（像素）
	Size     int    `json:"size,omitempty"`     // 大小（字节）
	Format   string `json:"format,omitempty"`   // 类型，如：mp4
	MD5      string `json:"md5,omitempty"`      // 文件内容MD5
	Snapshot string `json:"snapshot,omitempty"` // 封面图
}

// 消息
type Message struct {
	MessageId       string            `json:"message_id,omitempty"`       // 平台消息ID
	Channel         string            `json:"channel,omitempty"`          // 消息收发地址，账号回复消息时会发送到此地址
	Account         string            `json:"account,omitempty"`          // 归属账号的平台用户ID
	UserId          string            `json:"user_id,omitempty"`          // 消息发送人平台用户ID
	Type            MessageType       `json:"type,omitempty"`             // 消息类型
	Text            string            `json:"text,omitempty"`             // 文本消息
	Image           *ImageMessageBody `json:"image,omitempty"`            // 图片消息、视频消息封面
	Audio           *AudioMessageBody `json:"audio,omitempty"`            // 语音消息
	Video           *VideoMessageBody `json:"video,omitempty"`            // 视频消息
	MentionedType   MentionedType     `json:"mentioned_type,omitempty"`   // @用户类型
	MentionedUsers  []string          `json:"mentioned_users"`            // @用户列表，是平台用户ID
	SentAt          *time.Time        `json:"sent_at,omitempty"`          // 发送时间
	Revoked         bool              `json:"revoked,omitempty"`          // 是否撤回
	Metadata        map[string]any    `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any    `json:"private_metadata,omitempty"` // 私有元数据
	State           string            `json:"state,omitempty"`            // 发送消息时携带的业务自定义数据，发送后返回消息会透传给业务方
}

// 发送消息
type SendMessageRequest struct {
	Account          string            `json:"account,omitempty"`           // 归属账号的平台用户ID
	UserId           string            `json:"user_id,omitempty"`           // 消息发送人平台用户ID
	Channel          string            `json:"channel,omitempty"`           // 消息发送地址
	ConversationType ConversationType  `json:"conversation_type,omitempty"` // 消息所属的会话类型
	Type             MessageType       `json:"type,omitempty"`              // 消息类型
	Text             string            `json:"text,omitempty"`              // 文本消息
	Image            *ImageMessageBody `json:"image,omitempty"`             // 图片消息、视频消息封面
	Audio            *AudioMessageBody `json:"audio,omitempty"`             // 语音消息
	Video            *VideoMessageBody `json:"video,omitempty"`             // 视频消息
	Seq              int               `json:"seq,omitempty"`               // 序列号，在会话中唯一且有序增长，用于确保消息顺序
	MentionedType    MentionedType     `json:"mentioned_type,omitempty"`    // @用户类型
	MentionedUsers   []string          `json:"mentioned_users"`             // @用户列表，是平台用户ID
}

// 发送消息结果
type SendMessageResponse struct {
	BaseResponse
	Message
}

// 消息变更
type MessageUpdate struct {
	MessageId       string         `json:"message_id,omitempty"`       // 平台消息ID
	Revoked         *bool          `json:"revoked,omitempty"`          // 是否撤回
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 查询消息地址的信息
type GetChannelInfoRequest struct {
	Channel string `json:"channel,omitempty"` // 消息收发地址
}

// 查询消息地址的返回
type GetChannelInfoResponse struct {
	BaseResponse
	Group *Group  `json:"group,omitempty"` // 如果是群组的地址，返回群组信息
	User  *IMUser `json:"user,omitempty"`  // 如果是用户的地址，返回用户信息
}

// 群组
type Group struct {
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	Account         string         `json:"account,omitempty"`          // 归属账号的平台用户ID
	Owner           *IMUser        `json:"owner,omitempty"`            // 群主信息
	Name            string         `json:"name,omitempty"`             // 名称
	Alias           string         `json:"alias,omitempty"`            // 备注名
	Avatar          string         `json:"avatar,omitempty"`           // 头像URL
	Announcement    string         `json:"announcement,omitempty"`     // 群公告
	Description     string         `json:"description,omitempty"`      // 群介绍
	QRCode          string         `json:"qrcode,omitempty"`           // 二维码
	Remark          string         `json:"remark,omitempty"`           // 备注说明
	Marked          bool           `json:"marked,omitempty"`           // 是否星标
	Mute            bool           `json:"mute,omitempty"`             // 是否禁言
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 群组变更
type GroupUpdate struct {
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	Owner           *IMUser        `json:"owner,omitempty"`            // 群主变更
	Name            *string        `json:"name,omitempty"`             // 名称
	Alias           *string        `json:"alias,omitempty"`            // 备注名
	Avatar          *string        `json:"avatar,omitempty"`           // 头像URL
	Announcement    *string        `json:"announcement,omitempty"`     // 群公告
	Description     *string        `json:"description,omitempty"`      // 群介绍
	QRCode          *string        `json:"qrcode,omitempty"`           // 二维码
	Remark          *string        `json:"remark,omitempty"`           // 备注说明
	Marked          *bool          `json:"marked,omitempty"`           // 是否星标
	Mute            *bool          `json:"mute,omitempty"`             // 是否禁言
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 群成员角色
type GroupMemberRole int

const (
	GroupMemberRoleMember GroupMemberRole = iota // 普通成员
	GroupMemberRoleAdmin                         // 管理员
	GroupMemberRoleOwner                         // 群主
)

// 群组成员
type GroupMember struct {
	IMUser                          // 群成员用户资料
	GroupId         string          `json:"group_id,omitempty"`         // 平台群组ID
	MemberId        string          `json:"member_id,omitempty"`        // 平台群成员ID
	Role            GroupMemberRole `json:"role,omitempty"`             // 角色
	Alias           string          `json:"alias,omitempty"`            // 群内备注名
	JoinedAt        *time.Time      `json:"joined_at,omitempty"`        // 入群时间
	Metadata        map[string]any  `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any  `json:"private_metadata,omitempty"` // 私有元数据
}

// 群组成员变更
type GroupMemberUpdate struct {
	IMUserUpdate                     // 群成员用户资料更新
	GroupId         string           `json:"group_id,omitempty"`         // 平台群组ID
	MemberId        string           `json:"member_id,omitempty"`        // 平台群成员ID
	Role            *GroupMemberRole `json:"role,omitempty"`             // 角色
	Alias           *string          `json:"alias,omitempty"`            // 群内备注名
	JoinedAt        *time.Time       `json:"joined_at,omitempty"`        // 入群时间
	Metadata        map[string]any   `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any   `json:"private_metadata,omitempty"` // 私有元数据
}

// 入群邀请
type GroupInvitation struct {
	ID              string         `json:"id,omitempty"`               // 入群邀请ID
	UserId          string         `json:"user_id,omitempty"`          // 收到邀请的平台用户ID
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	Inviter         *IMUser        `json:"inviter,omitempty"`          // 邀请人信息
	HelloMessage    string         `json:"hello_message,omitempty"`    // 打招呼留言
	InvitedAt       *time.Time     `json:"invited_at,omitempty"`       // 邀请时间
	Source          string         `json:"source,omitempty"`           // 邀请来源
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 入群申请
type GroupApply struct {
	ID              string         `json:"id,omitempty"`               // 入群申请ID
	UserId          string         `json:"user_id,omitempty"`          // 收到申请的平台用户ID
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	ApplyUser       *IMUser        `json:"apply_user,omitempty"`       // 申请用户信息
	HelloMessage    string         `json:"hello_message,omitempty"`    // 打招呼留言
	AppliedAt       *time.Time     `json:"applied_at,omitempty"`       // 申请时间
	Source          string         `json:"source,omitempty"`           // 申请来源
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
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

// 通过好友请求
type AcceptFriendApplyRequest struct {
	ApplyId string `json:"apply_id,omitempty"` // 好友请求ID
	UserId  string `json:"user_id,omitempty"`  // 账号的平台用户ID
}

// 通过好友返回
type AcceptFriendApplyResponse struct {
	BaseResponse
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

// 消息类型
type MomentType string

const (
	MomentTypeText  MomentType = "text"  // 文本动态
	MomentTypeImage MomentType = "image" // 图文动态
	MomentTypeVideo MomentType = "video" // 视频动态
)

// 图片动态内容
type ImageMomentContent struct {
	Size   int          `json:"size,omitempty"`   // 大小（字节）
	Format string       `json:"format,omitempty"` // 类型，如：png、jpeg
	MD5    string       `json:"md5,omitempty"`    // 文件内容MD5
	Infos  []*ImageInfo `json:"infos,omitempty"`  // 图片信息，索引0是原图，1是中图，2是小图
}

// 视频动态内容
type VideoMomentContent struct {
	URL      string `json:"url,omitempty"`      // 视频URL
	Duration int    `json:"duration,omitempty"` // 时长（毫秒）
	Width    int    `json:"width,omitempty"`    // 宽度（像素）
	Height   int    `json:"height,omitempty"`   // 高度（像素）
	Size     int    `json:"size,omitempty"`     // 大小（字节）
	Format   string `json:"format,omitempty"`   // 类型，如：mp4
	MD5      string `json:"md5,omitempty"`      // 文件内容MD5
	Snapshot string `json:"snapshot,omitempty"` // 封面图
}

// 评论
type Comment struct {
	CommentId   string     `json:"comment_id,omitempty"`    // 平台评论ID
	User        *IMUser    `json:"user,omitempty"`          // 评论人的信息
	CommentedAt *time.Time `json:"commented_at,omitempty"`  // 评论时间
	ReplyTo     string     `json:"reply_to,omitempty"`      // 回复的平台评论ID
	ReplyToUser *IMUser    `json:"reply_to_user,omitempty"` // 回复的用户信息
	Text        string     `json:"text,omitempty"`          // 评论文本
}

// 点赞
type Like struct {
	LikeId  string     `json:"like_id,omitempty"`  // 平台点赞ID
	User    *IMUser    `json:"user,omitempty"`     // 点赞人的信息
	LikedAt *time.Time `json:"liked_at,omitempty"` // 点赞时间
}

// 动态
type Moment struct {
	MomentId    string                `json:"moment_id,omitempty"`    // 平台动态ID
	Account     string                `json:"account,omitempty"`      // 归属账号的平台用户ID
	User        *IMUser               `json:"user,omitempty"`         // 动态发布人的信息
	PublishedAt *time.Time            `json:"published_at,omitempty"` // 发布时间
	Type        MomentType            `json:"type,omitempty"`         // 动态类型
	Text        string                `json:"text,omitempty"`         // 文案
	Images      []*ImageMomentContent `json:"images,omitempty"`       // 图片
	Video       *VideoMomentContent   `json:"video,omitempty"`        // 视频
	Comments    CursorPage[*Comment]  `json:"comments,omitempty"`     // 评论
	Likes       CursorPage[*Like]     `json:"likes,omitempty"`        // 点赞
}

// 查询动态列表请求
type GetMomentListRequest struct {
	CursorQuery
	Account string `json:"account,omitempty"` // 所属账号的平台用户ID
	UserId  string `json:"user_id,omitempty"` // 发布人的平台用户ID，如果查询账号自己发布的动态，则 account 与 user_id 都为账号的平台用户ID
}

// 查询懂啊提列表请求返回
type GetMomentListResponse struct {
	BaseResponse
	CursorPage[*Moment]
}
