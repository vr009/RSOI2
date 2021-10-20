package delivery

import (
	"context"
	"errors"
	"fmt"
	"lib/services/models"
	"lib/services/proto/rating "
	"lib/services/rating/internal"
)

type GRPCHandler struct {
	usecase internal.RatingUsecase
}

func NewGRPCHandler(usecase internal.RatingUsecase) *GRPCHandler {
	return &GRPCHandler{usecase: usecase}
}

func (h *GRPCHandler) GetRating(ctx context.Context, req *rating.RatingRequest) (*rating.RatingResponse, error) {
	name := req.Name
	rate, status := h.usecase.GetRating(name)
	if status != models.OK {
		return nil, errors.New(fmt.Sprintf("%d", status))
	}

	return &rating.RatingResponse{Stars: int32(rate.Stars)} , nil
}


