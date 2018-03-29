# mlbcli
mlbcli is a command-line tool to display a basic MLB scoreboard in your terminal.
### Installation
On a Mac, use homebrew. If you don't have homebrew, get it [here](https://brew.sh/).
```
brew tap jimt1234/mlbcli
brew install mlbcli
```
You can also download the binary and put it in `/usr/local/bin/`:
```
wget https://github.com/jimt1234/mlbcli/releases/download/v0.1.0/mlbcli-0.1.0-darwin-amd64.zip
```
Same thing for Linux:
```
wget https://github.com/jimt1234/mlbcli/releases/download/v0.1.0/mlbcli-0.1.0-linux-386.zip
```
And [Windows](http://gnuwin32.sourceforge.net/packages/wget.htm):
```
wget https://github.com/jimt1234/mlbcli/releases/download/v0.1.0/mlbcli-0.1.0-windows-386.exe.zip
```
### Usage
Defaults to current day:
```
$ mlbcli
```
Same thing:
```
$ mlbcli today
```
Tomorrow's games and start times:
```
$ mlbcli tomorrow
```
Yesterday's scoreboard:
```
$ mlbcli yesterday
```
Schedule for the upcoming Saturday:
```
$ mlbcli Saturday
```
Scoreboard for a specific date; defaults to current year:
```
$ mlbcli 7/1
```
Scoreboard for a specific date:
```
$ mlbcli 08/27/2015
```
### Sample output:
```
$ mlbcli
+--------------+---------------------------+----+
| Dodgers      | 0 1 0 0 0 0 0 0 0         |  1 |
| Reds         | 0 0 0 0 0 0 0 0 0         |  0 |
+--------------+---------------------------+----+
| Angels       | 0 0 0 0 0 1 0 0 1         |  2 |
| Tigers       | 0 0 0 0 0 0 0 0 0         |  0 |
+--------------+---------------------------+----+
| Blue Jays    | 0 0 0 0 0 0 0 1 0         |  1 |
| Rangers      | 1 0 0 0 0 0 3 0           |  4 |
+--------------+---------------------------+----+
| Orioles      | 0 0 0 0 0 0 1 0 2         |  3 |
| Royals       | 0 0 0 2 1 1 1 0           |  5 |
+--------------+---------------------------+----+
| Cubs         | 0 1 0 0 0 0 0 0 0         |  1 |
| Giants       | 0 1 4 0 0 0 0 4           |  9 |
+--------------+---------------------------+----+
| Mets         | 0 0 0 2 3 0 0 0 0 0 0 0 4 |  9 |
| Phillies     | 0 0 5 0 0 0 0 0 0 0 0 0 0 |  5 |
+--------------+---------------------------+----+
| Padres       | 0 0 0 1 0 0 0 0 1         |  2 |
| Nationals    | 0 0 0 0 2 1 1 0           |  4 |
+--------------+---------------------------+----+
| Twins        | 0 1 1 2 0 0 0 0 0         |  4 |
| Rays         | 2 0 0 0 1 2 0 0           |  5 |
+--------------+---------------------------+----+
| Pirates      | 0 1 0 1 0 0 0 0 0         |  2 |
| Marlins      | 0 0 0 0 0 0 1 0 0         |  1 |
+--------------+---------------------------+----+
| Mariners     | 0 0 0 0 0 0 2 0 0         |  2 |
| White Sox    | 1 0 2 0 1 0 0 0           |  4 |
+--------------+---------------------------+----+
| Cardinals    | 1 2 0 0 2 0 0 0 0         |  5 |
| D-backs      | 0 3 0 0 0 0 0 0 0         |  3 |
+--------------+---------------------------+----+
```
And, of course, [the greatest game ever](https://www.youtube.com/watch?t=15623&v=K8XkzEJsKe4):
```
$ mlbcli 10/27/2011
+--------------+-----------------------+----+
| Rangers      | 1 1 0 1 1 0 3 0 0 2 0 |  9 |
| Cardinals    | 2 0 0 1 0 1 0 1 2 2 1 | 10 |
+--------------+-----------------------+----+
```
