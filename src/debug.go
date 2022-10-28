package main

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*
 * FILE: static_checking.go
 * PURPOSE: 
 */

func create_data_name(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	final := ""
	for i := 0; i < length; i += 1 {

		final += string(chars[int(math.Abs(float64((int64(rand.Intn((len(chars))))*time.Now().UnixNano()/int64(1000000))%int64(len(chars)))))])
	}
	return final
}
