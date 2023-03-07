package igdb

type IGDBService struct {
	igdbWrapperService *IGDBWrapperService
}

type IGDBClient interface {
	Test()
}

func NewIGDBService(
	igdbWrapperService *IGDBWrapperService,
) *IGDBService {
	return &IGDBService{
		igdbWrapperService: igdbWrapperService,
	}
}

func (w *IGDBService) Test() {
	w.igdbWrapperService.Test()
}

func isTokenExpired() bool {
	return false
}
