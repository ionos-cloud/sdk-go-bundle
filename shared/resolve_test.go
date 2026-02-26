package shared

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsUUID(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"550e8400-e29b-41d4-a716-446655440000", true},
		{"550E8400-E29B-41D4-A716-446655440000", true}, // uppercase
		{"550e8400-e29b-41d4-a716-44665544000", false},  // too short
		{"example.com", false},
		{"my-resource", false},
		{"", false},
		{"not-a-uuid-at-all-but-has-dashes", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.want, IsUUID(tt.input))
		})
	}
}

func TestResolve_AlreadyUUID(t *testing.T) {
	uuid := "550e8400-e29b-41d4-a716-446655440000"

	// listByName should never be called when input is a UUID
	id, err := Resolve(context.Background(), uuid, func(_ context.Context, _ string) ([]mockResource, error) {
		t.Fatal("listByName should not be called for UUID input")
		return nil, nil
	})

	require.NoError(t, err)
	assert.Equal(t, uuid, id)
}

func TestResolve_SingleMatch(t *testing.T) {
	id, err := Resolve(context.Background(), "example.com", func(_ context.Context, name string) ([]mockResource, error) {
		assert.Equal(t, "example.com", name)
		return []mockResource{
			{id: "550e8400-e29b-41d4-a716-446655440000", name: "example.com"},
		}, nil
	})

	require.NoError(t, err)
	assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", id)
}

func TestResolve_NoMatch(t *testing.T) {
	_, err := Resolve(context.Background(), "nonexistent.com", func(_ context.Context, _ string) ([]mockResource, error) {
		return []mockResource{}, nil
	})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "no resource found")
	assert.Contains(t, err.Error(), "nonexistent.com")
}

func TestResolve_Ambiguous(t *testing.T) {
	_, err := Resolve(context.Background(), "common-name", func(_ context.Context, _ string) ([]mockResource, error) {
		return []mockResource{
			{id: "aaa-bbb-ccc-ddd"},
			{id: "eee-fff-ggg-hhh"},
		}, nil
	})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ambiguous")
	assert.Contains(t, err.Error(), "2 resources")
}

func TestResolve_ListError(t *testing.T) {
	_, err := Resolve(context.Background(), "example.com", func(_ context.Context, _ string) ([]mockResource, error) {
		return nil, fmt.Errorf("connection refused")
	})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "resolving")
	assert.Contains(t, err.Error(), "connection refused")
}
