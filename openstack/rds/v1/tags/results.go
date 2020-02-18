package tags

import (
	"github.com/gophercloud/gophercloud"
)

type GetTagsResult struct {
    gophercloud.Result
}

type Tag struct {
    Key   string
    Value string
}

// Extract will extract root user information from a UserRootResult.
func (r GetTagsResult) Extract() (map[string]string, error) {
    var s struct {
        Tags []Tag
    }
    err := r.ExtractInto(&s)
    if err != nil {
        return nil, err
    }
    m := make(map[string]string)
    for _, tag := range s.Tags {
        m[tag.Key] = tag.Value
    }
    return m, err
}
