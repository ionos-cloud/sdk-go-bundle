package shared_test

import (
	"fmt"

	"github.com/ionos-cloud/sdk-go-bundle/shared"
)

// ──────────────────────────────────────────────────────────────────────────────
// Mock types that mirror the generated SDK model shape (e.g. dns.ZoneRead,
// dns.Zone, dns.ZoneReadList). All generated SDK types follow this exact
// pattern: pointer receivers, value-typed Items slices, nested Properties.
// ──────────────────────────────────────────────────────────────────────────────

// Zone mirrors a properties type like dns.Zone.
type Zone struct {
	ZoneName string
}

func (o *Zone) GetZoneName() string {
	if o == nil {
		return ""
	}
	return o.ZoneName
}

// ZoneRead mirrors a read-model type like dns.ZoneRead.
// Pointer receivers — exactly like the generated code.
type ZoneRead struct {
	Id         string
	Href       string
	Properties Zone
}

func (o *ZoneRead) GetId() string {
	if o == nil {
		return ""
	}
	return o.Id
}

func (o *ZoneRead) GetHref() string {
	if o == nil {
		return ""
	}
	return o.Href
}

func (o *ZoneRead) GetProperties() Zone {
	if o == nil {
		return Zone{}
	}
	return o.Properties
}

// ZoneReadList mirrors a list-response type like dns.ZoneReadList.
// GetItems() returns []ZoneRead (values, not pointers) — exactly like
// the generated code.
type ZoneReadList struct {
	Items []ZoneRead
}

func (o *ZoneReadList) GetItems() []ZoneRead {
	if o == nil {
		return nil
	}
	return o.Items
}

// Record mirrors a properties type like dns.Record.
type Record struct {
	Name    string
	Type    string
	Content string
}

func (o *Record) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// RecordRead mirrors a read-model type like dns.RecordRead.
// Same pattern as ZoneRead: pointer receivers, same interface shape.
type RecordRead struct {
	Id         string
	Href       string
	Properties Record
}

func (o *RecordRead) GetId() string {
	if o == nil {
		return ""
	}
	return o.Id
}

func (o *RecordRead) GetHref() string {
	if o == nil {
		return ""
	}
	return o.Href
}

// ──────────────────────────────────────────────────────────────────────────────
// Examples — these appear in godoc and are executed by `go test`.
// ──────────────────────────────────────────────────────────────────────────────

func ExampleExtractIDs() {
	// Simulate an API list response:
	//   list, _, err := client.ZonesApi.ZonesGet(ctx).Execute()
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
	// Simulate an API list response:
	//   list, _, err := client.ZonesApi.ZonesGet(ctx).Execute()
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

func ExampleListItems() {
	// Simulate an API list response:
	//   list, _, err := client.ZonesApi.ZonesGet(ctx).Execute()
	list := &ZoneReadList{
		Items: []ZoneRead{
			{Id: "aaa-111", Properties: Zone{ZoneName: "example.com"}},
			{Id: "bbb-222", Properties: Zone{ZoneName: "example.org"}},
		},
	}

	// ListItems is useful in generic contexts where the list type is a
	// type parameter. For direct use, list.GetItems() works just as well.
	items := shared.ListItems[ZoneRead](list)
	fmt.Println(len(items))
	// Output: 2
}

func ExampleProperties() {
	zone := &ZoneRead{
		Id:         "aaa-111",
		Properties: Zone{ZoneName: "example.com"},
	}

	// Properties is useful in generic contexts where both the model and
	// properties types are type parameters.
	props := shared.Properties[Zone](zone)
	fmt.Println(props.GetZoneName())
	// Output: example.com
}

// Example_interfaces demonstrates that different SDK types (zones, records,
// or any other product) can be combined into a single collection through the
// shared interfaces and operated on uniformly.
func Example_interfaces() {
	// Simulate fetching zones and records from different API calls.
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
