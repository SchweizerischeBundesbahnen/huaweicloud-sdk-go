package tags

import (
	"github.com/gophercloud/gophercloud"
)


func GetTags(client *gophercloud.ServiceClient, instanceId string) (r GetTagsResult) {
    _, r.Err = client.Get(getTagsURL(client, instanceId), &r.Body, nil)
    return
}
