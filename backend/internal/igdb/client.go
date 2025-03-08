package igdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"vgtracker/backend/internal/twitch"

	"github.com/pkg/errors"
)

type IgdbMethods interface {
	SearchMainGames(input string) (*VGTSearchResults, error)
}

type Client struct {
	AccessToken   *twitch.AccessTokenResponse
	ClientID      *string
	ImageResolver ImageResolverMethods
}

func NewClient(token *twitch.AccessTokenResponse, clientID *string) IgdbMethods {
	newResolver := NewImageResolver()

	return &Client{
		AccessToken:   token,
		ClientID:      clientID,
		ImageResolver: newResolver,
	}
}

// Search implements IgdbMethods.
func (c *Client) SearchMainGames(input string) (*VGTSearchResults, error) {
	client := &http.Client{}

	postURL := "https://api.igdb.com/v4/games"
	// currently searching infix naming (case insenstive contains string)
	// and only main/remake/remaster
	bodyString := fmt.Sprintf(`
		fields id,
		age_ratings.*,
		artworks.*,
		category,
		cover.*,
		first_release_date,
		game_type,
		genres.*,
		name,
		platforms.*,
		summary;
		sort name asc;
		where
			name ~ *"%s"*  & game_type = 0 |
			name ~ *"%s"*  & game_type = 8 |
			name ~ *"%s"*  & game_type = 9 ;
		limit 30;
	`, input, input, input)
	body := []byte(bodyString)
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

	//fmt.Printf("data %+v", data)

	for _, item := range data {
		timeStamp := time.Unix(item.FirstReleaseDate, 0)
		results.Items = append(results.Items, VGTGame{
			GameID:           item.ID,
			Name:             item.Name,
			GameType:         item.convertGameType(item.GameType),
			Genres:           item.getGenreList(item.Genres),
			Platforms:        item.getPlatformList(item.Platforms),
			ThumbnailURL:     c.ImageResolver.GetImageURL(Thumb, item.Cover.ImageID),
			CoverURL:         c.ImageResolver.GetImageURL(CoverBig, item.Cover.ImageID),
			FirstReleaseYear: timeStamp.Year(),
			Summary:          item.Summary,
		})
	}

	// fmt.Printf("\n data result: %+v \n", results)

	return results, nil
}
