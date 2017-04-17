## partyfy-server

Project to learn and improve my skills in:
- Go lang
- MongoDB
- Spotify API

## Install

- Clone `git clone https://github.com/felipeparpinelli/partyfy-server.git`
- Install dependencies
  - `go get gopkg.in/mgo.v2` MongoDB driver for Go
  - `go get github.com/gorilla/mux` A powerful URL router and dispatcher for golang
  - `go get github.com/zmb3/spotify` wrapper for Spotify Web API

- Install and run mongoDB
- Set Environment variables:
`export SPOTIFY_ID="client_ID"`
`export SPOTIFY_SECRET="secret_ID"`
  - https://developer.spotify.com/my-applications

## Tests
- `cd controllers`
- `go test *.go`

## Running

- Start the server `go run *.go`

- Follow the instructions to run the frontend service: `https://github.com/felipeparpinelli/partyfy-client`

* You can view the project up and running on [AWS](http://ec2-34-205-81-9.compute-1.amazonaws.com:8080)

* You can create the database with a collection(party) automatically, just uncomment the line 19 in db.go and run the application.

## Future features (TODO List)

- Create websocket between client and server - collaborativePlaylist
- Create a music search
- User be able to create parties
