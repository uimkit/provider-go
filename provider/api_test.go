package provider

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/hokaccha/go-prettyjson"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
	uim "github.com/uimkit/provider-go"
)

const defaultUserId = "wxid_SPdd_nkhEYnA_Yf5gN5sp"
const defaultGroupId = "wxid_6QOgsgb9QYM4od3rwZUM9"

func newProviderClient() *Client {
	clientId := os.Getenv("PROVIDER_CLIENT_ID")
	clientSecret := os.Getenv("PROVIDER_CLIENT_SECRET")
	audience := os.Getenv("PROVIDER_AUDIENCE")
	issuer := os.Getenv("ISSUER")
	return NewClient(
		uim.WithClient(clientId, clientSecret, audience),
		uim.WithServer(issuer, audience),
		WithProvider("provider-go", "test"),
		uim.WithBaseUrl("http://127.0.0.1:9000/providers/v1"),
		uim.WithDebug(true),
	)
}

func TestAuthorize(t *testing.T) {
	_, _, err := NewClient(uim.WithClient(
		os.Getenv("PROVIDER_CLIENT_ID"), os.Getenv("UIM_CLIENT_SECRET"), os.Getenv("UIM_AUDIENCE")),
	).Authorize()
	assert.Equal(t, uim.AuthenticationFailedErrorCode, err.(*uim.ClientError).ErrorCode())

	client := newProviderClient()
	accessToken, expiresIn, err := client.Authorize()
	assert.NotEmpty(t, accessToken)
	assert.Greater(t, expiresIn, int64(0))
	assert.Nil(t, err)
	t.Log(accessToken)

	claims, err := client.ValidateToken(accessToken)
	assert.Nil(t, err)
	t.Logf("%+v", claims)

	id1, _ := gonanoid.New()
	id2, _ := gonanoid.New()
	resourceId1 := "test_metafield_" + id1
	resourceId2 := "test_metafield_" + id2

	err = client.NewMetafield(&uim.Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId1,
		Type:       uim.MetafieldValueTypeString,
		Key:        "str_value",
		Value:      "this is the string value",
	})
	assert.Nil(t, err)

	client = NewClient(
		uim.WithClient(os.Getenv("UIM_CLIENT_ID"), os.Getenv("UIM_CLIENT_SECRET"), os.Getenv("UIM_AUDIENCE")),
		WithProvider("provider-go", "test"),
		uim.WithBaseUrl("http://127.0.0.1:9000/providers/v1"),
		uim.WithDebug(true),
	)
	err = client.NewMetafield(&uim.Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId2,
		Type:       uim.MetafieldValueTypeString,
		Key:        "str_value",
		Value:      "this is the string value",
	})
	assert.Equal(t, uim.UnauthorizedErrorCode, err.(*uim.ServerError).ErrorCode())
}

func TestRequestOptions(t *testing.T) {
	clientId := os.Getenv("PROVIDER_CLIENT_ID")
	clientSecret := os.Getenv("PROVIDER_CLIENT_SECRET")
	audience := os.Getenv("PROVIDER_AUDIENCE")
	client := NewClient(
		uim.WithClient(clientId, clientSecret, audience),
		WithProvider("provider-go", "test"),
		uim.WithDebug(true),
	)
	resourceId := "test_metafield_" + strconv.FormatInt(time.Now().UnixMilli(), 36)

	err := client.NewMetafield(&uim.Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       uim.MetafieldValueTypeString,
		Key:        "str_value",
		Value:      "this is the string value",
	}, uim.WithRequestBaseUrl("http://127.0.0.1:9000/providers/v1"))
	assert.Nil(t, err)
}

func TestMessage(t *testing.T) {
	var err error
	client := newProviderClient()

	messageId1, _ := gonanoid.New()
	messageId2, _ := gonanoid.New()
	messageId3, _ := gonanoid.New()
	userId1, _ := gonanoid.New()
	userId2, _ := gonanoid.New()
	userId3, _ := gonanoid.New()
	now := time.Now()

	message := &uim.Message{
		MessageId:      messageId1,
		Account:        defaultUserId,
		UserId:         userId1,
		Channel:        defaultGroupId,
		MentionedType:  uim.MentionedTypeAll,
		MentionedUsers: make([]string, 0),
		SentAt:         &now,
		Type:           uim.MessageTypeText,
		Text:           "hello",
		Revoked:        false,
	}
	err = client.NewMessage(message)
	assert.Nil(t, err)

	message.MessageId = messageId2
	message.MentionedType = uim.MentionedTypeUsers
	message.MentionedUsers = []string{userId2}
	message.Text = "yes"
	err = client.NewMessage(message)
	assert.Nil(t, err)

	message.MessageId = messageId3
	message.Account = defaultUserId
	message.UserId = defaultUserId
	message.Channel = userId3

	message.MentionedType = uim.MentionedTypeNone
	message.MentionedUsers = nil
	message.Text = "在不？"
	err = client.NewMessage(message)
	assert.Nil(t, err)

	updateRevoke := true
	err = client.MessageUpdated(&uim.MessageUpdate{
		MessageId: messageId3,
		Revoked:   &updateRevoke,
		Metadata:  map[string]any{"test": true},
	})
	assert.Nil(t, err)
}

