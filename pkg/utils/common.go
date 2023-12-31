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

func Left[T any](v T, err error) T {
	return v
}

func Right[T any](v T, err error) error {
	return err
}

func LeftLogRight[T any](v T, err error) T {
	if err != nil {
		log.Println(err)
	}
	return v
}

func LeftCallbackIfErr[T any](v T, err error) func(callback func(err error)) T {
	return func(f func(err error)) T {
		if err != nil {
			f(err)
		}
		return v
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
