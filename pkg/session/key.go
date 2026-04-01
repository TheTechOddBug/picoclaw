package session

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

const sessionKeyV1Prefix = "sk_v1_"

// BuildOpaqueSessionKey returns a stable opaque session key derived from a
// canonical alias string. The alias remains available through metadata for
// compatibility and migration purposes.
func BuildOpaqueSessionKey(alias string) string {
	normalized := strings.TrimSpace(strings.ToLower(alias))
	if normalized == "" {
		return ""
	}
	sum := sha256.Sum256([]byte(normalized))
	return sessionKeyV1Prefix + hex.EncodeToString(sum[:])
}

// IsOpaqueSessionKey returns true when the key matches the current opaque
// session-key format.
func IsOpaqueSessionKey(key string) bool {
	return strings.HasPrefix(strings.ToLower(strings.TrimSpace(key)), sessionKeyV1Prefix)
}

// CanonicalScopeSignature returns a stable serialized representation of scope.
func CanonicalScopeSignature(scope SessionScope) string {
	parts := []string{
		fmt.Sprintf("v=%d", scope.Version),
		fmt.Sprintf("agent=%s", strings.TrimSpace(strings.ToLower(scope.AgentID))),
		fmt.Sprintf("channel=%s", strings.TrimSpace(strings.ToLower(scope.Channel))),
		fmt.Sprintf("account=%s", strings.TrimSpace(strings.ToLower(scope.Account))),
	}
	for _, dimension := range scope.Dimensions {
		dimension = strings.TrimSpace(strings.ToLower(dimension))
		if dimension == "" {
			continue
		}
		value := strings.TrimSpace(strings.ToLower(scope.Values[dimension]))
		parts = append(parts, fmt.Sprintf("%s=%s", dimension, value))
	}
	return strings.Join(parts, "|")
}

// BuildSessionKey returns the current opaque key for a structured session scope.
func BuildSessionKey(scope SessionScope) string {
	return BuildOpaqueSessionKey(CanonicalScopeSignature(scope))
}
