module Example exposing (..)

import Expect
import Json.Decode exposing (decodeString)
import Main exposing (Msg(..), Note, Page(..), noteDecoder)
import Test exposing (Test, test)
import Url exposing (Protocol(..))


parseJsonNote : Test
parseJsonNote =
    test "json should be parsed" (\_ -> Expect.equal (Ok (Note "My Title" "My Desc" (Just 1))) (decodeString noteDecoder """{"title":"My Title","desc":"My Desc","id":1}"""))