func TestGroupJoinApply(t *testing.T) {
	var err error
	client := newProviderClient()

	userId := defaultUserId
	groupId := defaultGroupId
	applyUserId, _ := gonanoid.New()
	applyUserId = fmt.Sprintf("wxid_%s", applyUserId)
	applyId, _ := gonanoid.New()
	appliedAt := time.Now().Add(-10 * time.Minute)

	err = client.NewJoinGroupApply(&uim.JoinGroupApply{
		ID:      applyId,
		UserId:  userId,
		GroupId: groupId,
		ApplyUser: &uim.IMUser{
			UserId: applyUserId,
			Name:   "John Stockton",
		},
		HelloMessage: "你好啊",
		AppliedAt:    &appliedAt,
		Metadata:     map[string]any{"test": true},
	})
	assert.Nil(t, err)
}

func TestGroupInvitation(t *testing.T) {
	var err error
	client := newProviderClient()

	userId := defaultUserId
	groupId := defaultGroupId
	inviterUserId, _ := gonanoid.New()
	inviterUserId = fmt.Sprintf("wxid_%s", inviterUserId)
	inviteId, _ := gonanoid.New()
	invitedAt := time.Now().Add(-10 * time.Minute)

	err = client.NewGroupInvitation(&uim.GroupInvitation{
		ID:      inviteId,
		UserId:  userId,
		GroupId: groupId,
		Inviter: &uim.IMUser{
			UserId: inviterUserId,
			Name:   "Karl Marlone",
		},
		HelloMessage: "你好啊",
		InvitedAt:    &invitedAt,
		Metadata:     map[string]any{"test": true},
	})
	assert.Nil(t, err)
}

func TestGroupMember(t *testing.T) {
	var err error
	client := newProviderClient()

	userId, _ := gonanoid.New()
	userId = fmt.Sprintf("wxid_%s", userId)
	userId2, _ := gonanoid.New()
	userId2 = fmt.Sprintf("wxid_%s", userId2)
	groupId := defaultGroupId
	memberId, _ := gonanoid.New()
	memberId2, _ := gonanoid.New()

	err = client.NewGroupMember(&uim.GroupMember{
		GroupId:  groupId,
		MemberId: memberId,
		IMUser: uim.IMUser{
			UserId:   userId,
			CustomId: "Chris Webber",
		},
		IsOwner:  false,
		IsAdmin:  true,
		Alias:    "李伟波",
		Metadata: map[string]any{"test": true},
	})
	assert.Nil(t, err)

	updateIsOwner := true
	err = client.GroupMemberUpdated(&uim.GroupMemberUpdate{
		GroupId:         groupId,
		MemberId:        memberId,
		IsOwner:         &updateIsOwner,
		PrivateMetadata: map[string]any{"test": false},
	})
	assert.Nil(t, err)

	updateName := "韦伯"
	err = client.GroupMemberUpdated(&uim.GroupMemberUpdate{
		GroupId:  groupId,
		MemberId: memberId,
		IMUserUpdate: uim.IMUserUpdate{
			UserId: userId,
			Name:   &updateName,
		},
	})
	assert.Nil(t, err)

	err = client.NewGroupMember(&uim.GroupMember{
		GroupId:  groupId,
		MemberId: memberId2,
		IMUser: uim.IMUser{
			UserId:   userId2,
			CustomId: "Mike Bibby",
		},
		IsOwner:  false,
		IsAdmin:  true,
		Alias:    "麦贝比",
		Metadata: map[string]any{"test": true},
	})
	assert.Nil(t, err)
}

func TestFriendApply(t *testing.T) {
	var err error
	client := newProviderClient()

	userId := defaultUserId
	applyUserId, _ := gonanoid.New()
	applyUserId = fmt.Sprintf("wxid_%s", applyUserId)
	applyId, _ := gonanoid.New()
	appliedAt := time.Now()

	err = client.NewFriendApply(&uim.FriendApply{
		IMUser: uim.IMUser{
			UserId:    applyUserId,
			CustomId:  "Kobe",
			Name:      "Kobe Bryant",
			Mobile:    "18666633332",
			Avatar:    "https://avatar.url",
			Gender:    uim.GenderMale,
			Country:   "美国",
			Province:  "加利福尼亚",
			City:      "洛杉矶",
			Signature: "",
		},
		ID:           applyId,
		Account:      userId,
		HelloMessage: "play with me",
		AppliedAt:    &appliedAt,
		Metadata:     map[string]any{"test": true},
	})
	assert.Nil(t, err)
}

