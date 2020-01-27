package main

import (
	"fmt"

	"github.com/gcash/bchutil/hdkeychain"
)

func main() {
	serialized := "xprv9s21ZrQH143K25nXfd4XuZ6d955Y7XQ7GzSrew7bEYCx1eC6JixjXoaULDjgpB4BJaLL4BmCTka4L9Jq5rxstLcXLnVkGK4tkMk8D9iv1y8"

	hdkey, err := hdkeychain.NewKeyFromString(serialized)
	if err != nil {
		panic(err)
	}

	pubkey, err := hdkey.Neuter()
	if err != nil {
		panic(err)
	}

	fmt.Println("serialized xprv:", serialized)
	fmt.Println("xpub key:", pubkey.String())
}
