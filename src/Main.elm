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
    { newNote : Maybe NewNoteForm }


type alias NewNoteForm =
    { title : String
    , description : String
    }


type Msg
    = EditNote
    | SaveNote


init : Model
init =
    { newNote = Nothing
    }


update : Msg -> Model -> Model
update msg model =
    case msg of
        EditNote ->
            { model
                | newNote = Just { title = "", description = "" }
            }

        SaveNote ->
            { model | newNote = Nothing }



-- VIEW


view : Model -> Html Msg
view model =
    div
        [ css
            [ backgroundColor (rgb 100 100 100), color (rgb 200 200 200) ]
        ]
        [ nav []
            [ h1 [] [ text "Notes" ]
            ]
        , main_ []
            [ button [ onClick EditNote ] [ text "Add New Note" ]
            , case model.newNote of
                Just note ->
                    Html.Styled.form [ onSubmit SaveNote ]
                        [ input [ value note.title ] []
                        , input [ value note.description ] []
                        , button [ onClick SaveNote ] [ text "Save" ]
                        ]

                Nothing ->
                    div [] []
            ]
        ]
