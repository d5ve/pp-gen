package main

/*
pp-gen

Use EFF's general short wordlist (with unique prefixes) to generate random
passphrases.

https://www.eff.org/files/2016/09/08/eff_short_wordlist_2_0.txt contains
1269 memorable and distinct words for passphrases, all with unique 3-char
prefixes and edit distances of at least 3.

See https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases for
the full rationale behind the wordlist.

eff_short_wordlist_2_0.txt is released by the EFF under the Creative Commons
Attribution License - https://creativecommons.org/licenses/by/3.0/us/

Each generated passphrase is printed in full, as well as with just the unique
3-char prefixes for a shorter string with the same entropy.
*/

// TODO: Embed the short wordlist directly into the binary.

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strings"
)

func usage() {
	fmt.Fprintln(os.Stderr, "pp-gen - generate random passphrases from EFF's general short wordlist (with unique prefixes).")
	flag.PrintDefaults()
}

func generatePassphrase(length int, wordlist []string) {
	var pp []string
	numWords := len(wordlist)
	for i := 0; i < length; i++ {
		rndb, err := rand.Int(rand.Reader, big.NewInt(int64(numWords)))
		if err != nil {
			panic(err)
		}
		rndi := rndb.Int64()
		pp = append(pp, wordlist[rndi])
	}
	for _, word := range pp {
		fmt.Printf("%v ", word)
	}
	fmt.Print(":: ")
	for _, word := range pp {
		fmt.Printf("%v ", word[0:3])
	}
	fmt.Print("\n")

}

func loadWordlist() (wordlist []string) {
	data, err := ioutil.ReadFile("eff_short_wordlist_2_0.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if len(line) > 0 {
			f := strings.Fields(line)
			wordlist = append(wordlist, f[1])
		}
	}
	return wordlist
}

func main() {
	flag.Usage = usage
	length := flag.Int("len", 0, "Length of passphrase in words")
	number := flag.Int("num", 0, "Number of passphrases to generate at a time")
	flag.Parse()

	if *length > 0 && *number > 0 {
		wordlist := loadWordlist()
		fmt.Printf("# Each %d-word passphrase has about %0.2f bits of entropy\n", *length, math.Log2(float64(len(wordlist)))*float64(*length))
		for i := 0; i < *length; i++ {
			generatePassphrase(*length, wordlist)
		}
	} else {
		flag.Usage()
	}
}

/*
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org/>
*/
