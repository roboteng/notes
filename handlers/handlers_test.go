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
	givens := []struct {
		desc    string
		service ty.NotesViewer
	}{
		{
			desc: "Given a serive with no notes",
			service: &ty.AnonNotesViewer{
				View: func() []ty.Note {
					return make([]ty.Note, 0)
				},
			},
		},
	}
	for _, tC := range givens {
		t.Run(tC.desc, func(t *testing.T) {
			t.Run("When they go to /api/notes", func(t *testing.T) {
				router := handlers.MakeRouter()
				wr := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/api/notes", nil)
				router.ServeHTTP(wr, r)
				res := wr.Result()
				thens := []struct {
					desc      string
					want      []ty.Note
					transform func(http.Response) []ty.Note
				}{
					{
						desc: "Then they get an empty list",
						want: make([]ty.Note, 0),
						transform: func(r http.Response) []ty.Note {
							b, err := io.ReadAll(res.Body)
							if err != nil {
								t.Error(err.Error())
							}
							notes := make([]ty.Note, 0)
							err = json.Unmarshal(b, &notes)
							if err != nil {
								t.Error(err.Error())
							}
							return notes
						},
					},
				}
				for _, then := range thens {
					t.Run(then.desc, func(t *testing.T) {
						ts.AssertEqualSlice(t, then.want, then.transform(*res))
					})
				}
			})
		})
	}
}
