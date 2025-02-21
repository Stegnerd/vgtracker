package controllers

import (
	"vgtracker/backend/internal/igdb"

	"github.com/pkg/errors"
)

type IgdbControllerMethods interface {
	Search(input string) (*igdb.VGTSearchResults, error)
}

type IgdbController struct {
	client igdb.IgdbMethods
}

func NewIgdbController(
	client igdb.IgdbMethods,
) IgdbControllerMethods {
	return &IgdbController{
		client: client,
	}
}

// Search implements IgdbControllerMethods.
func (i *IgdbController) Search(input string) (*igdb.VGTSearchResults, error) {
	result, err := i.client.Search(input)
	if err != nil {
		return nil, errors.WithMessage(err, "could not complete search")
	}

	return result, nil
}
