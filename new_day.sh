#!/bin/bash

# check if args length is 2
if [ $# -ne 2 ]; then
    echo "Usage: ./new_day.sh <directory> <day>"
    exit 1
fi

# directory where the new day will be created
DIR=$1
DAY=$2

if [ ! -d "$DIR" ]; then
    echo "Directory $DIR does not exist"
    exit 1
fi

if [ ! -d "$DIR/inputs" ]; then
    mkdir $DIR/inputs
fi

cd $DIR/inputs

EXAMPLE="day_${DAY}_example.txt"
INPUT="day_${DAY}_input.txt"

if [ ! -f "$EXAMPLE" ]; then
    touch $EXAMPLE
else
    echo "File $EXAMPLE already exists"
fi

if [ ! -f "$INPUT" ]; then
    touch $INPUT
else
    echo "File $INPUT already exists"
fi

# back in $DIR
cd ..

# create programming files

if [ $DIR = "python" ]; then
    if [ ! -f "day${DAY}.py" ]; then
        touch day${DAY}.py
        echo "def part_one(filename):" >> day${DAY}.py
        echo "    pass" >> day${DAY}.py
        echo "" >> day${DAY}.py
        echo "def part_two(filename):" >> day${DAY}.py
        echo "    pass" >> day${DAY}.py
    else
        echo "File day${DAY}.py already exists"
    fi
fi

if [ $DIR = "go" ]; then
    if [ ! -d "pkg" ]; then
        mkdir pkg
    fi 
    cd pkg

    if [ ! -d "day${DAY}" ]; then
        mkdir day${DAY}
    fi

    cd day${DAY}

    if [ ! -f "day${DAY}.go" ]; then
        touch day${DAY}.go
        echo "package day${DAY}" >> day${DAY}.go
        echo "" >> day${DAY}.go
        echo "import (" >> day${DAY}.go
        echo "    \"fmt\"" >> day${DAY}.go
        echo ")" >> day${DAY}.go
        echo "" >> day${DAY}.go
        echo "func Day${DAY}Part1(filename string) (int, error) {" >> day${DAY}.go
        echo "}" >> day${DAY}.go
        echo "" >> day${DAY}.go
        echo "func Day${DAY}Part2(filename string) (int, error) {" >> day${DAY}.go
        echo "}" >> day${DAY}.go
    else
        echo "File day${DAY}.go already exists"
    fi

    if [ ! -f "day${DAY}_test.go" ]; then
        touch day${DAY}_test.go
        echo "package day${DAY}" >> day${DAY}_test.go
        echo "" >> day${DAY}_test.go
        echo "import (" >> day${DAY}_test.go
        echo "    \"testing\"" >> day${DAY}_test.go
        echo ")" >> day${DAY}_test.go
        echo "" >> day${DAY}_test.go
        echo "func TestDay${DAY}Part1Example(t *testing.T) {" >> day${DAY}_test.go
        echo "    num, err := Day${DAY}Part1(\"inputs/day_${DAY}_example.txt\")" >> day${DAY}_test.go
        echo "    if err != nil {" >> day${DAY}_test.go
        echo "        t.Error(err)" >> day${DAY}_test.go
        echo "    }" >> day${DAY}_test.go
        echo "    expected := 0" >> day${DAY}_test.go
        echo "    if num != expected {" >> day${DAY}_test.go
        echo "        t.Errorf(\"Expected 0, got %d\", num)" >> day${DAY}_test.go
        echo "    }" >> day${DAY}_test.go
        echo "}" >> day${DAY}_test.go
        echo "" >> day${DAY}_test.go
        echo "func TestDay${DAY}Part1Input(t *testing.T) {" >> day${DAY}_test.go
        echo "    num, err := Day${DAY}Part1(\"inputs/day_${DAY}_example.txt\")" >> day${DAY}_test.go
        echo "    if err != nil {" >> day${DAY}_test.go
        echo "        t.Error(err)" >> day${DAY}_test.go
        echo "    }" >> day${DAY}_test.go
        echo "    expected := 0" >> day${DAY}_test.go
        echo "    if num != expected {" >> day${DAY}_test.go
        echo "        t.Errorf(\"Expected 0, got %d\", num)" >> day${DAY}_test.go
        echo "    }" >> day${DAY}_test.go
        echo "}" >> day${DAY}_test.go
    else
        echo "File day${DAY}_test.go already exists"
    fi
fi
