package delphix

import (
	"fmt"
	"gopkg.in/resty.v1"
	"strings"
	"time"
	"encoding/json"
)

// CreateGroup creates a new group in Delphix
func (c *Client) CreateGroup(h *GroupStruct) (
	interface{}, error) {
	url := "/group"
	reference, err := c.executePostJobAndReturnObjectReference(url, h)

	return reference, err
}

// DeleteGroup deletes an group in Delphix
func (c *Client) DeleteGroup(r string) error {
	url := fmt.Sprintf("/group/%s/delete", strings.ToUpper(r))
	err := c.executePostJobAndReturnErrOnly(url, "{}")

	return err
}

// UpdateGroup updates an group in Delphix
func (c *Client) UpdateGroup(r string, h *GroupStruct) error {
	url := fmt.Sprintf("/group/%s", r)

	err := c.executePostJobAndReturnErrOnly(url, h)

	return err
}


// WaitforDelphixJob waits for a job to complete
func (c *Client) WaitforDelphixJob(j string) error {
	var jobState string
	var err error
	for jobState != "COMPLETED" && jobState != "FAILED" && jobState != "CANCELED" {
		time.Sleep(3 * time.Second)
		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(c.url + "/job/" + j)
		// explore response object
		if err != nil {
			panic(err)
		}
		s := resp.Body()

		var dat map[string]interface{}
		if err = json.Unmarshal(s, &dat); err != nil { //convert the json to go objects
			return err
		}
		results := dat["result"].(map[string]interface{}) //grab the query results
		jobState = results["jobState"].(string)
		fmt.Printf("\n%v", results["jobState"])
	}
	//If the job is failed or cancelled, return an error
	if jobState == "FAILED" {
		err = fmt.Errorf("Job Failed")
	} else if jobState == "CANCELED" {
		err = fmt.Errorf("Job Canceled")
	}
	return err
}
