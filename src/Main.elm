module Main exposing (..)

import Browser
import Css exposing (..)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css)


main : Program () {} a
main =
    Browser.sandbox
        { init = init
        , update = update
        , view = view >> toUnstyled
        }


type alias Model =
    {}


init : {}
init =
    {}


update : a -> Model -> Model
update _ x =
    x



-- VIEW


view : Model -> Html a
view _ =
    div
        [ css
            [ backgroundColor (rgb 100 100 100), color (rgb 200 200 200) ]
        ]
        [ h1 [] [ text "Notes" ]
        ]
