package util

import "strings"

func SanitizeSearchQuery(query string) string {
	// Remove SQL wildcards
	query = strings.ReplaceAll(query, "%", "")
	query = strings.ReplaceAll(query, "_", "")
	// Remove quotes
	query = strings.ReplaceAll(query, "'", "")
	query = strings.ReplaceAll(query, "\"", "")
	
	return query
}