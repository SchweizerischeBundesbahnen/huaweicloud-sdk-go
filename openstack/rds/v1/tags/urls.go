package tags

import (
    "github.com/gophercloud/gophercloud"
    "strings"
)

func getTagsURL(sc *gophercloud.ServiceClient, instanceId string) string {
    endpoint := sc.Endpoint
    // map: https://rds.eu-ch.o13bb.otc.t-systems.com/rds/v1/$(tenant_id)s
    // to: https://rds.eu-ch.o13bb.otc.t-systems.com/v1/$(tenant_id)s/rds
    sc.Endpoint = strings.Replace(endpoint, "rds/", "", 1)

	return sc.ServiceURL("rds", instanceId, "tags")
}
