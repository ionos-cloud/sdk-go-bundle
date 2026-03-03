package shared_test

import (
	"context"
	"fmt"

	"github.com/ionos-cloud/sdk-go-bundle/shared"
)

func ExampleResolve() {
	// Simulate: if input is already a UUID, no API call is made.
	id, _ := shared.Resolve(context.Background(), "550e8400-e29b-41d4-a716-446655440000",
		func(_ context.Context, _ string) ([]ZoneRead, error) {
			panic("should not be called")
		},
	)
	fmt.Println("UUID passthrough:", id)

	// Simulate: resolve a zone name to its ID via a list callback.
	id, _ = shared.Resolve(context.Background(), "example.com",
		func(_ context.Context, name string) ([]ZoneRead, error) {
			// In real code this would be:
			//   list, _, err := client.ZonesApi.ZonesGet(ctx).FilterZoneName(name).Limit(2).Execute()
			//   return list.GetItems(), err
			return []ZoneRead{
				{Id: "550e8400-e29b-41d4-a716-446655440000", Properties: Zone{ZoneName: name}},
			}, nil
		},
	)
	fmt.Println("Resolved:", id)
	// Output:
	// UUID passthrough: 550e8400-e29b-41d4-a716-446655440000
	// Resolved: 550e8400-e29b-41d4-a716-446655440000
}

func ExampleIsUUID() {
	fmt.Println(shared.IsUUID("550e8400-e29b-41d4-a716-446655440000"))
	fmt.Println(shared.IsUUID("example.com"))
	// Output:
	// true
	// false
}
