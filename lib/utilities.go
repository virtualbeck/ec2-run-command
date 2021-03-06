package ec2

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// Hash returns a random hash n characters long
func Hash(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

// Generate a New SSH key in AWS based on instance options returns pointers to the key name and identity
func newKeyPair(logKey bool) (sshKeyName, sshKeyIdentity *string, err error) {
	name := "ec2-cli#" + Hash(10)

	input := &ec2.CreateKeyPairInput{
		KeyName: &name,
	}

	result, err := ec2Svc.CreateKeyPair(input)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to create AWS Key Pair %s: %s", name, err)
	}

	if logKey {
		fmt.Println(*result.KeyMaterial)
	}

	return &name, result.KeyMaterial, nil

}

func stringPointerValueOrNil(i *string, s string) string {
	if i == nil {
		return s
	}
	return *i
}
