package kzutils

func ConvertValToPtr[T any](val T) *T {
	return &val
}

func ConvertPtrToVal[T any](ptr *T) T {
	if ptr == nil {
		var z T
		return z
	}
	return *ptr
}

func ConvertPtrsToVals[A any](items []*A) (ret []A) {
	ret = make([]A, 0, len(items))
	for _, item := range items {
		ret = append(ret, ConvertPtrToVal(item))
	}
	return ret
}

func ConvertValsToPtrs[A any](items []A) (ret []*A) {
	ret = make([]*A, 0, len(items))
	for _, item := range items {
		ret = append(ret, ConvertValToPtr(item))
	}
	return ret
}
