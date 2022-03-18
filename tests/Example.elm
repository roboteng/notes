module Example exposing (..)

import Expect
import Main exposing (Msg(..), init, initNote, update)
import Test exposing (Test, describe, test)


suite : Test
suite =
    describe "Given an empty model"
        [ describe
            "WHen the user clicks on the Add Note Button"
            [ let
                action =
                    EditNote
              in
              test "Then we have state for entry"
                (\_ -> Expect.equal (update action init) { init | newNote = Just initNote })
            ]
        ]
