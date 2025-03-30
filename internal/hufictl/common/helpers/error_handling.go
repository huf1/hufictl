package helpers

import "github.com/sirupsen/logrus"

func FatalOnError(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
