package uim

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
)

const defaultUserId = "wxid_SPdd_nkhEYnA_Yf5gN5sp"

func newProviderClient() *Client {
	return NewClient(
		WithAppSecret("", ""),
		WithProvider("provider-go", "test"),
		WithDomain("localhost"),
		WithPort(9000),
		WithScheme(HTTP),
		WithDebug(true),
	)
}

func TestFriendApply(t *testing.T) {
	var err error
	client := newProviderClient()

	userId := defaultUserId
	applyUserId, _ := gonanoid.New()
	applyUserId = fmt.Sprintf("wxid_%s", applyUserId)
	applyId, _ := gonanoid.New()
	appliedAt := time.Now()

	err = client.NewFriendApply(&FriendApply{
		ID:     applyId,
		UserId: userId,
		ApplyUser: &IMUser{
			UserId:    applyUserId,
			CustomId:  "Kobe",
			Name:      "Kobe Bryant",
			Mobile:    "18666633332",
			Avatar:    "https://avatar.url",
			Gender:    GenderMale,
			Country:   "美国",
			Province:  "加利福尼亚",
			City:      "洛杉矶",
			Signature: "",
		},
		HelloMessage: "play with me",
		AppliedAt:    &appliedAt,
		Metadata:     map[string]any{"test": true},
	})
	assert.Nil(t, err)
}

