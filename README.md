<p align="center">
<img height="150" width="150" src="https://cdn.simpleicons.org/premierleague/gray"/>
</p>

<div align="center">
    <h1>Premier League CLI</h1>
    <img src="https://img.shields.io/github/v/release/digitalghost-dev/premier-league-cli?style=flat-square&logo=git&logoColor=38003C&label=Release%20Version&labelColor=EEE&color=38003C">
    <img src="https://img.shields.io/github/actions/workflow/status/digitalghost-dev/premier-league-cli/go_tests.yml?style=flat-square&logo=go&logoColor=00ADD8&label=Tests&labelColor=EEE&color=00ADD8">
    <img src="https://img.shields.io/github/go-mod/go-version/digitalghost-dev/premier-league-cli?style=flat-square&logo=Go&labelColor=EEE&color=00ADD8">
</div>

## Overview
A tool for viewing data relating to the Premier League but through a command line!

> [!IMPORTANT]
> This project is an addition to a data engineering project that I've been working on which includes full end-to-end data pipelines with visualiztions in Streamlit.

## Links
* [Data Engineering Project](https://github.com/digitalghost-dev/premier-league)

## Infrastructure

## Usage
Currently, the CLI support one command and three flags:
```
Root Command:
    pl-cli
Options:
    -d, --delete             Deletes standings.csv file
    -s, --standings          Prints current standings
    -u, --update             Updates standings.csv file
```

Example output:
```
+-------+-------+-------------------+--------------+-------+-------+-------+-------------+--------+-----------+---------------+-----------------+
|       | RANK  |       TEAM        | GAMES PLAYED | WINS  | DRAWS | LOSES | RECENT FORM | POINTS | GOALS FOR | GOALS AGAINST | GOAL DIFFERENCE |
+-------+-------+-------------------+--------------+-------+-------+-------+-------------+--------+-----------+---------------+-----------------+
|  0:   |   1   |     Liverpool     |      27      |  19   |   6   |   2   |    WWWWL    |   63   |    64     |      25       |       39        |
|  1:   |   2   |  Manchester City  |      27      |  19   |   5   |   3   |    WWWDW    |   62   |    62     |      27       |       35        |
|  2:   |   3   |      Arsenal      |      27      |  19   |   4   |   4   |    WWWWW    |   61   |    68     |      23       |       45        |
|  3:   |   4   |    Aston Villa    |      27      |  17   |   4   |   6   |    WWWLW    |   55   |    59     |      37       |       22        |
|  4:   |   5   |     Tottenham     |      26      |  15   |   5   |   6   |    WLWDW    |   50   |    55     |      39       |       16        |
|  5:   |   6   | Manchester United |      27      |  14   |   2   |  11   |    LLWWW    |   44   |    37     |      39       |       -2        |
|  6:   |   7   |     West Ham      |      27      |  12   |   6   |   9   |    WWLLL    |   42   |    43     |      47       |       -4        |
|  7:   |   8   |     Newcastle     |      27      |  12   |   4   |  11   |    WLDWD    |   40   |    57     |      45       |       12        |
|  8:   |   9   |     Brighton      |      27      |  10   |   9   |   8   |    LDWLW    |   39   |    49     |      44       |        5        |
|  9:   |  10   |      Wolves       |      27      |  11   |   5   |  11   |    LWWLW    |   38   |    40     |      43       |       -3        |
|  10:  |  11   |      Chelsea      |      26      |  10   |   6   |  10   |    DDWLL    |   36   |    44     |      43       |        1        |
|  11:  |  12   |      Fulham       |      27      |  10   |   5   |  12   |    WWLWD    |   35   |    39     |      42       |       -3        |
|  12:  |  13   |    Bournemouth    |      26      |   8   |   7   |  11   |    WLDLD    |   31   |    35     |      47       |       -12       |
|  13:  |  14   |  Crystal Palace   |      27      |   7   |   7   |  13   |    LWDLL    |   28   |    32     |      47       |       -15       |
|  14:  |  15   |     Brentford     |      27      |   7   |   5   |  15   |    DLLLW    |   26   |    39     |      50       |       -11       |
|  15:  |  16   |      Everton      |      27      |   8   |   7   |  12   |    LDDLD    |   25   |    29     |      37       |       -8        |
|  16:  |  17   | Nottingham Forest |      27      |   6   |   6   |  15   |    LLWLD    |   24   |    34     |      49       |       -15       |
|  17:  |  18   |       Luton       |      26      |   5   |   5   |  16   |    LLLLD    |   20   |    37     |      54       |       -17       |
|  18:  |  19   |      Burnley      |      27      |   3   |   4   |  20   |    LLLLD    |   13   |    25     |      60       |       -35       |
|  19:  |  20   |   Sheffield Utd   |      27      |   3   |   4   |  20   |    LLLWL    |   13   |    22     |      72       |       -50       |
+-------+-------+-------------------+--------------+-------+-------+-------+-------------+--------+-----------+---------------+-----------------+
| 20X11 | INT64 |      STRING       |    INT64     | INT64 | INT64 | INT64 |   STRING    | INT64  |   INT64   |     INT64     |      INT64      |
+-------+-------+-------------------+--------------+-------+-------+-------+-------------+--------+-----------+---------------+-----------------+
```

## Security