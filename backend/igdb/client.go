package igdb

type IGDBService struct {
	WrapperService *IGDBWrapper
}

type IGDBClient interface {
	Test()
}

func NewIGDBClient(
	igdbWrapperService *IGDBWrapper,
) *IGDBService {
	return &IGDBService{
		WrapperService: igdbWrapperService,
	}
}

func (w *IGDBService) Test() {
	w.WrapperService.Test()
}

func isTokenExpired() bool {
	return false
}
