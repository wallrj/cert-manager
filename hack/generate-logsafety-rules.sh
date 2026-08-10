#!/bin/bash
# Generates hack/logsafety.rules.go: ruleguard rules flagging pointer values
# passed to logr structured logging calls when the pointed-to type implements
# fmt.Stringer (a nil pointer panics inside the logger; wrap in klog.SafePtr).
#
# One rule function per (method, key/value-pair position) because of two
# ruleguard matcher limitations (golangci-lint v2.12.2):
#   - gogrep does not backtrack $*_ wildcards when a Where filter rejects
#     the first binding, so each value position needs its own pattern, and
#   - when several alternative patterns of a single m.Match structurally
#     match the same call and pass the filter, nothing is reported, so the
#     patterns must live in separate rule functions. Rules with identical
#     Report messages also interfere, so every message is unique.
set -e
cd "$(dirname "$0")"
out=logsafety.rules.go

underscores() { # n -> "$_, " x n
  local n=$1 s=""
  for ((j = 0; j < n; j++)); do s+='$_, '; done
  printf '%s' "$s"
}

rule() { # name method fixedargs pair
  local name=$1 method=$2 fixed=$3 pair=$4
  cat <<RULE

func ${name}(m dsl.Matcher) {
	m.Import(\`github.com/go-logr/logr\`)
	m.Match(\`\$log.${method}($(underscores "$fixed")\$v, \$*_)\`).
		Where(m[\`log\`].Type.Is(\`logr.Logger\`) &&
			m[\`v\`].Type.Is(\`*\$_\`) &&
			m[\`v\`].Type.Implements(\`fmt.Stringer\`)).
		Report(\`a nil \$v (${method} key/value pair ${pair}) panics when the logger stringifies it; wrap it in klog.SafePtr\`)
}
RULE
}

{
  cat boilerplate-go.txt
  cat <<'HEADER'

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
HEADER
  for i in 1 2 3 4 5 6; do
    # logr.Logger.Info(msg, kv...): value of pair i is at arg index 2i
    rule "nilUnsafeLogInfoValue${i}" "Info" $((2 * i)) "${i}"
  done
  for i in 1 2 3 4 5 6; do
    # logr.Logger.Error(err, msg, kv...): value of pair i is at arg index 2i+1
    rule "nilUnsafeLogErrorValue${i}" "Error" $((2 * i + 1)) "${i}"
  done
  for i in 1 2 3 4 5 6; do
    # logr.Logger.WithValues(kv...): value of pair i is at arg index 2i-1
    rule "nilUnsafeLogWithValues${i}" "WithValues" $((2 * i - 1)) "${i}"
  done
} > "$out"
echo "wrote $out"
