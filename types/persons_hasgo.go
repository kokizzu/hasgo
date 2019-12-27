// Package types contains code generated by hasgo. [DO NOT EDIT!]
package types

import (
	"fmt"
)

// =============== all.go =================

// All returns true if all elements of the slice satisfy the predicate.
// Can be generated for any type.
func (s persons) All(f func(person) bool) bool {
	if f == nil {
		return false
	}
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

// =============== any.go =================

// Any returns true if any of the elements satisfy the predicate.
// Can be generated for any type.
func (s persons) Any(f func(person) bool) bool {
	if f == nil {
		return false
	}
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// =============== break.go =================

// Break returns a tuple of any elements that do not satisfy the predicate up until the first time it passes, followed
// by the rest of the elements.
// Can be generated on any type.
func (s persons) Break(f func(person) bool) (before persons, after persons) {
	if f == nil {
		return before, s
	}

	passed := false

	for _, v := range s {
		if passed || f(v) {
			after = append(after, v)
			passed = true
		} else {
			before = append(before, v)
		}
	}

	return
}

// =============== delete.go =================

// Delete returns a slice with the first matching element
// removed from the slice.
// Can be generated for any type.
func (s persons) Delete(e person) (out persons) {
	deleted := false
	for _, v := range s {
		if deleted || v != e {
			out = append(out, v)
		} else {
			deleted = true
		}
	}
	return
}

// =============== drop.go =================

// Drop returns a new Slice with the elements after the provided key.
// Can be generated for any type.
func (s persons) Drop(i int) (out persons) {
	for i < len(s) {
		out = append(out, s[i])
		i++
	}
	return
}

// =============== dropwhile.go =================

// DropWhile returns a new slice containing all elements after the predicate fails for the first time.
// Can be generated for any type.
func (s persons) DropWhile(f func(person) bool) (out persons) {
	if f == nil {
		return s
	}
	failed := false
	for _, v := range s {
		if failed {
			out = append(out, v)
			continue
		}
		if !f(v) {
			out = append(out, v)
			failed = true
		}
	}
	return
}

// =============== elem.go =================

// Elem returns true if the slice contains the element
// Can be generated for any type.
func (s persons) Elem(el person) bool {
	for _, e := range s {
		if e == el {
			return true
		}
	}
	return false
}

// =============== filter.go =================

// Filter returns a slice containing only the elements that match the predicate.
// Can be generated for any type.
func (s persons) Filter(f func(person) bool) (out persons) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}

// =============== foldl.go =================

// Foldl reduces a list by iteratively applying f from left->right. Thus, for an empty slice, the result is the default zero-value.
func (s persons) Foldl(z person, f func(e1, e2 person) person) (out person) {
	if len(s) == 0 {
		return
	}
	out = f(z, s[0])
	for _, v := range s[1:] {
		out = f(out, v)
	}
	return
}

// =============== foldl1.go =================

// Foldl1 reduces a list by iteratively applying f from left->right. Thus, for an empty slice, the result is the default zero-value.
func (s persons) Foldl1(f func(e1, e2 person) person) (out person) {
	if len(s) == 0 {
		return
	}
	out = s[0]
	for _, v := range s[1:] {
		out = f(out, v)
	}
	return
}

// =============== foldr.go =================

// Foldr reduces a list by iteratively applying f from right -> left. Thus, for an empty slice, the result is the default zero-value.
func (s persons) Foldr(e person, f func(e1, e2 person) person) (out person) {
	if len(s) == 0 {
		return
	}

	end := len(s) - 1
	out = f(s[end], e)

	for i := end - 1; i >= 0; i-- {
		out = f(s[i], out)
	}

	return
}

// =============== foldr1.go =================

// Foldr1 reduces a list by iteratively applying f from right -> left. Thus, for an empty slice, the result is the default zero-value.
func (s persons) Foldr1(f func(e1, e2 person) person) (out person) {
	if len(s) == 0 {
		return
	}

	end := len(s) - 1
	out = s[end]

	for i := end - 1; i >= 0; i-- {
		out = f(s[i], out)
	}

	return
}

// =============== group.go =================

// Group returns a list of lists where each list contains only equal elements and the concatenation of the
// result is equal to the argument.
// Can be generated for any type.
func (s persons) Group() (out [][]person) {
	current := persons{}

	for i, v := range s {
		current = append(current, v)
		if i == len(s)-1 || v != s[i+1] {
			out = append(out, current)
			current = persons{}
		}
	}

	return
}

// =============== head.go =================

// Head returns the first element in the slice.
// If no element is found, returns the zero-value of the type.
// Can be generated for any type.
func (s persons) Head() (out person) {
	if len(s) > 0 {
		out = s[0]
	}
	return
}

// =============== init.go =================

// Init takes n-1 elements from a slice, where n = len(list).
// Can be generated for any type.
func (s persons) Init() (out persons) {
	if len(s) == 0 {
		return
	}
	slicecopy := append([]person(nil), s...)
	return slicecopy[:len(s)-1]
}

// =============== inits.go =================

// Inits returns all inits of a sequence, in order of small to large, as if it were called recursively.
// Can be generated for any type.
func (s persons) Inits() (out [][]person) {
	out = append(out, make(persons, 0))
	for i := range s {
		init := make(persons, i+1)
		for n := 0; n <= i; n++ {
			init[n] = s[n]
		}
		out = append(out, init)
	}
	return
}

// =============== intercalate.go =================

// Intercalate inserts the method receiver slice into the function slice at each step.
// Can be generated for any type.
func (s persons) Intercalate(ss [][]person) (out persons) {
	for i, slice := range ss {
		for _, el := range slice {
			out = append(out, el)
		}
		if i == len(ss)-1 {
			break
		}
		for _, el := range s {
			out = append(out, el)
		}
	}
	return out
}

// =============== intersperse.go =================

// Intersperse inserts the receiving value between each element of the method receiver.
// Can be generated for any type.
func (s persons) Intersperse(value person) (out persons) {
	for i, el := range s {
		out = append(out, el)
		if i == len(s)-1 {
			break
		}
		out = append(out, value)
	}
	return
}

// =============== last.go =================

// Last returns the last element in the slice
// If no element is found, returns the zero-value of the type
// Can be generated for any type.
func (s persons) Last() (out person) {
	if len(s) > 0 {
		out = s[len(s)-1]
	}
	return
}

// =============== length.go =================

// Length returns the length (len) of a slice.
// Can be generated for any type.
func (s persons) Length() int {
	return len(s)
}

// =============== map.go =================

// Map return a new slice with the map operation applied to each element.
// Can be generated for any type.
func (s persons) Map(f func(person) person) (out persons) {
	if f == nil {
		return s
	}
	for _, v := range s {
		out = append(out, f(v))
	}
	return
}

// =============== maximumby.go =================

// MaximumBy returns the maximum elements according to a custom comparator.
// Can be generated for any type.
func (s persons) MaximumBy(f func(e1, e2 person) person) (max person) {
	if len(s) == 0 {
		return
	}
	max = s[0]
	for _, el := range s[1:] {
		max = f(max, el)
	}
	return
}

// =============== modes.go =================

// Modes returns the elements with highest frequency in the slice.
// Can be generated for any type.
func (s persons) Modes() (out persons) {
	if len(s) == 0 {
		return
	}

	counts := make(map[person]int)
	for _, v := range s {
		counts[v]++
	}

	var max int
	for _, v := range counts {
		if v > max {
			max = v
		}
	}

	for k, v := range counts {
		if v == max {
			out = append(out, k)
		}
	}

	return
}

// =============== nub.go =================

// Nub returns a slice containing only the unique elements of the receiver.
// The order of the elements is preserved.
// Can be generated for any type.
func (s persons) Nub() (out persons) {
	if len(s) == 0 {
		return
	}

	contains := make(map[person]struct{})
	for _, v := range s {
		if _, ok := contains[v]; !ok {
			contains[v] = struct{}{}
			out = append(out, v)
		}
	}
	return
}

// =============== null.go =================

// Null returns true the slice is empty.
// Can be generated for any type.
func (s persons) Null() bool {
	return len(s) == 0
}

// =============== reverse.go =================

// Reverse returns the reversed slice.
// Can be generated for any type.
func (s persons) Reverse() (out persons) {
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return
}

// =============== scanl.go =================

// Scanl reduces a list by iteratively applying f from left->right and then returns each iteration in a slice.
func (s persons) Scanl(e person, f func(e1, e2 person) person) (out persons) {
	if len(s) == 0 {
		return
	}

	out = append(out, e)
	last := e

	for _, v := range s {
		last = f(last, v)
		out = append(out, last)
	}

	return
}

// =============== span.go =================

// Span returns a tuple of any elements that satisfy the predicate up until the first failure, followed by
// the rest of the elements.
// Can be generated for any type.
func (s persons) Span(f func(person) bool) (before persons, after persons) {
	if f == nil {
		return before, s
	}

	failed := false

	for _, v := range s {
		if failed || !f(v) {
			after = append(after, v)
			failed = true
		} else {
			before = append(before, v)
		}
	}

	return
}

// =============== splitat.go =================

// SplitAt splits the slice at the given index, returning before and after as a tuple.
// Can be generated for any type.
func (s persons) SplitAt(i int) (before, after persons) {
	for k, v := range s {
		if k < i {
			before = append(before, v)
		} else {
			after = append(after, v)
		}
	}
	return
}

// =============== tail.go =================

// Tail takes [1 -> n] elements from a slice, where n = len(list)
// Returns an empty slice if there are less than 2 elements in slice
// Can be generated for any type.
func (s persons) Tail() (out persons) {
	if len(s) <= 1 {
		return
	}
	slicecopy := append([]person(nil), s...)
	return slicecopy[1:]
}

// =============== tails.go =================

// Tails returns all tails of a sequence, in order of large to small, as if it were called recursively.
// Can be generated for any type.
func (s persons) Tails() (out [][]person) {
	slicecopy := append([]person(nil), s...)
	for range s {
		out = append(out, slicecopy)
		slicecopy = slicecopy[1:]
	}
	out = append(out, make(persons, 0))
	return
}

// =============== take.go =================

// Take takes the first n elements of the slice, or the entire slice if n > len(slice).
// Can be generated for any type.
func (s persons) Take(n uint64) (out persons) {
	if len(s) == 0 {
		return
	}
	out = make(persons, len(s))
	copy(out, s)
	if n < uint64(len(s)) {
		return out[:n]
	}
	return
}

// =============== takewhile.go =================

// TakeWhile continues appending to the output as long as the predicate is satisfied.
// Can be generated for any type.
func (s persons) TakeWhile(p func(person) bool) (out persons) {
	if len(s) == 0 {
		return
	}
	for _, e := range s {
		if !p(e) {
			return
		}
		out = append(out, e)
	}
	return
}

// =============== uncons.go =================

// Uncons decomposes a slice into the head and tail component.
// Can be generated for any type.
func (s persons) Uncons() (head person, tail persons) {
	return s.Head(), s.Tail()
}

// =============== unlines.go =================

// Unlines joins together the string representation of the slice with newlines after each element.
// Can be generated for any type.
func (s persons) Unlines() (out string) {
	for i, v := range s {
		out += fmt.Sprintf("%v", v)
		if i != len(s)-1 {
			out += "\n"
		}
	}
	return
}

// =============== unwords.go =================

// Unwords joins together the string representation of the slice with newlines after each element.
// Can be generated for any type.
func (s persons) Unwords() (out string) {
	for i, v := range s {
		out += fmt.Sprintf("%v", v)
		if i != len(s)-1 {
			out += " "
		}
	}
	return
}
