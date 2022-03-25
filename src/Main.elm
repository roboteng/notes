module Main exposing (..)

import Browser exposing (UrlRequest)
import Css exposing (..)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (value)
import Html.Styled.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, field, list, map2, string)
import Url


main : Program () Model Msg
main =
    Browser.application
        { init = init
        , update = update
        , view = view
        , subscriptions = subscriptions
        , onUrlChange = onUrlChange
        , onUrlRequest = onUrlRequest
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
    = None
    | EditNote
    | SaveNote
    | UpdateTitle String
    | UpdateDescription String
    | GotNotes (Result Http.Error (List Note))


init : () -> a -> b -> ( Model, Cmd Msg )
init _ _ _ =
    ( Model Nothing [], getNotes )


initNote : Note
initNote =
    Note "" ""


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none


onUrlChange : Url.Url -> Msg
onUrlChange _ =
    None


onUrlRequest : UrlRequest -> Msg
onUrlRequest _ =
    None


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        None ->
            ( model, Cmd.none )

        EditNote ->
            ( { model
                | newNote = Just initNote
              }
            , Cmd.none
            )

        SaveNote ->
            case model.newNote of
                Nothing ->
                    ( model, Cmd.none )

                Just note ->
                    ( { model
                        | newNote = Nothing
                        , notes = model.notes ++ [ note ]
                      }
                    , Cmd.none
                    )

        UpdateTitle title ->
            case model.newNote of
                Nothing ->
                    ( model, Cmd.none )

                Just note ->
                    ( { model
                        | newNote = Just { note | title = title }
                      }
                    , Cmd.none
                    )

        UpdateDescription desc ->
            case model.newNote of
                Nothing ->
                    ( model, Cmd.none )

                Just note ->
                    ( { model
                        | newNote = Just { note | description = desc }
                      }
                    , Cmd.none
                    )

        GotNotes res ->
            case res of
                Ok ns ->
                    ( { model | notes = ns }, Cmd.none )

                Err _ ->
                    ( model, Cmd.none )



-- HTTP


getNotes : Cmd Msg
getNotes =
    Http.get
        { url = "/api/notes"
        , expect = Http.expectJson GotNotes notesDecoder
        }


notesDecoder : Decoder (List Note)
notesDecoder =
    list noteDecoder


noteDecoder : Decoder Note
noteDecoder =
    map2 Note
        (field "title" string)
        (field "desc" string)



-- VIEW


view : Model -> Browser.Document Msg
view model =
    { title = "Notes"
    , body = [ toUnstyled (myBody model) ]
    }


myBody : Model -> Html Msg
myBody model =
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
