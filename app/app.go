package app

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/urfave/cli/v3"
)

func Launch() *cli.Command {
	app := &cli.Command{
		Name: "pass-gen",
		Description: "Generate random and secure passwords with a specified length.",
		Version: "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "len",
				Usage: "Specifies the length of the password - Max: 255",
				Value: "10",
			},
		},
		Action: generatePass,
	}
	return app
}

func generatePass(ctx context.Context, cmd *cli.Command) error {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*_-+=?"
	passLenStr := cmd.String("len")
	passLen, err := strconv.ParseUint(passLenStr, 10, 8)

	if err != nil {
		log.Fatal(err)
	}

	password := make([]byte, passLen)

	for i := uint64(0); i < passLen; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

		if err != nil {
			log.Fatal(err)
		}

		password[i] = charset[index.Int64()]
	}

	fmt.Println(string(password))
	return nil
}
