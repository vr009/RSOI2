package delivery

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	models2 "lib/services/models"
	"lib/services/proto/library"
	"lib/services/proto/rating "
	"lib/services/proto/reservation"
)

type Client struct {
	RatingServiceClient      rating.RatingServiceClient
	ReservationServiceClient reservation.ReservationServiceClient
	LibraryServiceClient     library.LibraryServiceClient
}

func NewGRPCClient (Rating, Reservation, Library grpc.ClientConnInterface) *Client {
	return &Client{
		RatingServiceClient: rating.NewRatingServiceClient(Rating),
		ReservationServiceClient: reservation.NewReservationServiceClient(Reservation),
		LibraryServiceClient: library.NewLibraryServiceClient(Library),
	}
}

func (cl * Client) 	GetLibraries(page, size int64, city string) ([]models2.LibraryPaginationResponse, models2.StatusCode) {
	request := &library.LibraryRequest{Page: page, Size: size, City: city}
	response, err := cl.LibraryServiceClient.FetchLibs(context.Background(), request,)
	if err != nil {
		return nil, models2.InternalError
	}
	libs := []models2.LibraryPaginationResponse{}
	for _, el := range response.Items {
		libEl := models2.LibraryPaginationResponse{}
		libEl.TotalElements = el.TotalElements
		libEl.Page = el.Page
		libEl.PageSize = el.Size
		for _, item := range el.Item {
			uid, err := uuid.Parse(item.LibraryUid)
			if err != nil {
				return nil, models2.BadRequest
			}
			libEl.Items = append(libEl.Items, models2.LibraryResponse{
				LibraryUid: uid,
				Name: item.Name,
				Address: item.Address,
				City: item.City,
			})
		}
	}
	return libs, models2.OK
}
func (cl * Client) GetBooks(page, size int64, showAll bool, LibUid uuid.UUID) ([]models2.LibraryBookPaginationResponse, models2.StatusCode) {
	request := &library.BookRequest{Page: page, Size: size, ShowAll: showAll}
	response, err := cl.LibraryServiceClient.FetchBooks(context.Background(), request)
	if err != nil {
		return nil, models2.InternalError
	}
	libs := []models2.LibraryBookPaginationResponse{}
	for _, el := range response.Items {
		libEl := models2.LibraryBookPaginationResponse{}
		libEl.TotalElements = el.TotalElements
		libEl.Page = el.Page
		libEl.PageSize = el.Size
		for _, item := range el.Item {
			uid, err := uuid.Parse(item.BookUid)
			if err != nil {
				return nil, models2.BadRequest
			}
			libEl.Items = append(libEl.Items, models2.LibraryBookResponse{
				BookId: uid,
				Name: item.Name,
				Author: item.Author,
				Genre: item.Genre,
				Condition: models2.BookCondition(item.Condition.String()),
				AvailableCount: item.AvailableCount,
			})
		}
	}
	return libs, models2.OK
}
func (cl * Client) GetReservations() ([]models2.BookReservationResponse, models2.StatusCode) {

	return nil, 0
}
func (cl * Client) TakeBook() (models2.TakeBookResponse, models2.StatusCode) {
	return models2.TakeBookResponse{}, 0
}
func (cl * Client) ReturnBook() models2.StatusCode {
	return 0
}
func (cl * Client) GetRating() (models2.UserRatingResponse, models2.StatusCode) {
	return models2.UserRatingResponse{}, 0
}