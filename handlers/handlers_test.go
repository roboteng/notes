package handlers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"notes/handlers"
	ts "notes/testHelpers"
	ty "notes/types"
	"testing"
)

func TestGetNotes(t *testing.T) {
	t.Run("Given a serive with no notes", func(t *testing.T) {
		service := &ty.AnonNotesViewer{
			View: func() []ty.Note {
				return make([]ty.Note, 0)
			},
		}
		t.Run("When they go to /api/notes/", func(t *testing.T) {
			res := getResponse(service)
			t.Run("Then they get an empty list", func(t *testing.T) {
				notes := parseBody[[]ty.Note](t, res)
				ts.AssertEqualSlice(t, make([]ty.Note, 0), notes)
			})

			t.Run("Then the status should be 200", func(t *testing.T) {
				ts.AssertEquals(t, http.StatusOK, res.StatusCode, "status")
			})
		})
	})
	t.Run("Given the service has one note", func(t *testing.T) {
		note := ty.Note{
			Title:    "Note Title",
			Contents: "Note Contents",
		}
		service := &ty.AnonNotesViewer{
			View: func() []ty.Note {
				return []ty.Note{note}
			},
		}
		t.Run("When they go to /api/notes", func(t *testing.T) {
			res := getResponse(service)
			t.Run("Then the response should have that note", func(t *testing.T) {
				got := parseBody[[]ty.Note](t, res)
				want := []ty.Note{note}
				ts.AssertEqualSlice(t, want, got)
			})
		})
	})
}

func parseBody[T any](t *testing.T, res *http.Response) T {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err.Error())
	}
	body := new(T)
	err = json.Unmarshal(b, body)
	if err != nil {
		t.Error(err.Error())
	}
	return *body
}

func getResponse(service *ty.AnonNotesViewer) *http.Response {
	router := handlers.MakeRouter(service)
	wr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/notes", nil)
	router.ServeHTTP(wr, r)
	res := wr.Result()
	return res
}
