package stat

import (
	"fmt"
	"net/http"
	"time"

	"github.com/n-korel/shortcut-api/configs"
	"github.com/n-korel/shortcut-api/pkg/middleware"
)

const (
	FilterByDay = "day"
	FilterByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router * http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return 
		}

		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return 
		}

		by := r.URL.Query().Get("by")
		if by != FilterByDay && by != FilterByMonth {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return 
		}
		fmt.Println(from, to, by)
	}
}