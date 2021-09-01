package delphix

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"log"
	"net/http"
)

// FindGroupByName returns the group object (interface) of the named (n) group
//func (c *Client) FindGroupByName(n string) (interface{}, error) {
//	obj, err := c.FindObjectByName("/group", n)
//	return obj, err
//}

// FindGroupByName returns the environment object (interface) of the named (n) environment
func (c *Client) FindGroupByName(n string) (interface{}, error) {
	var err error
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(c.url + "/group") //grab all the groups
	if err != nil {
		return nil, err
	}

	if http.StatusOK != resp.StatusCode() { //check to make sure our query was good
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return nil, err
		}
	}

	groups := resp.Body()

	var dat map[string]interface{}
	if err = json.Unmarshal(groups, &dat); err != nil { //convert the json to go objects
		return nil, err
	}
	results := dat["result"].([]interface{}) //grab the query results
	for _, result := range results {         //loop through the groups
		name := result.(map[string]interface{})["name"] //grab the group name
		if n == name {                                  //if the name matches our specified group
			return result, nil //return the group object
		}
	}
	log.Printf("Was unable to find group \"%s\" in %s", n, string(resp.Body()))
	return nil, nil
}


// FindGroupByRef returns the group object (interface) of the named (n) group
func (c *Client) FindGroupByRef(n string) (interface{}, error) {
	log.Println("------n:" + n)
	obj, err := c.FindObjectByReference("/group", n)
	return obj, err
}

// FindGroupRefByName returns the group reference (string) of the named (n) group
func (c *Client) FindGroupRefByName(n string) (interface{}, error) {
	obj, err := c.FindGroupByName(n)
	return obj.(map[string]interface{})["reference"].(string), err
}

// FindObjectByReference returns the referenced (r) object type (u)
func (c *Client) FindObjectByReference(u string, r string) (interface{}, error) {
	url := fmt.Sprintf("%s/%s", u, r)
	obj, err := c.executeReadAndReturnObject(url)
	if err != nil {
		return nil, err
	}
	return obj, err
}


