module Example exposing (..)

import Expect
import Main exposing (Msg(..), init, initNote, update)
import Test exposing (Test, describe, test)


suite : Test
suite =
    describe "Given an empty model"
        (let
            model =
                init
         in
         [ describe
            "WHen the user clicks on the Add Note Button"
            (let
                action =
                    EditNote
             in
             [ test "Then we have state for entry"
                (\_ -> Expect.equal (update action model) { model | newNote = Just initNote })
             ]
            )
         ]
        )


other : Test
other =
    describe "Given the Edit Note page is open"
        (let
            model =
                { init | newNote = Just { title = "Title", description = "Description" } }
         in
         [ describe "When the user save a note with data"
            (let
                action =
                    SaveNote
             in
             [ test "Then the newNote is Nothing"
                (\_ ->
                    Expect.equal
                        { notes = [ { title = "Title", description = "Description" } ], newNote = Nothing }
                        (update action model)
                )
             ]
            )
         ]
        )
