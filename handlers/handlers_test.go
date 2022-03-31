package handlers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"notes/handlers"
	ts "notes/testHelpers"
	ty "notes/types"
	"testing"

	"github.com/julienschmidt/httprouter"
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
				notes := parse[[]ty.Note](res)
				ts.AssertEqualSlice(t, make([]ty.Note, 0), *notes)
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
				got := parse[[]ty.Note](res)
				want := []ty.Note{note}
				ts.AssertEqualSlice(t, want, *got)
			})
		})
	})
}

func TestCreateNote(t *testing.T) {
	t.Run("When an empty request comes in", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/notes", nil)
		r := httptest.NewRecorder()
		handler := handlers.CreateNote(&ty.AnonNoteCreator{
			Create: func(note ty.Note) (int, error) {
				panic("This should not be reached")
			},
		})
		handler(r, req, httprouter.Params{})
		res := r.Result()
		t.Run("Then the response code is 400", func(t *testing.T) {
			ts.AssertEquals(t, http.StatusBadRequest, res.StatusCode, "Status Code")
		})
	})
	t.Run("When a request comes in with title in the query", func(t *testing.T) {
		keys := url.Values{}
		keys.Add("title", "my title")
		query := keys.Encode()
		req := httptest.NewRequest("POST", "/api/notes?"+query, nil)
		r := httptest.NewRecorder()
		handler := handlers.CreateNote(&ty.AnonNoteCreator{
			Create: func(note ty.Note) (int, error) {
				return 1, nil
			},
		})
		handler(r, req, httprouter.Params{})
		res := r.Result()
		t.Run("Then the response is 201 - Created", func(t *testing.T) {
			ts.AssertEquals(t, http.StatusCreated, res.StatusCode, "Status Code")
		})
		t.Run("Then the body should be the new id", func(t *testing.T) {

			got := parse[struct {
				Id int `json:"id"`
			}](res)

			ts.AssertEquals(t, struct {
				Id int `json:"id"`
			}{1}, *got, "body")
		})
	})
}

func parse[T any](res *http.Response) *T {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	body := new(T)
	err = json.Unmarshal(b, body)
	if err != nil {
		panic(err)
	}
	return body
}

func getResponse(service *ty.AnonNotesViewer) *http.Response {
	wr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/notes", nil)
	handlers.GetNtes(service)(wr, r, httprouter.Params{})
	res := wr.Result()
	return res
}
