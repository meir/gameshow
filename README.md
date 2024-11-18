# Gameshow

A small program to host a gameshow for a friend group, nothing special.

## Idea of the gameshow
The idea of the entire gameshow is that I prepare lots of questions/rounds.
The mechanics of the gameshow are as followed:
- 1 Person is the host, they control everything themselves, nothing is automated because that would take too long except for some stuff to keep track
- Every person gets assigned a team, more teams can be created or deleted by the Host and players can be dragged and dropped into their assigned teams by the Host.
- All the questions are prepared by the Host
- All teams start with a certain amount of HEALTH
- During a teams turn, they can answer a question to gain GOLD
- After they get a question correct they can choose to either USE an ITEM or BUY one
- the available items are ATTACK, DEFEND or PROTECT
- ATTACK damages a teams HEALTH
- DEFEND defends a teams HEALTH
- PROTECT destroys a teams ATTACK ITEM
- once one team is left, they WIN

## How to use

Install (golang)[https://go.dev/] and run `go run ./cmd/gs/main.go` to start the server.
If you dont want to port forward, i recommend using (ngrok)[https://ngrok.com/] to make it accessible for your players.

After that, navigate to the url and join the game and assign yourself as the host, only one person can do this.

And then just play the game with your friends!

## Why does it not do X?

It's only for keeping track of turns, teams, gold, health, etc. its not to automate the game, if youre searching for something like that, look into (jackbox)[https://jackbox.tv/]