func TestConversation(t *testing.T) {
	var err error
	client := newProviderClient()

	userId := defaultUserId
	partyGroupId, _ := gonanoid.New()
	partyGroupId = fmt.Sprintf("wxid_%s", partyGroupId)
	partyUserId, _ := gonanoid.New()
	partyUserId = fmt.Sprintf("wxid_%s", partyUserId)
	conversationId1, _ := gonanoid.New()
	conversationId2, _ := gonanoid.New()

	err = client.NewConversation(&Conversation{
		ConversationId: conversationId1,
		UserId:         userId,
		Type:           ConversationTypePrivate,
		Party: &ConversationParty{
			PartyId: partyUserId,
			Name:    "Jackson",
			Avatar:  "https://avatar.url",
		},
		Metadata: map[string]any{"test": true},
	})
	assert.Nil(t, err)

	err = client.NewConversation(&Conversation{
		ConversationId: conversationId2,
		UserId:         userId,
		Type:           ConversationTypeGroup,
		Party: &ConversationParty{
			PartyId: partyGroupId,
			Name:    "公司群",
			Avatar:  "https://avatar.url",
		},
		PrivateMetadata: map[string]any{"test": true},
	})
	assert.Nil(t, err)

	err = client.ConversationUpdated(&ConversationUpdate{
		UserId:          userId,
		Type:            ConversationTypePrivate,
		PartyId:         partyUserId,
		PrivateMetadata: map[string]any{"test": false},
	})
	assert.Nil(t, err)

	err = client.ConversationUpdated(&ConversationUpdate{
		UserId:   userId,
		Type:     ConversationTypeGroup,
		PartyId:  partyGroupId,
		Metadata: map[string]any{"test": false},
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

	err = client.NewGroup(&Group{
		UserId:  userId,
		GroupId: groupId,
		Owner: &IMUser{
			UserId:    ownerUserId,
			CustomId:  "Angela",
			Name:      "Angela（网红合作）☀️",
			Mobile:    "13000192287",
			Avatar:    "https://avatar.url",
			Gender:    GenderFemale,
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
	err = client.GroupUpdated(&GroupUpdate{
		UserId:  userId,
		GroupId: groupId,
		Owner: &IMUser{
			UserId:   updateOwnerUserId,
			CustomId: "Fiona",
			Name:     "Fiona（网红合作）☀️",
			Mobile:   "18988776655",
			Avatar:   "https://avatar.url",
			Gender:   GenderFemale,
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

	err = client.NewContact(&Contact{
		UserId:  userId,
		Alias:   "老李",
		Remark:  "公司同事",
		Blocked: false,
		Marked:  true,
		ContactUser: &IMUser{
			UserId:    contactUserId,
			CustomId:  "Angela",
			Name:      "Angela（网红合作）☀️",
			Mobile:    "13000192287",
			Avatar:    "https://avatar.url",
			Gender:    GenderFemale,
			Country:   "中国",
			Province:  "广东",
			City:      "深圳",
			Signature: "长期招募主播",
			Birthday:  &birthday,
		},
	})
	assert.Nil(t, err)

	updateAlias := "小孙"
	updateProvince := "江苏"
	updateCity := "苏州"
	err = client.ContactUpdated(&ContactUpdate{
		UserId: userId,
		ContactUser: &IMUserUpdate{
			UserId:   contactUserId,
			Province: &updateProvince,
			City:     &updateCity,
		},
		Alias: &updateAlias,
	})
	assert.Nil(t, err)
}

func TestIMAccount(t *testing.T) {
	var err error
	client := newProviderClient()

	err = client.NewAccount(&IMAccount{
		User: &IMUser{},
	})
	assert.Equal(t, InvalidEventDataErrorCode, err.(*ServerError).errorCode)

	birthday := time.Now().Add(-365 * 10 * 24 * 3600 * time.Second)
	userId, _ := gonanoid.New()
	userId = fmt.Sprintf("wxid_%s", userId)
	account := &IMAccount{
		User: &IMUser{
			UserId:    userId,
			CustomId:  "Angela",
			Name:      "Angela（网红合作）☀️",
			Mobile:    "13000192287",
			Avatar:    "https://avatar.url",
			Gender:    GenderFemale,
			Country:   "中国",
			Province:  "广东",
			City:      "深圳",
			Signature: "长期招募主播",
			Birthday:  &birthday,
		},
		Presence: PresenceInitializing,
	}
	err = client.NewAccount(account)
	assert.Nil(t, err)

	updatePresence := PresenceOnline
	updateMobile := "18900010002"
	updateName := "jenny"
	err = client.AccountUpdated(&IMAccountUpdate{
		User: &IMUserUpdate{
			UserId: userId,
			Name:   &updateName,
			Mobile: &updateMobile,
		},
		Presence: &updatePresence,
		Metadata: map[string]any{"test": true},
	})
	assert.Nil(t, err)

	err = client.AccountUpdated(&IMAccountUpdate{
		User: &IMUserUpdate{},
	})
	assert.Equal(t, InvalidEventDataErrorCode, err.(*ServerError).errorCode)

	err = client.AccountUpdated(&IMAccountUpdate{
		User: &IMUserUpdate{
			UserId: "fakeid",
		},
	})
	assert.Equal(t, ResourceNotFoundErrorCode, err.(*ServerError).errorCode)
}

func TestMetafield(t *testing.T) {
	var err error
	client := newProviderClient()

	resourceId := "test_metafield_" + strconv.FormatInt(time.Now().UnixMilli(), 36)

	// string value
	err = client.NewMetafield(&Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       MetafieldValueTypeString,
		Key:        "str_value",
		Value:      "this is the string value",
	})
	assert.Nil(t, err)

	// integer value
	err = client.NewMetafield(&Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       MetafieldValueTypeInteger,
		Key:        "int_value",
		Value:      2789132749,
	})
	assert.Nil(t, err)

	// bool value
	err = client.NewMetafield(&Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       MetafieldValueTypeBoolean,
		Key:        "boolean_value",
		Value:      true,
	})
	assert.Nil(t, err)

	// map value
	err = client.NewMetafield(&Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       MetafieldValueTypeJsonMap,
		Key:        "map_value",
		Value: map[string]any{
			"id":   resourceId,
			"name": "map_value",
		},
	})
	assert.Nil(t, err)

	err = client.MetafieldUpdated(&MetafieldUpdate{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       MetafieldValueTypeBoolean,
		Key:        "boolean_value",
		Value:      false,
	})
	assert.Nil(t, err)

	err = client.MetafieldUpdated(&MetafieldUpdate{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       MetafieldValueTypeString,
		Key:        "str_value",
		Value:      "hello world",
	})
	assert.Nil(t, err)

	err = client.MetafieldUpdated(&MetafieldUpdate{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       MetafieldValueTypeString,
		Key:        "not_found_value",
		Value:      "hello world",
	})
	assert.NotNil(t, err)
	assert.Equal(t, ResourceNotFoundErrorCode, err.(*ServerError).errorCode)
	t.Logf("%+v", err)

	var getMetafieldResp *GetMetafieldResponse
	getMetafieldResp, err = client.GetMetafield(&GetMetafieldRequest{
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Namespace:  "test",
		Key:        "str_value",
	})
	assert.Nil(t, err)
	assert.Equal(t, "hello world", getMetafieldResp.Value)

	_, err = client.GetMetafield(&GetMetafieldRequest{
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Namespace:  "test",
		Key:        "not_found_value",
	})
	assert.Equal(t, ResourceNotFoundErrorCode, err.(*ServerError).errorCode)
}
