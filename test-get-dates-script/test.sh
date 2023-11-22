#!/bin/bash

# curl -s https://adventofcode.com/2022/day/3 | grep -o -E "\-\-\- Day [0-9][0-9]?: .+ ---" |  sed -E "s/---\sDay\s(.*)\s---$/\1/g" | sed -E "s/[\:\<\>\"\?\*]//g" | sed -E "s/\s/-/g" | tr '[:upper:]' '[:lower:]' | sed -E "s/\<([0-9])-/0\1-/g"

SESSION_ID=53616c7465645f5f173a84ef1191ca3fdf822237221e5ea4aa9e76b78888e30735a0f264d4c10822c33f1e289ce57714e9564733ece6669364e894bf3f2a2897

get_linux_style () { \
    curl -s -H "Cookie: session=$SESSION_ID" "https://adventofcode.com/$YEAR/day/$DAY" | \
    grep -o -E "\-\-\- Day [0-9][0-9]?: .+ ---" | \
    sed -E "s/---\sDay\s(.*)\s---$/\1/g" | \
    sed -E "s/[\:\<\>\"\?\*]//g" | \
    sed -E "s/\s/-/g" | \
    tr '[:upper:]' '[:lower:]' | \
    sed -E "s/\<([0-9])-/0\1-/g" \
    ; \
}

get_python_style () {
    python 'get_problem_details_legacy.py' "$(curl -s -H "Cookie: session=$SESSION_ID" https://adventofcode.com/$YEAR/day/$DAY)"
}

get_dir_names () {
    get_linux_style
    get_python_style
}

for YEAR in {2015..2022}
do 
    for DAY in {1..25} 
    do
        get_dir_names >> out.txt
        echo "" >> out.txt
    done 
done 
