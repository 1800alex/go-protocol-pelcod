# pelcod [![GoDoc](https://godoc.org/github.com/1800alex/go-protocol-pelcod?status.svg)](https://godoc.org/github.com/1800alex/go-protocol-pelcod) [![Build Status](https://travis-ci.com/1800alex/go-protocol-pelcod.svg?branch=master)](https://travis-ci.com/1800alex/go-protocol-pelcod)
*/}} [![Coverage Status](https://coveralls.io/repos/github/1800alex/go-protocol-pelcod/badge.svg?branch=master)](https://coveralls.io/github/1800alex/go-protocol-pelcod?branch=master)
*/}} [![Go Report Card](https://goreportcard.com/badge/github.com/1800alex/go-protocol-pelcod)](https://goreportcard.com/report/github.com/1800alex/go-protocol-pelcod)
Package go-protocol-pelcod is a package capable of building and parsing PelcoD messages.

Download:
```shell
go get github.com/1800alex/go-protocol-pelcod
```

* * *
Package go-protocol-pelcod is a package capable of building and parsing PelcoD messages.
The API can be used to communicate with Pan/Tilts or other PelcoD devices.

TODO Build NACK.





# Examples

BuildACK
Code:

```
{
	result := BuildACK([]byte{0x00, 0x01, 0x09})
	fmt.Printf("%x\n", result)
}
```


BuildSTX
Code:

```
{
	result := BuildSTX([]byte{0x00, 0x01, 0x09})
	fmt.Printf("%x\n", result)
}
```


Parse
Code:

```
{
	result, err := Parse([]byte{0x02, 0x00, 0x01, 0x09, 0x08, 0x03})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%x\n", result)
}
```



