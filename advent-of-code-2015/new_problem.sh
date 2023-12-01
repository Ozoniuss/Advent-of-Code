#!/bin/bash

# copy this script to the advent of code folder of the year you're playing
# make sure to have AOC_SESSION_ID set in order to also get the problems

source .env
SESSION_ID=$AOC_SESSION_ID

DAY=1
YEAR=2015

if [[ ! -z $1 ]]
then
    DAY=$1
fi

if [[ ! -z $2 ]]
then
    YEAR=$2
fi

URL="https://adventofcode.com/$YEAR/day/$DAY"
echo $URL

# works regardless the session is actually set or not
DIR=$(python 'get_problem_details.py' "$(curl -s -H "Cookie: session=$SESSION_ID" "$URL")")

if [ -d "$DIR" ]
then
    echo 'directory for that year already exists, exiting...'
    exit 0
fi

mkdir "$DIR"

touch "$DIR/main.go"
touch "$DIR/statement.txt"
touch "$DIR/input.txt"

STATUS_CODE=$(curl -s -I -H "Cookie: session=$SESSION_ID" https://adventofcode.com/2015/day/1/input 2>/dev/null | head -n 1 | cut -d$' ' -f2)

if [ $STATUS_CODE == 200 ]
then
    curl -s -H "Cookie: session=$SESSION_ID" "$URL/input" > "$DIR/input.txt"
else 
    echo 'unauthenticated, input not set'
fi

# use your own template
cp ../aocommon/template.go "$DIR/main.go"
echo $URL > "$DIR/statement.txt"
