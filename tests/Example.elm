module Example exposing (..)

import Expect
import Main exposing (Msg(..), init, initNote, update)
import Test exposing (Test, describe, test)


emptyModel : Test
emptyModel =
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
                (\_ -> Expect.equal { model | newNote = Just initNote } (update action model))
             ]
            )
         ]
        )


editNotePage : Test
editNotePage =
    describe "Given the Edit Note page is open"
        (let
            model =
                { init | newNote = Just { title = "Title", description = "Description" } }
         in
         [ describe "When the user save a note"
            (let
                action =
                    SaveNote
             in
             [ test "Then the newNote is Nothing"
                (\_ ->
                    Expect.equal
                        Nothing
                        (update action model).newNote
                )
             , test "Then the notes list is populated"
                (\_ ->
                    Expect.equal
                        [ { title = "Title", description = "Description" } ]
                        (update action model).notes
                )
             ]
            )
         , describe "When the title is updated"
            (let
                action =
                    UpdateTitle "Title2"
             in
             [ test "Then the title should be updated"
                (\_ ->
                    Expect.equal "Title2"
                        (case (update action model).newNote of
                            Just note ->
                                note.title

                            Nothing ->
                                "Wrong"
                        )
                )
             ]
            )
         ]
        )
