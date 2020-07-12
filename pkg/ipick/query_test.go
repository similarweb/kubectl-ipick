package ipick

import (
	"testing"
)

func TestNewQueryOptions(t *testing.T) {

	namespace := "test-namespace"
	queryOption := NewQueryOptions(namespace)

	if *queryOption.configFlags.Context != namespace {
		t.Fatalf("unexpected query context option, got %s expected %s", *queryOption.configFlags.Context, namespace)
	}
}
