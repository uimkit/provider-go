package uim

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
