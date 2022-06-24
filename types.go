package uim

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
