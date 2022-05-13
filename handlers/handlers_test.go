package handlers_test

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"notes/features"
	"notes/handlers"
	"notes/services"
	ts "notes/testHelpers"
	ty "notes/types"
	"strings"
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
		feature := &features.ViewNotes{Service: service}
		t.Run("When they go to /api/notes/", func(t *testing.T) {
			res := getResponse(feature)
			t.Run("Then they get an empty list", func(t *testing.T) {
				notes := parse[[]ty.Note](res)
				ts.AssertEqualSlice(t, make([]ty.Note, 0), *notes)
			})

			t.Run("Then the status should be 200", func(t *testing.T) {
				ts.AssertEquals(t, http.StatusOK, res.StatusCode)
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
		feature := &features.ViewNotes{Service: service}
		t.Run("When they go to /api/notes", func(t *testing.T) {
			res := getResponse(feature)
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
			ts.AssertEquals(t, http.StatusBadRequest, res.StatusCode)
		})
	})
	t.Run("When a request comes in with title in the query, and desc in the body", func(t *testing.T) {
		keys := url.Values{}
		keys.Add("title", "my title")
		query := keys.Encode()
		req := httptest.NewRequest("POST", "/api/notes?"+query, strings.NewReader("my desc"))
		r := httptest.NewRecorder()
		service := &services.InMemoryNoteService{}
		handler := handlers.CreateNote(service)
		handler(r, req, httprouter.Params{})
		res := r.Result()
		t.Run("Then the response is 201 - Created", func(t *testing.T) {
			ts.AssertEquals(t, http.StatusCreated, res.StatusCode)
		})
		t.Run("Then the body should be the new id", func(t *testing.T) {
			got := parse[ty.CreateNoteResponse](res)
			ts.AssertEquals(t, ty.CreateNoteResponse{Id: 1}, *got)
		})
		t.Run("Then the service should have a note added", func(t *testing.T) {
			ts.AssertEquals(t, "my title", service.ViewNotes()[0].Title)
			ts.AssertEquals(t, "my desc", service.ViewNotes()[0].Contents)
		})
	})
	t.Run("When two create note requests come in, they should have the correct ids", func(t *testing.T) {
		keys := url.Values{}
		keys.Add("title", "my title")
		query := keys.Encode()
		req := httptest.NewRequest("POST", "/api/notes?"+query, nil)
		r := httptest.NewRecorder()
		handler := handlers.CreateNote(&inMemoryNoteCreator{})
		handler(httptest.NewRecorder(), httptest.NewRequest("POST", "/api/notes?"+query, nil), httprouter.Params{})
		handler(r, req, httprouter.Params{})
		res := r.Result()
		t.Run("The second response should have an id of 2", func(t *testing.T) {
			got := parse[ty.CreateNoteResponse](res)

			ts.AssertEquals(t, ty.CreateNoteResponse{Id: 2}, *got)
		})
	})
	t.Run("When a valid reponse comes in, and when the service gives an error", func(t *testing.T) {
		keys := url.Values{}
		keys.Add("title", "my title")
		query := keys.Encode()
		req := httptest.NewRequest("POST", "/api/notes?"+query, nil)
		r := httptest.NewRecorder()
		service := &ty.AnonNoteCreator{
			Create: func(note ty.Note) (int, error) {
				return 0, errors.New("It Failed!")
			},
		}
		handlers.CreateNote(service)(r, req, httprouter.Params{})
		res := r.Result()
		t.Run("Then the response should have a 500 status code", func(t *testing.T) {
			ts.AssertEquals(t, http.StatusInternalServerError, res.StatusCode)
		})
	})
}

func TestViewSingleNote(t *testing.T) {
	t.Run("Given a service that has one note", func(t *testing.T) {
		note := ty.Note{
			Title:    "My Title",
			Contents: "My Contents",
			Id:       1,
		}
		service := &ty.AnonSingleNoteViewer{
			View: func(id int) (ty.Note, error) {
				if id == 1 {
					return note, nil
				}
				return ty.Note{}, errors.New("No note found at that ID")
			},
		}
		handler := handlers.GetSingleNote(service)
		t.Run("When a request comes in for id 1", func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/api/notes/1", nil)
			w := httptest.NewRecorder()
			handler(w, r, httprouter.Params{
				httprouter.Param{
					Key:   "id",
					Value: "1",
				},
			})
			res := w.Result()
			t.Run("then the status code should be 200", func(t *testing.T) {
				ts.AssertEquals(t, http.StatusOK, res.StatusCode)
			})
			t.Run("Then the body should contain the json for the note", func(t *testing.T) {
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					panic(err)
				}
				got := ty.Note{}
				err = json.Unmarshal(body, &got)
				if err != nil {
					panic(err)
				}
				ts.AssertEquals(t, note, got)
			})
		})
		t.Run("When a request comes in for a different ID", func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/api/notes/2", nil)
			w := httptest.NewRecorder()
			handler(w, r, httprouter.Params{
				httprouter.Param{
					Key:   "id",
					Value: "2",
				},
			})
			res := w.Result()
			t.Run("Then the status should be 404", func(t *testing.T) {
				ts.AssertEquals(t, http.StatusNotFound, res.StatusCode)
			})
		})
	})
}

type inMemoryNoteCreator struct {
	notes int
}

func (i *inMemoryNoteCreator) CreateNote(note ty.Note) (int, error) {
	i.notes++
	return i.notes, nil
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

func getResponse(feature *features.ViewNotes) *http.Response {
	wr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/notes", nil)
	handlers.GetNtes(feature)(wr, r, httprouter.Params{})
	res := wr.Result()
	return res
}