func TestGroup(t *testing.T) {
	var err error
	client := newProviderClient()

	userId := defaultUserId
	groupId, _ := gonanoid.New()
	groupId = fmt.Sprintf("wxid_%s", groupId)
	ownerUserId, _ := gonanoid.New()
	ownerUserId = fmt.Sprintf("wxid_%s", ownerUserId)
	birthday := time.Now().Add(-365 * 10 * 24 * 3600 * time.Second)

	err = client.NewGroup(&uim.Group{
		Account: userId,
		GroupId: groupId,
		Owner: &uim.IMUser{
			UserId:    ownerUserId,
			CustomId:  "Angela",
			Name:      "Angela（网红合作）☀️",
			Mobile:    "13000192287",
			Avatar:    "https://avatar.url",
			Gender:    uim.GenderFemale,
			Country:   "中国",
			Province:  "广东",
			City:      "深圳",
			Signature: "长期招募主播",
			Birthday:  &birthday,
		},
		Name:         "福利一群",
		Avatar:       "https://avatar.url",
		Announcement: "大家记得修改群公告",
	})
	assert.Nil(t, err)

	updateOwnerUserId, _ := gonanoid.New()
	updateOwnerUserId = fmt.Sprintf("wxid_%s", updateOwnerUserId)
	updateBirthday := time.Now().Add(-365 * 5 * 24 * 3600 * time.Second)
	updateAnnouncement := "大家记得修改群公告，发广告者提出"
	err = client.GroupUpdated(&uim.GroupUpdate{
		GroupId: groupId,
		Owner: &uim.IMUser{
			UserId:   updateOwnerUserId,
			CustomId: "Fiona",
			Name:     "Fiona（网红合作）☀️",
			Mobile:   "18988776655",
			Avatar:   "https://avatar.url",
			Gender:   uim.GenderFemale,
			Country:  "中国",
			Province: "江苏",
			City:     "苏州",
			Birthday: &updateBirthday,
		},
		Announcement: &updateAnnouncement,
	})
	assert.Nil(t, err)
}

func TestContact(t *testing.T) {
	var err error
	client := newProviderClient()

	birthday := time.Now().Add(-365 * 10 * 24 * 3600 * time.Second)
	userId := defaultUserId
	contactUserId, _ := gonanoid.New()
	contactUserId = fmt.Sprintf("wxid_%s", contactUserId)

	err = client.NewContact(&uim.Contact{
		IMUser: uim.IMUser{
			UserId:    contactUserId,
			CustomId:  "Angela",
			Name:      "Angela（网红合作）☀️",
			Mobile:    "13000192287",
			Avatar:    "https://avatar.url",
			Gender:    uim.GenderFemale,
			Country:   "中国",
			Province:  "广东",
			City:      "深圳",
			Signature: "长期招募主播",
			Birthday:  &birthday,
		},
		Account: userId,
		Alias:   "老李",
		Remark:  "公司同事",
		Blocked: false,
		Marked:  true,
	})
	assert.Nil(t, err)
}

func TestIMAccount(t *testing.T) {
	var err error
	client := newProviderClient()

	err = client.NewAccount(&uim.IMAccount{})
	assert.Equal(t, uim.InvalidEventDataErrorCode, err.(*uim.ServerError).ErrorCode())

	birthday := time.Now().Add(-365 * 10 * 24 * 3600 * time.Second)
	userId, _ := gonanoid.New()
	userId = fmt.Sprintf("wxid_%s", userId)
	account := &uim.IMAccount{
		IMUser: uim.IMUser{
			UserId:    userId,
			CustomId:  "Angela",
			Name:      "Angela（网红合作）☀️",
			Mobile:    "13000192287",
			Avatar:    "https://avatar.url",
			Gender:    uim.GenderFemale,
			Country:   "中国",
			Province:  "广东",
			City:      "深圳",
			Signature: "长期招募主播",
			Birthday:  &birthday,
		},
		Presence: uim.PresenceInitializing,
	}
	err = client.NewAccount(account)
	assert.Nil(t, err)

	updatePresence := uim.PresenceOnline
	updateMobile := "18900010002"
	updateName := "jenny"
	err = client.AccountUpdated(&uim.IMAccountUpdate{
		IMUserUpdate: uim.IMUserUpdate{
			UserId: userId,
			Name:   &updateName,
			Mobile: &updateMobile,
		},
		Presence: &updatePresence,
		Metadata: map[string]any{"test": true},
	})
	assert.Nil(t, err)

	err = client.AccountUpdated(&uim.IMAccountUpdate{})
	assert.Equal(t, uim.InvalidEventDataErrorCode, err.(*uim.ServerError).ErrorCode())

	err = client.AccountUpdated(&uim.IMAccountUpdate{
		IMUserUpdate: uim.IMUserUpdate{
			UserId: "fakeid",
		},
	})
	assert.Equal(t, uim.ResourceNotFoundErrorCode, err.(*uim.ServerError).ErrorCode())
}

