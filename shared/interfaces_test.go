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
func (m *ptrMockResource) GetProperties() mockProperties {
	return mockProperties{Name: m.name}
}

type mockProperties struct {
	Name string
}

type mockList struct {
	items []ptrMockResource
}

func (l mockList) GetItems() []ptrMockResource { return l.items }

// Compile-time interface satisfaction checks (pointer receivers — matches real SDKs)
var _ Identifiable = &ptrMockResource{}
var _ HasHref = &ptrMockResource{}
var _ Resource = &ptrMockResource{}
var _ Listable[ptrMockResource] = mockList{}
var _ HasProperties[mockProperties] = &ptrMockResource{}

func TestExtractIDs(t *testing.T) {
	items := []ptrMockResource{
		{id: "aaa-111"},
		{id: "bbb-222"},
		{id: "ccc-333"},
	}
	ids := ExtractIDs(items)
	assert.Equal(t, []string{"aaa-111", "bbb-222", "ccc-333"}, ids)
}

func TestExtractIDs_Empty(t *testing.T) {
	ids := ExtractIDs([]ptrMockResource{})
	assert.Empty(t, ids)
}

func TestFindByID_Found(t *testing.T) {
	items := []ptrMockResource{
		{id: "aaa-111", name: "first"},
		{id: "bbb-222", name: "second"},
	}
	item, found := FindByID(items, "bbb-222")
	assert.True(t, found)
	assert.Equal(t, "second", item.name)
}

func TestFindByID_NotFound(t *testing.T) {
	items := []ptrMockResource{
		{id: "aaa-111"},
	}
	item, found := FindByID(items, "zzz-999")
	assert.False(t, found)
	assert.Nil(t, item)
}

func TestFindByID_Empty(t *testing.T) {
	item, found := FindByID([]ptrMockResource{}, "aaa-111")
	assert.False(t, found)
	assert.Nil(t, item)
}

func TestFindByID_ReturnsPointerIntoSlice(t *testing.T) {
	items := []ptrMockResource{
		{id: "aaa-111", name: "original"},
	}
	item, found := FindByID(items, "aaa-111")
	assert.True(t, found)

	// Mutating via returned pointer should modify the original slice
	item.name = "mutated"
	assert.Equal(t, "mutated", items[0].name)
}

func TestListItems(t *testing.T) {
	list := mockList{
		items: []ptrMockResource{
			{id: "aaa-111"},
			{id: "bbb-222"},
		},
	}
	items := ListItems(list)
	assert.Len(t, items, 2)
	assert.Equal(t, "aaa-111", (*ptrMockResource)(&items[0]).GetId())
}

func TestProperties(t *testing.T) {
	r := &ptrMockResource{id: "aaa-111", name: "test-zone"}
	props := Properties(r)
	assert.Equal(t, "test-zone", props.Name)
}
