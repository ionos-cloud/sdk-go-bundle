package shared_test

import (
	"fmt"

	"github.com/ionos-cloud/sdk-go-bundle/shared"
)

// ──────────────────────────────────────────────────────────────────────────────
// Mock types that mirror the generated SDK model shape (e.g. dns.ZoneRead,
// dns.Zone, dns.ZoneReadList). All generated SDK types use pointer receivers
// and value-typed Items slices — these mocks follow the same pattern.
// ──────────────────────────────────────────────────────────────────────────────

// Zone mirrors a properties type like dns.Zone.
type Zone struct {
	ZoneName string
}

func (o *Zone) GetZoneName() string { return o.ZoneName }

// ZoneRead mirrors a read-model type like dns.ZoneRead.
type ZoneRead struct {
	Id         string
	Href       string
	Properties Zone
}

func (o *ZoneRead) GetId() string         { return o.Id }
func (o *ZoneRead) GetHref() string       { return o.Href }
func (o *ZoneRead) GetProperties() Zone   { return o.Properties }

// ZoneReadList mirrors a list-response type like dns.ZoneReadList.
type ZoneReadList struct {
	Items []ZoneRead
}

func (o *ZoneReadList) GetItems() []ZoneRead { return o.Items }

// RecordRead mirrors a read-model type like dns.RecordRead.
type RecordRead struct {
	Id   string
	Href string
}

func (o *RecordRead) GetId() string   { return o.Id }
func (o *RecordRead) GetHref() string { return o.Href }

// ──────────────────────────────────────────────────────────────────────────────
// Examples — these appear in godoc and are executed by `go test`.
// ──────────────────────────────────────────────────────────────────────────────

func ExampleExtractIDs() {
	// Simulate: list, _, err := client.ZonesApi.ZonesGet(ctx).Execute()
	list := &ZoneReadList{
		Items: []ZoneRead{
			{Id: "aaa-111", Href: "/zones/aaa-111", Properties: Zone{ZoneName: "example.com"}},
			{Id: "bbb-222", Href: "/zones/bbb-222", Properties: Zone{ZoneName: "example.org"}},
			{Id: "ccc-333", Href: "/zones/ccc-333", Properties: Zone{ZoneName: "example.net"}},
		},
	}

	// Works directly with GetItems() — no []T to []*T conversion needed.
	ids := shared.ExtractIDs(list.GetItems())
	fmt.Println(ids)
	// Output: [aaa-111 bbb-222 ccc-333]
}

func ExampleFindByID() {
	// Simulate: list, _, err := client.ZonesApi.ZonesGet(ctx).Execute()
	list := &ZoneReadList{
		Items: []ZoneRead{
			{Id: "aaa-111", Properties: Zone{ZoneName: "example.com"}},
			{Id: "bbb-222", Properties: Zone{ZoneName: "example.org"}},
		},
	}

	// Returns a *ZoneRead pointer into the original slice.
	zone, ok := shared.FindByID(list.GetItems(), "bbb-222")
	if ok {
		props := zone.GetProperties()
		fmt.Println(props.GetZoneName())
	}
	// Output: example.org
}

// ExampleResource demonstrates that different SDK types (zones, records,
// or any other product) can be combined into a single collection through the
// shared interfaces and operated on uniformly.
func ExampleResource() {
	zones := []ZoneRead{
		{Id: "zone-aaa", Href: "/zones/zone-aaa"},
		{Id: "zone-bbb", Href: "/zones/zone-bbb"},
	}
	records := []RecordRead{
		{Id: "rec-111", Href: "/records/rec-111"},
		{Id: "rec-222", Href: "/records/rec-222"},
		{Id: "rec-333", Href: "/records/rec-333"},
	}

	// Both *ZoneRead and *RecordRead satisfy shared.Resource (GetId + GetHref),
	// so they can be combined into a single slice and operated on uniformly.
	var all []shared.Resource
	for i := range zones {
		all = append(all, &zones[i])
	}
	for i := range records {
		all = append(all, &records[i])
	}

	for _, r := range all {
		fmt.Printf("%s -> %s\n", r.GetId(), r.GetHref())
	}
	// Output:
	// zone-aaa -> /zones/zone-aaa
	// zone-bbb -> /zones/zone-bbb
	// rec-111 -> /records/rec-111
	// rec-222 -> /records/rec-222
	// rec-333 -> /records/rec-333
}
