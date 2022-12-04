package logger

import "github.com/sirupsen/logrus"

func New(lvl string) (*logrus.Logger, error) {
	logg := logrus.New()
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return nil, err
	}

	logg.SetLevel(level)

	return logg, err
}
