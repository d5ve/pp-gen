# pp-gen

Use EFF's general short wordlist (with unique prefixes) to generate random
passphrases.

https://www.eff.org/files/2016/09/08/eff_short_wordlist_2_0.txt contains
1269 memorable and distinct words for passphrases, all with unique 3-char
prefixes and edit distances of at least 3.

See https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases for
the full rationale behind the wordlist.

Each generated passphrase is printed in full, as well as with just the unique
3-char prefixes for a shorter string with the same entropy.

## Run: `go run pp-gen.go -len= -num=`

```bash
$ go run pp-gen.go -len 6 -num 2
# Each 6-word passphrase has about 62.04 bits of entropy
krypton gothic bouquet vivacious fountain scythe :: kry got bou viv fou scy
already jiffy clubhouse doll garlic podiatrist :: alr jif clu dol gar pod
```

## Build: `go build pp-gen.go`

# LICENCE

https://www.eff.org/files/2016/09/08/eff_short_wordlist_2_0.txt is released by
the EFF under the Creative Commons Attribution License -
https://creativecommons.org/licenses/by/3.0/us/

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
