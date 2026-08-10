/*
Copyright The cert-manager Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/


//go:build ruleguard

// Package gorules contains ruleguard lint rules, run by the "ruleguard"
// checker of gocritic via golangci-lint (see .golangci.yaml).
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// The rules below flag pointer values passed to structured logging calls
// when the pointed-to type implements fmt.Stringer. The logger formats
// such a value by calling String() through the pointer; String is
// promoted from the embedded type with a value receiver, so a nil pointer
// panics inside the logger, which recovers and renders the field as
// "<panic: runtime error: invalid memory address or nil pointer
// dereference>". See https://github.com/cert-manager/cert-manager/issues/6799.
// The fix at the call site is to wrap the value in klog.SafePtr.
//
// One rule function per (method, key/value-pair position) because:
//   - gogrep does not backtrack $*_ wildcards when a Where filter rejects
//     the first binding, so each value position needs its own pattern, and
//   - ruleguard reports nothing when several alternative patterns of a
//     single m.Match structurally match the same call and pass the filter,
//     so the patterns must live in separate rule functions.
//
// Only the first 6 key/value pairs of each call are checked.

func nilUnsafeLogInfoValue1(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Info($_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Info key/value pair 1) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogInfoValue2(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Info($_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Info key/value pair 2) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogInfoValue3(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Info($_, $_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Info key/value pair 3) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogInfoValue4(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Info($_, $_, $_, $_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Info key/value pair 4) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogInfoValue5(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Info($_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Info key/value pair 5) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogInfoValue6(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Info($_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Info key/value pair 6) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogErrorValue1(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Error($_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Error key/value pair 1) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogErrorValue2(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Error($_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Error key/value pair 2) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogErrorValue3(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Error($_, $_, $_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Error key/value pair 3) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogErrorValue4(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Error($_, $_, $_, $_, $_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Error key/value pair 4) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogErrorValue5(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Error($_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Error key/value pair 5) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogErrorValue6(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.Error($_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (Error key/value pair 6) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogWithValues1(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.WithValues($_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (WithValues key/value pair 1) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogWithValues2(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.WithValues($_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (WithValues key/value pair 2) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogWithValues3(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.WithValues($_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (WithValues key/value pair 3) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogWithValues4(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.WithValues($_, $_, $_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (WithValues key/value pair 4) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogWithValues5(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.WithValues($_, $_, $_, $_, $_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (WithValues key/value pair 5) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}

func nilUnsafeLogWithValues6(m dsl.Matcher) {
	m.Import(`github.com/go-logr/logr`)
	m.Match(`$log.WithValues($_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $_, $v, $*_)`).
		Where(m[`log`].Type.Is(`logr.Logger`) &&
			m[`v`].Type.Is(`*$_`) &&
			m[`v`].Type.Implements(`fmt.Stringer`)).
		Report(`a nil $v (WithValues key/value pair 6) panics when the logger stringifies it; wrap it in klog.SafePtr`)
}
