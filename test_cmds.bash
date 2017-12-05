#!/bin/bash


function assert_exists() {
    if [ "$#" != "2" ]; then
        echo "wrong number of parameters for $1, $@"
        exit 1
    fi
    if [[ ! -f "$2" && ! -d "$2" ]]; then
        echo "$1: $2 does not exists"
        exit 1
    fi

}

function assert_equal() {
    if [ "$#" != "3" ]; then
        echo "wrong number of parameters for $1, $@"
        exit 1
    fi
    if [ "$2" != "$3" ]; then
        echo "$1: expected |$2| got |$3|"
        exit 1
    fi
}



#
# Tests
#

function test_byline() {
    EXPECTED="By J. Q. Public 2017-12-04"
    # Test reading from file
    RESULT=$(bin/byline -i demo/byline/index.md)
    assert_equal "test_byline" "$EXPECTED" "$RESULT"
    # Test with standard input
    RESULT=$(cat demo/byline/index.md | bin/byline -i - )
    assert_equal "test_byline" "$EXPECTED" "$RESULT"
    echo "test_byline OK"
}

function test_mkpage() {
    bin/mkpage content=demo/mkpage/helloworld.md page.tmpl > temp.html
    EXPECTED=""
    RESULT=$(cmp temp.html demo/mkpage/helloworld.html)
    assert_exists "test_mkpage (simple)" "temp.html"
    assert_equal "test_mkpage (simple)" "$EXPECTED" "$RESULT"
    rm temp.html
    echo "test_mkpage() OK"
}

function test_mkrss() {
    echo "test_mkrss() not implemented."
}

function test_mkslides() {
    echo "test_mkslides() not implemented."
}

function test_reldocpath() {
    echo "test_reldocpath() not implemented."
}

function test_sitemapper() {
    echo "test_sitemapper() not implemented."
}

function test_titleline() {
    echo "test_titleline() not implemented."
}

function test_urldecode() {
    echo "test_urldecode() not implemented."
}

function test_urlencode() {
    echo "test_urlencode() not implemented."
}

function test_ws() {
    echo "test_ws() not implemented."
}

echo "Testing command line tools"
test_byline
test_mkpage
test_mkrss
test_mkslides
test_reldocpath
test_sitemapper
test_titleline
test_urldecode
test_urlencode
test_ws
echo 'Success!'
