package student

import (
	"log/slog"
	"net/http"
	"student_api/internal/types"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		
		slog.Info("student created successfully")
		
		w.Write([]byte("Hello World"))
	}
}
