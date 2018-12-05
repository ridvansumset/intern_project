# Restaurant Menu App

is a project that is developed by my first intern and me in his internship period.

## General Menu Schema

                Category
               /        \
            Product1     Product2
            |            /       \
        Option1      Option1    Option2
         /     \                   |
     Choice1   Choice2          Choice1

## How to Get & Run

`go get -u github.com/ridvansumset/restaurant-menu-app/`

or, simply, use `git clone`. Then,

`cd ~/YOUR_GO_FOLDER/src/github.com/ridvansumset/restaurant-menu-app`

and then, open 2 terminal screens in the current path. In the first screen,

`cd server && go run main.go`

In the second screen,

`cd web/src/app`

then, run `python -m SimpleHTTPServer` and go to http://localhost:8000/#/ in your browser.
