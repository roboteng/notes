module Main exposing (..)

import Browser exposing (UrlRequest)
import Browser.Navigation as Nav
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (value)
import Html.Styled.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, field, int, list, map3, string)
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
    { page : Page
    , notes : List Note
    }


type Page
    = HomePage
    | EditNotePage Note
    | ViewNotePage Int (Maybe Note)


type alias Note =
    { title : String
    , description : String
    , id : Int
    }


type Msg
    = None
    | SaveNote
    | UpdateTitle String
    | UpdateDescription String
    | GotNotes (Result Http.Error (List Note))
    | GotNote (Result Http.Error Note)
    | PostedNote (Result Http.Error PostNoteResponse)
    | View Page
    | UrlChanged Url.Url


init : () -> Url.Url -> Nav.Key -> ( Model, Cmd Msg )
init _ _ _ =
    ( Model HomePage [], getNotes )


initNote : Note
initNote =
    Note "" "" 0


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none


onUrlChange : Url.Url -> Msg
onUrlChange url =
    case url.path of
        "/edit" ->
            View (EditNotePage initNote)

        "/1" ->
            View (ViewNotePage 1 Nothing)

        _ ->
            View HomePage


onUrlRequest : UrlRequest -> Msg
onUrlRequest _ =
    None


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        None ->
            ( model, Cmd.none )

        SaveNote ->
            case model.page of
                EditNotePage note ->
                    ( { model
                        | page = HomePage
                        , notes = model.notes ++ [ note ]
                      }
                    , postNote note
                    )

                _ ->
                    ( model, Cmd.none )

        UpdateTitle title ->
            case model.page of
                EditNotePage note ->
                    ( { model
                        | page = EditNotePage { note | title = title }
                      }
                    , Cmd.none
                    )

                _ ->
                    ( model, Cmd.none )

        UpdateDescription desc ->
            case model.page of
                EditNotePage note ->
                    ( { model
                        | page = EditNotePage { note | description = desc }
                      }
                    , Cmd.none
                    )

                _ ->
                    ( model, Cmd.none )

        GotNotes res ->
            case res of
                Ok ns ->
                    ( { model | notes = ns }, Cmd.none )

                Err _ ->
                    ( model, Cmd.none )

        GotNote res ->
            case res of
                Ok n ->
                    ( { model
                        | page = ViewNotePage n.id (Just n)
                      }
                    , Cmd.none
                    )

                Err _ ->
                    ( model, Cmd.none )

        PostedNote res ->
            case res of
                Ok _ ->
                    ( model, getNotes )

                Err _ ->
                    ( model, Cmd.none )

        View page ->
            ( { model
                | page = page
              }
            , case page of
                ViewNotePage id _ ->
                    getNote id

                _ ->
                    Cmd.none
            )

        UrlChanged _ ->
            ( model, Cmd.none )



-- HTTP


postNote : Note -> Cmd Msg
postNote note =
    Http.post
        { url = "/api/notes?title=" ++ note.title
        , body = Http.stringBody "text/plain" note.description
        , expect = Http.expectJson PostedNote idDecoder
        }


type alias PostNoteResponse =
    { id : Int }


idDecoder : Decoder PostNoteResponse
idDecoder =
    Json.Decode.map
        PostNoteResponse
        (field "id" int)


getNotes : Cmd Msg
getNotes =
    Http.get
        { url = "/api/notes"
        , expect = Http.expectJson GotNotes notesDecoder
        }


getNote : Int -> Cmd Msg
getNote id =
    Http.get
        { url = "/api/notes/" ++ String.fromInt id
        , expect = Http.expectJson GotNote noteDecoder
        }


notesDecoder : Decoder (List Note)
notesDecoder =
    list noteDecoder


noteDecoder : Decoder Note
noteDecoder =
    map3 Note
        (field "title" string)
        (field "desc" string)
        (field "id" int)



-- VIEW


view : Model -> Browser.Document Msg
view model =
    { title = "Notes"
    , body =
        [ toUnstyled
            (case model.page of
                ViewNotePage id n ->
                    div []
                        (case n of
                            Just note ->
                                [ p [] [ text (String.fromInt id) ]
                                , showNote note
                                ]

                            Nothing ->
                                [ p [] [ text (String.fromInt id) ]
                                , p [] [ text "Loading" ]
                                ]
                        )

                _ ->
                    myBody model
            )
        ]
    }


myBody : Model -> Html Msg
myBody model =
    div
        []
        [ nav []
            [ h1 [] [ text "Notes" ]
            ]
        , main_ []
            [ button [ onClick (View (EditNotePage initNote)) ] [ text "Add New Note" ]
            , case model.page of
                EditNotePage note ->
                    newNoteForm note

                ViewNotePage id _ ->
                    p [] [ text (String.fromInt id) ]

                HomePage ->
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
        , button [ onClick (View (ViewNotePage note.id Nothing)) ] [ text "View" ]
        ]
