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
    Model Nothing []


initNote : Note
initNote =
    Note "" ""


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
        [ formInput "Title" note.title UpdateTitle
        , formInput "Description" note.description UpdateDescription
        , button [ onClick SaveNote ] [ text "Save" ]
        ]


formInput : String -> String -> (String -> Msg) -> Html Msg
formInput l v action =
    label []
        [ text l
        , input [ value v, onInput action ] []
        ]


showNote : Note -> Html Msg
showNote note =
    li []
        [ p [] [ text note.title ]
        , p [] [ text note.description ]
        ]
