#!/bin/bash

# copy this script to the advent of code folder of the year you're playing
# make sure to have AOC_SESSION_ID set in order to also get the problems

# works regardless the session is actually set or not
get_dir_name () { \
    curl -s -H "Cookie: session=$SESSION_ID" "https://adventofcode.com/$YEAR/day/$DAY" | \

    # find the title
    grep -o -E "\-\-\- Day [0-9][0-9]?: .+ ---" | \

    # remove --- Day ... --- border
    sed -E "s/---\sDay\s(.*)\s---$/\1/g" | \

    # remove special characters, including "," "!" and ''' because go modules
    # don't play well with them
    sed -E "s/[\:\<\>\"\?\*\,\'\!]//g" | \

    # replace whitespaces with -
    sed -E "s/\s/-/g" | \

    # lowercase
    tr '[:upper:]' '[:lower:]' | \

    # add 0 if day has 1 digit
    sed -E "s/\<([0-9])-/0\1-/g" | \

    # html unescape
    sed -E "s/\&apos\;//g" \
    ; \
}

SCRIPT_DIR=$(dirname $0)
source "$SCRIPT_DIR/.env"

SESSION_ID=$AOC_SESSION_ID

DAY=1
YEAR=2024

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

STATUS_CODE=$(curl -s -I -H "Cookie: session=$SESSION_ID" "$URL" 2>/dev/null | head -n 1 | cut -d$' ' -f2)
if [ $STATUS_CODE == 404 ]
then
    echo 'puzzle not up yet'
    exit 0
fi


# works regardless the session is actually set or not
DIR=$(get_dir_name)

if [ -d "$DIR" ]
then
    echo 'directory for that year already exists, exiting...'
    exit 0
fi

mkdir "$DIR"
touch "$DIR/main.py"
touch "$DIR/statement.txt"
touch "$DIR/input.txt"

STATUS_CODE=$(curl -s -I -H "Cookie: session=$SESSION_ID" https://adventofcode.com/$YEAR/day/$DAY/input 2>/dev/null | head -n 1 | cut -d$' ' -f2)

if [ $STATUS_CODE == 200 ]
then
    curl -s -H "Cookie: session=$SESSION_ID" "$URL/input" > "$DIR/input.txt"
else 
    echo 'unauthenticated, puzzle input not set'
fi

# use your own template
cp ../aocommon/template.py "$DIR/main.py"
echo $URL > "$DIR/statement.txt"
