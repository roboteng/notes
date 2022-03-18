module Main exposing (..)

import Browser
import Css exposing (..)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (..)
import Html.Styled.Events exposing (..)


main : Program () Model Msg
main =
    Browser.sandbox
        { init = init
        , update = update
        , view = view >> toUnstyled
        }


type alias Model =
    { newNote : Maybe Note
    , notes : List Note
    }


type alias Note =
    { title : String
    , description : String
    }


type Msg
    = EditNote
    | SaveNote
    | UpdateTitle String
    | UpdateDescription String


init : Model
init =
    { newNote = Nothing
    , notes = []
    }


initNote : Note
initNote =
    { title = "", description = "" }


update : Msg -> Model -> Model
update msg model =
    case msg of
        EditNote ->
            { model
                | newNote = Just initNote
            }

        SaveNote ->
            case model.newNote of
                Nothing ->
                    model

                Just note ->
                    { model
                        | newNote = Nothing
                        , notes = model.notes ++ [ note ]
                    }

        UpdateTitle title ->
            case model.newNote of
                Nothing ->
                    model

                Just note ->
                    { model
                        | newNote = Just { note | title = title }
                    }

        UpdateDescription desc ->
            case model.newNote of
                Nothing ->
                    model

                Just note ->
                    { model
                        | newNote = Just { note | description = desc }
                    }



-- VIEW


view : Model -> Html Msg
view model =
    div
        []
        [ nav []
            [ h1 [] [ text "Notes" ]
            ]
        , main_ []
            [ button [ onClick EditNote ] [ text "Add New Note" ]
            , case model.newNote of
                Just note ->
                    newNoteForm note

                Nothing ->
                    div [] []
            , ul [] (List.map showNote model.notes)
            ]
        ]


newNoteForm : Note -> Html Msg
newNoteForm note =
    Html.Styled.form [ onSubmit SaveNote ]
        [ input [ value note.title, onInput UpdateTitle ] []
        , input [ value note.description, onInput UpdateDescription ] []
        , button [ onClick SaveNote ] [ text "Save" ]
        ]


showNote : Note -> Html Msg
showNote note =
    li [] [ text note.title ]
