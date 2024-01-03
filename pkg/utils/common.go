package utils

import (
	"errors"
	"fmt"
	"log"
)

// Panic if error
func Pife(err error) {
	if err != nil {
		panic(err)
	}
}

// Log if the err != nil
func Loge(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ErrIfNotOk(ok bool, msg string) error {
	if ok {
		return nil
	}
	return errors.New(msg)
}

func AnyErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// Define a generic Map function type
type MapFunc[A, B any] func(A) B

// Implement the Map function
func Map[A any, B any](input []A, m MapFunc[A, B]) []B {
	output := make([]B, len(input))
	for i, element := range input {
		output[i] = m(element)
	}
	return output
}

func Filter[T any](tSlice []T, keep func(t T) bool) []T {
	out := []T{}
	for _, t := range tSlice {
		if keep(t) {
			out = append(out, t)
		}
	}
	return out
}

func ResultOkToErr[T any](v T, ok bool) func(string) (T, error) {
	return func(msg string) (T, error) {
		return v, ErrIfNotOk(ok, msg)
	}
}

func Left[L any, R any](l L, _ R) L {
	return l
}

func Right[L any, R any](_ L, r R) R {
	return r
}

func LeftLogRight[L any, R any](l L, err error) L {
	if err != nil {
		log.Println(err)
	}
	return l
}

func LeftCallbackIfErr[L any](l L, err error) func(callback func(err error)) L {
	return func(f func(err error)) L {
		if err != nil {
			f(err)
		}
		return l
	}
}

func MapGetDef[K comparable, V any](map_ map[K]V, key K, def V) V {
	v, ok := map_[key]
	if ok {
		return v
	}
	return def
}

func MapGetErr[K comparable, V any](map_ map[K]V, key K) (V, error) {
	v, ok := map_[key]
	if ok {
		return v, nil
	}
	return v, errors.New(fmt.Sprint("key not found in map:", key))
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
