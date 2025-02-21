package igdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"vgtracker/backend/internal/twitch"

	"github.com/pkg/errors"
)

type IgdbMethods interface {
	Search(input string) (*VGTSearchResults, error)
}

type Client struct {
	AccessToken *twitch.AccessTokenResponse
	ClientID    *string
}

func NewClient(token *twitch.AccessTokenResponse, clientID *string) IgdbMethods {
	return &Client{
		AccessToken: token,
		ClientID:    clientID,
	}
}

// Search implements IgdbMethods.
func (c *Client) Search(input string) (*VGTSearchResults, error) {
	client := &http.Client{}

	postURL := "https://api.igdb.com/v4/games"
	body := []byte(`fields *;where id=5;`)
	request, err := http.NewRequest("POST", postURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed making request")
	}

	request.Header.Add("Client-ID", *c.ClientID)
	token := fmt.Sprintf("Bearer %s", c.AccessToken.AccessToken)
	request.Header.Add("Authorization", token)

	response, err := client.Do(request)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed the request")
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to read all of request")
	}

	// data := &SearchResult{}
	results := &VGTSearchResults{}
	var data []IGDBSearchResultGame
	derr := json.Unmarshal(responseBody, &data)
	if derr != nil {
		return nil, errors.WithMessage(derr, "Failed to unmarshall")
	}

	for _, item := range data {
		results.Items = append(results.Items, VGTGame{
			Name:     item.Name,
			GameType: item.convertGameType(item.GameType),
			Genres:   item.convertGenre(item),
		})
	}

	fmt.Printf("\n data result: %+v \n", data)

	return results, nil
}
