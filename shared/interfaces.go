package shared

// Identifiable is satisfied by any SDK model type that exposes a string ID.
// All *Read types across SDK products (dns.ZoneRead, compute.Datacenter,
// vpn.WireguardGatewayRead, etc.) implement this interface via their
// generated GetId() method.
type Identifiable interface {
	GetId() string
}

// HasHref is satisfied by any SDK model type that exposes an href link.
type HasHref interface {
	GetHref() string
}

// Resource combines Identifiable and HasHref — the two universally
// consistent accessor methods across all SDK product model types.
type Resource interface {
	Identifiable
	HasHref
}

// Listable is satisfied by SDK list response types that expose their
// items via GetItems(). Most generated list types (dns.ZoneReadList,
// vpn.IPSecGatewayReadList, compute.Datacenters, etc.) satisfy this.
type Listable[T any] interface {
	GetItems() []T
}

// HasProperties is satisfied by SDK model types that wrap a properties
// sub-object (e.g., dns.ZoneRead exposes GetProperties() dns.Zone).
type HasProperties[P any] interface {
	GetProperties() P
}

// ExtractIDs returns the IDs from a slice of Identifiable items.
func ExtractIDs[T Identifiable](items []T) []string {
	ids := make([]string, len(items))
	for i, item := range items {
		ids[i] = item.GetId()
	}
	return ids
}

// FindByID returns the first item matching the given ID, and true if found.
func FindByID[T Identifiable](items []T, id string) (T, bool) {
	for _, item := range items {
		if item.GetId() == id {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// ListItems extracts items from a Listable response. This is a convenience
// wrapper useful in generic contexts where the list type is a type parameter.
func ListItems[T any, L Listable[T]](list L) []T {
	return list.GetItems()
}

// Properties extracts the properties from a HasProperties type. Useful in
// generic contexts where both the model and properties types are parameters.
func Properties[P any, T HasProperties[P]](item T) P {
	return item.GetProperties()
}
