package libsubscription

import (
	"encoding/base64"
	"strings"

	E "github.com/sagernet/sing/common/exceptions"
)

func ParseRawSubscription(content string) ([]Server, error) {
	if base64Content, err := decodeBase64URLSafe(content); err == nil {
		servers, _ := parseRawSubscription(base64Content)
		if len(servers) > 0 {
			return servers, err
		}
	}
	return parseRawSubscription(content)
}

func parseRawSubscription(content string) ([]Server, error) {
	var servers []Server
	content = strings.ReplaceAll(content, "\r\n", "\n")
	linkList := strings.Split(content, "\n")
	for _, linkLine := range linkList {
		if server, err := ParseSubscriptionLink(linkLine); err == nil {
			servers = append(servers, server)
		}
	}
	if len(servers) == 0 {
		return nil, E.New("no servers found")
	}
	return servers, nil
}

func decodeBase64URLSafe(content string) (string, error) {
	content = strings.ReplaceAll(content, " ", "-")
	content = strings.ReplaceAll(content, "/", "_")
	content = strings.ReplaceAll(content, "+", "-")
	content = strings.ReplaceAll(content, "=", "")
	result, err := base64.StdEncoding.DecodeString(content)
	return string(result), err
}
