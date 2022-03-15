module Example exposing (..)

import Expect exposing (Expectation)
import Html exposing (h1, text)
import Main exposing (view)
import Test exposing (Test, describe, test)


suite : Test
suite =
    describe "Given the user is on the home page"
        [ test
            "Then \"Ntes\" is visible"
          <|
            \_ -> Expect.equal (h1 [] [ text "Notes" ]) (view {})
        ]
