package myerr

import "errors"

func InitError()(err error)  {
	return  errors.New("init fail")
}