package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockResource is a minimal type satisfying Resource, Listable, and HasProperties
// for testing. It mirrors the shape of generated SDK *Read types.
type mockResource struct {
	id   string
	href string
	typ  string
	name string
}

func (m mockResource) GetId() string   { return m.id }
func (m mockResource) GetHref() string { return m.href }
func (m mockResource) GetType() string { return m.typ }
func (m mockResource) GetProperties() mockProperties {
	return mockProperties{Name: m.name}
}

type mockProperties struct {
	Name string
}

type mockList struct {
	items []mockResource
}

func (l mockList) GetItems() []mockResource { return l.items }

// Compile-time interface satisfaction checks
var _ Identifiable = mockResource{}
var _ HasHref = mockResource{}
var _ Resource = mockResource{}
var _ Listable[mockResource] = mockList{}
var _ HasProperties[mockProperties] = mockResource{}

func TestExtractIDs(t *testing.T) {
	items := []mockResource{
		{id: "aaa-111"},
		{id: "bbb-222"},
		{id: "ccc-333"},
	}
	ids := ExtractIDs(items)
	assert.Equal(t, []string{"aaa-111", "bbb-222", "ccc-333"}, ids)
}

func TestExtractIDs_Empty(t *testing.T) {
	ids := ExtractIDs([]mockResource{})
	assert.Empty(t, ids)
}

func TestFindByID_Found(t *testing.T) {
	items := []mockResource{
		{id: "aaa-111", name: "first"},
		{id: "bbb-222", name: "second"},
	}
	item, found := FindByID(items, "bbb-222")
	assert.True(t, found)
	assert.Equal(t, "second", item.name)
}

func TestFindByID_NotFound(t *testing.T) {
	items := []mockResource{
		{id: "aaa-111"},
	}
	_, found := FindByID(items, "zzz-999")
	assert.False(t, found)
}

func TestFindByID_Empty(t *testing.T) {
	_, found := FindByID([]mockResource{}, "aaa-111")
	assert.False(t, found)
}

func TestListItems(t *testing.T) {
	list := mockList{
		items: []mockResource{
			{id: "aaa-111"},
			{id: "bbb-222"},
		},
	}
	items := ListItems(list)
	assert.Len(t, items, 2)
	assert.Equal(t, "aaa-111", items[0].GetId())
}

func TestProperties(t *testing.T) {
	r := mockResource{id: "aaa-111", name: "test-zone"}
	props := Properties(r)
	assert.Equal(t, "test-zone", props.Name)
}
