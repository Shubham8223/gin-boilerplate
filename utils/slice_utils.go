package utils

func SearchQuerySlice[T comparable](queryValue T, searchList []T) bool {
	for _, value := range searchList {
		if value == queryValue {
			return true
		}
	}
	return false
}


func All[T any](slice []T, condition func(T) bool) bool {
    for _, item := range slice {
        if !condition(item) {
            return false
        }
    }
    return true
}

func Any[T any](slice []T, condition func(T) bool) bool {
    for _, item := range slice {
        if condition(item) {
            return true
        }
    }
    return false
}

func None[T any](slice []T, condition func(T) bool) bool {
    for _, item := range slice {
        if condition(item) {
            return false
        }
    }
    return true
}