func TestMetafield(t *testing.T) {
	var err error
	client := newProviderClient()

	resourceId := "test_metafield_" + strconv.FormatInt(time.Now().UnixMilli(), 36)

	// string value
	err = client.NewMetafield(&uim.Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       uim.MetafieldValueTypeString,
		Key:        "str_value",
		Value:      "this is the string value",
	})
	assert.Nil(t, err)

	// integer value
	err = client.NewMetafield(&uim.Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       uim.MetafieldValueTypeInteger,
		Key:        "int_value",
		Value:      2789132749,
	})
	assert.Nil(t, err)

	// bool value
	err = client.NewMetafield(&uim.Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       uim.MetafieldValueTypeBoolean,
		Key:        "boolean_value",
		Value:      true,
	})
	assert.Nil(t, err)

	// map value
	err = client.NewMetafield(&uim.Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       uim.MetafieldValueTypeJsonMap,
		Key:        "map_value",
		Value: map[string]any{
			"id":   resourceId,
			"name": "map_value",
		},
	})
	assert.Nil(t, err)

	err = client.MetafieldUpdated(&uim.MetafieldUpdate{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       uim.MetafieldValueTypeBoolean,
		Key:        "boolean_value",
		Value:      false,
	})
	assert.Nil(t, err)

	err = client.MetafieldUpdated(&uim.MetafieldUpdate{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       uim.MetafieldValueTypeString,
		Key:        "str_value",
		Value:      "hello world",
	})
	assert.Nil(t, err)

	err = client.MetafieldUpdated(&uim.MetafieldUpdate{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       uim.MetafieldValueTypeString,
		Key:        "not_found_value",
		Value:      "hello world",
	})
	assert.NotNil(t, err)
	assert.Equal(t, uim.ResourceNotFoundErrorCode, err.(*uim.ServerError).ErrorCode())
	t.Logf("%+v", err)

	var getMetafieldResp *uim.GetMetafieldResponse
	getMetafieldResp, err = client.GetMetafield(&uim.GetMetafieldRequest{
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Namespace:  "test",
		Key:        "str_value",
	})
	assert.Nil(t, err)
	assert.Equal(t, "hello world", getMetafieldResp.Value)

	_, err = client.GetMetafield(&uim.GetMetafieldRequest{
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Namespace:  "test",
		Key:        "not_found_value",
	})
	assert.Equal(t, uim.ResourceNotFoundErrorCode, err.(*uim.ServerError).ErrorCode())
}

func TestJSON(t *testing.T) {
	birthday := time.Now().Add(-365 * 10 * 24 * 3600 * time.Second)
	userId, _ := gonanoid.New()
	userId = fmt.Sprintf("wxid_%s", userId)
	b, err := prettyjson.Marshal(&uim.IMAccount{
		IMUser: uim.IMUser{
			UserId:    userId,
			CustomId:  "Angela",
			Name:      "Angela（网红合作）☀️",
			Mobile:    "13000192287",
			Avatar:    "https://avatar.url",
			Gender:    uim.GenderFemale,
			Country:   "中国",
			Province:  "广东",
			City:      "深圳",
			Signature: "长期招募主播",
			Birthday:  &birthday,
			Metadata:  map[string]any{"user": true},
		},
		Presence: uim.PresenceInitializing,
		Metadata: map[string]any{"account": true},
	})
	assert.Nil(t, err)
	t.Logf("%s", string(b))

	str := `{
		"avatar": "https://avatar.url",
		"birthday": "2012-11-07T01:07:43.296042+08:00",
		"city": "深圳",
		"country": "中国",
		"custom_id": "Angela",
		"gender": 2,
		"metadata": {
			"account": true
		},
		"mobile": "13000192287",
		"name": "Angela（网红合作）☀️",
		"province": "广东",
		"signature": "长期招募主播",
		"user_id": "wxid_dAsOimcs5OWwQwWu_57aI"
	}`
	account := &uim.IMAccount{}
	err = json.Unmarshal([]byte(str), account)
	assert.Nil(t, err)
	t.Logf("%+v", account)
}
