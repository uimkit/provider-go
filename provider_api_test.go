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
		Type:       MetafieldTypeString,
		Key:        "str_value",
		Value:      "this is a string value",
	})
	assert.Nil(t, err)

	// integer value
	err = client.NewMetafield(&Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       MetafieldTypeInteger,
		Key:        "int_value",
		Value:      2789132749,
	})
	assert.Nil(t, err)

	// bool value
	err = client.NewMetafield(&Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       MetafieldTypeBool,
		Key:        "bool_value",
		Value:      true,
	})
	assert.Nil(t, err)

	// map value
	err = client.NewMetafield(&Metafield{
		Namespace:  "test",
		Resource:   "test_metafield",
		ResourceId: resourceId,
		Type:       MetafieldTypeMap,
		Key:        "map_value",
		Value: map[string]any{
			"id":   resourceId,
			"name": "map_value",
		},
	})
	assert.Nil(t, err)
}
