module Spa exposing (..)

import Browser
import Browser.Navigation as Nav
import Html exposing (Html, a, br, button, div, text)
import Html.Attributes exposing (href)
import Html.Events exposing (onClick)
import Url


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
    { key : Nav.Key
    , url : Url.Url
    }


init : () -> Url.Url -> Nav.Key -> ( Model, Cmd Msg )
init _ url key =
    ( Model key url, Cmd.none )


type Msg
    = None
    | LinkClicked Browser.UrlRequest
    | UrlChanged Url.Url


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        None ->
            ( model, Cmd.none )

        LinkClicked link ->
            case link of
                Browser.Internal url ->
                    ( model, Nav.pushUrl model.key (Url.toString url) )

                Browser.External href ->
                    ( model, Nav.load href )

        UrlChanged url ->
            ( { model | url = url }, Cmd.none )


view : Model -> Browser.Document Msg
view model =
    Browser.Document
        "Sample"
        [ div []
            [ text "Hello"
            , br [] []
            , text model.url.path
            , br [] []
            , a [ href "/edit" ] [ text "Go" ]
            ]
        ]
