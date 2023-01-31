/*
 * Container Registry service
 *
 * Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls.
 *
 * API version: 1.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerregistry

import (
	"github.com/ionos-cloud/sdk-go-bundle/common"
	"strings"
	"time"
)

type IonosTime struct {
	time.Time
}

func (t *IonosTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	if common.Strlen(str) == 0 {
		t = nil
		return nil
	}
	if str[0] == '"' {
		str = str[1:]
	}
	if str[len(str)-1] == '"' {
		str = str[:len(str)-1]
	}
	if !strings.Contains(str, "Z") {
		/* forcefully adding timezone suffix to be able to parse the
		 * string using RFC3339 */
		str += "Z"
	}
	tt, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}
	*t = IonosTime{tt}
	return nil
}