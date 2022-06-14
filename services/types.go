package services

// 性别
type Gender int

const (
	GenderUnknown Gender = iota // 未知
	GenderMale                  // 男
	GenderFemale                // 女
)

// 账号在想状态
type Presence int

const (
	PresenceOnline             Presence = iota + 1 // 在线
	PresenceOffline                                // 离线
	PresenceLogout                                 // 登出
	PresenceDisabled                               // 禁用
	PresenceDisabledByProvider                     // 服务提供者封禁
)

// 评论类型
type CommentType int

const (
	CommentTypeUndefined CommentType = iota // 未定义评论
	CommentTypeText                         // 文本评论
)

// 关系类型
type RelationType string

const (
	RelationTypeFriendship RelationType = "friendship" // 好友关系
	RelationTypeFollowing  RelationType = "following"  // 关注关系
)

// 会话类型
type ConversationType int

const (
	ConversationTypePrivate         ConversationType = iota + 1 // 私聊会话
	ConversationTypeGroup                                       // 群聊会话
	ConversationTypeDiscussion                                  // 聊天室/讨论组会话
	ConversationTypeSystem                                      // 系统会话
	ConversationTypeCustomerService                             // 客服会话
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

// 消息@用户类型
type MentionedType int

const (
	MentionedTypeALL    MentionedType = iota + 1 // 所有人
	MentionedTypeSINGLE                          // 单个人
)
