package marche

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
 "strings"
)


type item struct {
	 name  string
	 price int
}

type player struct {
	money int
	inv   []item
}

func MarcheDuSoleil() {