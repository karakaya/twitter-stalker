package twitter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type UserFollower struct {
	Data []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
	Meta struct {
		ResultCount int `json:"result_count"`
	} `json:"meta"`
}
type Tw struct {
	ID int64
}

var bearer string

func init() {
	if bearer = os.Getenv("Bearer"); bearer == "" {
		panic("bearer token is missing")
	}

}
func (t *Tw) Followers() UserFollower {
	url := fmt.Sprintf("https://api.twitter.com/2/users/%d/followers", t.ID)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Authorization", "Bearer "+bearer)

	response, err := client.Do(request)

	if err != nil {
		panic(err)

	}
	defer response.Body.Close()

	var follower UserFollower

	json.NewDecoder(response.Body).Decode(&follower)
	return follower
}
