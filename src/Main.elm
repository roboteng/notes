module Main exposing (..)

import Browser
import Browser.Navigation as Nav
import Css exposing (border3, cursor, listStyleType, none, pointer, px, rgb, solid)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, value)
import Html.Styled.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, field, int, list, map3, string)
import List
import Url
import Url.Builder


main : Program () Model Msg
main =
    Browser.application
        { init = init
        , update = update
        , view = view
        , subscriptions = subscriptions
        , onUrlChange = UrlChanged
        , onUrlRequest = LinkClicked
        }


type alias Model =
    { page : Page
    , notes : List Note
    , key : Nav.Key
    , url : Url.Url
    , message : String
    }


type Page
    = HomePage
    | EditNotePage Note
    | ViewNotePage Int (Maybe (Result String Note))


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
    | LinkClicked Browser.UrlRequest


init : () -> Url.Url -> Nav.Key -> ( Model, Cmd Msg )
init _ url key =
    let
        paths_ =
            String.split "/" url.path |> List.tail
    in
    case paths_ of
        Just paths ->
            case List.head paths of
                Just "note" ->
                    ( { page = ViewNotePage 1 Nothing
                      , notes = []
                      , key = key
                      , url = url
                      , message = ""
                      }
                    , getNote 1
                    )

                Just "edit" ->
                    ( { page = EditNotePage initNote
                      , notes = []
                      , key = key
                      , url = url
                      , message = ""
                      }
                    , Cmd.none
                    )

                _ ->
                    ( { page = HomePage
                      , notes = []
                      , key = key
                      , url = url
                      , message = ""
                      }
                    , getNotes
                    )

        Nothing ->
            ( { page = HomePage
              , notes = []
              , key = key
              , url = url
              , message = ""
              }
            , getNotes
            )


parseNoteId : List String -> Maybe Int
parseNoteId url =
    List.drop 1 url |> List.head |> Maybe.andThen String.toInt


initNote : Note
initNote =
    Note "" "" 0


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        None ->
            ( model, Cmd.none )

        SaveNote ->
            case model.page of
                EditNotePage note ->
                    saveNote model note

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
                    ( { model | message = "Couldn't load notes" }, Cmd.none )

        GotNote res ->
            case res of
                Ok n ->
                    ( { model
                        | page = ViewNotePage n.id (Just (Ok n))
                      }
                    , Cmd.none
                    )

                Err err ->
                    ( { model
                        | page = handleGotNoteError err
                      }
                    , Cmd.none
                    )

        PostedNote res ->
            case res of
                Ok _ ->
                    ( model, getNotes )

                Err _ ->
                    ( { model | message = "Failed to save note" }, Cmd.none )

        View page ->
            ( { model
                | page = page
              }
            , case page of
                ViewNotePage id _ ->
                    Cmd.batch [ getNote id, Nav.pushUrl model.key "/note/1" ]

                EditNotePage _ ->
                    Nav.pushUrl model.key "/edit"

                _ ->
                    Cmd.none
            )

        UrlChanged url ->
            init () url model.key

        LinkClicked req ->
            case req of
                Browser.Internal url ->
                    ( model, Nav.pushUrl model.key (Url.toString url) )

                Browser.External url ->
                    ( model, Nav.load url )


saveNote : Model -> Note -> ( Model, Cmd Msg )
saveNote model note =
    ( { model
        | page = HomePage
        , notes = model.notes ++ [ note ]
      }
    , Cmd.batch
        [ postNote note
        , Nav.pushUrl model.key (Url.Builder.relative [ "/" ] [])
        ]
    )


handleGotNoteError : Http.Error -> Page
handleGotNoteError err =
    ViewNotePage 0
        (Just
            (Err
                (case err of
                    Http.BadStatus 404 ->
                        "Couldn't find the note"

                    _ ->
                        "Something went wrong"
                )
            )
        )



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
        [ viewNav model |> toUnstyled
        ]
    }


viewNav : Model -> Html Msg
viewNav model =
    div
        []
        [ nav []
            [ h1 [] [ text "Notes" ]
            ]
        , main_ []
            [ text model.message
            , viewPages model
            ]
        ]


viewPages : Model -> Html Msg
viewPages model =
    case model.page of
        ViewNotePage id res ->
            viewNoteDetails id res

        EditNotePage note ->
            newNoteForm note

        HomePage ->
            div [] (viewHomePage model)


viewNoteDetails : Int -> Maybe (Result String Note) -> Html Msg
viewNoteDetails id res =
    div []
        (case res of
            Just r ->
                [ p [] [ text (String.fromInt id) ]
                , case r of
                    Ok note ->
                        showNote note

                    Err message ->
                        text message
                ]

            Nothing ->
                [ p [] [ text (String.fromInt id) ]
                , p [] [ text "Loading" ]
                ]
        )


viewHomePage : Model -> List (Html Msg)
viewHomePage model =
    [ a [ href "/edit" ] [ text "Add New note" ]
    , div [] []
    , ul [] (List.map showNote model.notes)
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
    li
        [ onClick (View (ViewNotePage 1 Nothing))
        , css
            [ border3 (px 1) solid (rgb 127 127 127)
            , cursor pointer
            , listStyleType none
            ]
        ]
        [ p
            []
            [ text note.title ]
        , p [] [ text note.description ]
        ]
