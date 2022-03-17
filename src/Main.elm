module Main exposing (..)

import Browser exposing (sandbox)
import Html exposing (Html, h1, text)


main : Program () {} a
main =
    sandbox
        { init = init
        , update = update
        , view = view
        }


view : Model -> Html a
view _ =
    h1 [] [ text "Notes" ]


type alias Model =
    {}


update : a -> Model -> Model
update _ x =
    x


init : {}
init =
    {}
