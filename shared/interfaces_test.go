package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ptrMockResource mirrors generated SDK types that use pointer receivers
// for GetId()/GetHref(). All generated SDK *Read types follow this pattern.
type ptrMockResource struct {
	id   string
	href string
	name string
}

func (m *ptrMockResource) GetId() string   { return m.id }
func (m *ptrMockResource) GetHref() string { return m.href }

// Compile-time interface satisfaction checks (pointer receivers — matches real SDKs)
var _ Identifiable = &ptrMockResource{}
var _ HasHref = &ptrMockResource{}
var _ Resource = &ptrMockResource{}

func TestExtractIDs(t *testing.T) {
	items := []ptrMockResource{
		{id: "aaa-111"},
		{id: "bbb-222"},
		{id: "ccc-333"},
	}
	ids := ExtractIDs(items)
	assert.Equal(t, []string{"aaa-111", "bbb-222", "ccc-333"}, ids)
}

func TestExtractIDsEmpty(t *testing.T) {
	ids := ExtractIDs([]ptrMockResource{})
	assert.Empty(t, ids)
}

func TestFindByIDFound(t *testing.T) {
	items := []ptrMockResource{
		{id: "aaa-111", name: "first"},
		{id: "bbb-222", name: "second"},
	}
	item, found := FindByID(items, "bbb-222")
	assert.True(t, found)
	assert.Equal(t, "second", item.name)
}

func TestFindByIDNotFound(t *testing.T) {
	items := []ptrMockResource{
		{id: "aaa-111"},
	}
	item, found := FindByID(items, "zzz-999")
	assert.False(t, found)
	assert.Nil(t, item)
}

func TestFindByIDEmpty(t *testing.T) {
	item, found := FindByID([]ptrMockResource{}, "aaa-111")
	assert.False(t, found)
	assert.Nil(t, item)
}

func TestFindByIDReturnsPointerIntoSlice(t *testing.T) {
	items := []ptrMockResource{
		{id: "aaa-111", name: "original"},
	}
	item, found := FindByID(items, "aaa-111")
	assert.True(t, found)

	// Mutating via returned pointer should modify the original slice
	item.name = "mutated"
	assert.Equal(t, "mutated", items[0].name)
}
