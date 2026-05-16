package interviewtransport

import (
	"net/http"

	interviewserver "gateway-service/server/interviewserver"
	commonhttp "gateway-service/transport/common"
)

type InterviewHTTPHandlers struct {
	service interviewserver.InterviewServiceInterface
}

func NewInterviewHTTPHandlers(svc interviewserver.InterviewServiceInterface) *InterviewHTTPHandlers {
	return &InterviewHTTPHandlers{
		service: svc,
	}
}

func (h *InterviewHTTPHandlers) GetInterviews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		commonhttp.RespondWithError(w, http.StatusUnauthorized, "user_id not found in context")
		return
	}

	resp, err := h.service.GetInterviews(ctx, userID)
	if err != nil {
		commonhttp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	commonhttp.RespondWithJSON(w, http.StatusOK, resp)
}
