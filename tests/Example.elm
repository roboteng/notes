module Example exposing (..)

import Expect
import Test exposing (Test, describe, test)


suite : Test
suite =
    describe "Given the user is on the home page"
        [ test
            "2 + 2 = 4"
          <|
            \_ -> Expect.equal (2 + 2) 4
        ]
