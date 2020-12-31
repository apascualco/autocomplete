package main

type StringSorter []string

func (a StringSorter) Len() int {
	return len(a)
}

func (a StringSorter) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a StringSorter) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
