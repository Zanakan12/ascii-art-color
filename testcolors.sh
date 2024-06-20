#!/bin/bash

echo "Try passing as arguments --color red banana."
go run . --color red "banana"
echo ""

echo "Try passing as arguments --color=green 1 + 1 = 2."
go run . --color=green "1 + 1 = 2"
echo ""

echo "Try passing as arguments --color=yellow (%&) ??"
go run . --color=yellow "(%&) ??"
echo ""

echo "Try specifying a substring to be colored (the second until the last letter)."
go run . --color=blue "xample text" "example text"
echo ""

echo "Try specifying letter to be colored (the second letter)."
go run . --color=purple "r" "tribute"
echo ""

echo "Try specifying a substring to be colored (just two letters)."
go run . --color=cyan "FT" "RAFTA LE BIG NARCISSIQUE"
echo ""

echo "Try passing as arguments --color=orange GuYs \"HeY GuYs\", in order to color GuYs."
go run . --color=orange "GuYs" "HeY GuYs"
echo ""

echo "Try passing as arguments --color=blue B 'RGB()', in order to color just the B."
go run . --color=blue "B" "RGB()"
echo ""

echo "Try passing as arguments a random string with lower and upper case letters, and a random color in the color flag (\"--color=\")."
go run . --color=random "RandomString"
echo ""

echo "Try passing as arguments a random string with lower case letters, numbers and spaces, and a random color in the color flag (\"--color=\")."
go run . --color=random "random 1234 string"
echo ""

echo "Try passing as arguments a random string with special characters, and a random color in the color flag (\"--color=\")."
go run . --color=random "#" "@!#&"
echo ""

echo "Try passing as arguments a random string with lower case letters, upper case letters, spaces and numbers with a random color in the color flag (\"--color=\")."
go run . --color=random "String" "Complex String 123"
echo ""

echo "Try passing as arguments --color=#ff0000 \"there\" \"Hello there\""
go run . --color=#ff0000 "there" "Hello there"
echo ""

echo "Try passing as arguments --color='rgb(255,0,0)' \"test\" \"color test\""
go run . --color='rgb(255,0,0)' "test" "color test"
echo ""

echo "Try passing as arguments --color='hsl(0,100%,50%)' \"THANKS\" \"THANKS FOR YOUR TIME\""
go run . --color='hsl(0,100%,50%)' "THANKS" "THANKS FOR YOUR TIME"
echo ""

echo "//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////"
