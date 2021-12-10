// Code generated by go generate. DO NOT EDIT.

//go:generate rm pkg.go
//go:generate go run ../gen/gen.go

package strconv

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"
)

func init() {
	internal.Register("strconv", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{{
		Name: "Unquote",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret, c.Err = Unquote(s)
			}
		},
	}, {
		Name: "ParseBool",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			str := c.String(0)
			if c.Do() {
				c.Ret, c.Err = ParseBool(str)
			}
		},
	}, {
		Name: "FormatBool",
		Params: []internal.Param{
			{Kind: adt.BoolKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			b := c.Bool(0)
			if c.Do() {
				c.Ret = FormatBool(b)
			}
		},
	}, {
		Name: "ParseFloat",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.IntKind},
		},
		Result: adt.NumKind,
		Func: func(c *internal.CallCtxt) {
			s, bitSize := c.String(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = ParseFloat(s, bitSize)
			}
		},
	}, {
		Name:  "IntSize",
		Const: "64",
	}, {
		Name: "ParseUint",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.IntKind},
			{Kind: adt.IntKind},
		},
		Result: adt.IntKind,
		Func: func(c *internal.CallCtxt) {
			s, base, bitSize := c.String(0), c.Int(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = ParseUint(s, base, bitSize)
			}
		},
	}, {
		Name: "ParseInt",
		Params: []internal.Param{
			{Kind: adt.StringKind},
			{Kind: adt.IntKind},
			{Kind: adt.IntKind},
		},
		Result: adt.IntKind,
		Func: func(c *internal.CallCtxt) {
			s, base, bitSize := c.String(0), c.Int(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = ParseInt(s, base, bitSize)
			}
		},
	}, {
		Name: "Atoi",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.IntKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret, c.Err = Atoi(s)
			}
		},
	}, {
		Name: "FormatFloat",
		Params: []internal.Param{
			{Kind: adt.NumKind},
			{Kind: adt.IntKind},
			{Kind: adt.IntKind},
			{Kind: adt.IntKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			f, fmt, prec, bitSize := c.Float64(0), c.Byte(1), c.Int(2), c.Int(3)
			if c.Do() {
				c.Ret = FormatFloat(f, fmt, prec, bitSize)
			}
		},
	}, {
		Name: "FormatUint",
		Params: []internal.Param{
			{Kind: adt.IntKind},
			{Kind: adt.IntKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			i, base := c.Uint64(0), c.Int(1)
			if c.Do() {
				c.Ret = FormatUint(i, base)
			}
		},
	}, {
		Name: "FormatInt",
		Params: []internal.Param{
			{Kind: adt.IntKind},
			{Kind: adt.IntKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			i, base := c.Int64(0), c.Int(1)
			if c.Do() {
				c.Ret = FormatInt(i, base)
			}
		},
	}, {
		Name: "Quote",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret = Quote(s)
			}
		},
	}, {
		Name: "QuoteToASCII",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret = QuoteToASCII(s)
			}
		},
	}, {
		Name: "QuoteToGraphic",
		Params: []internal.Param{
			{Kind: adt.StringKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			s := c.String(0)
			if c.Do() {
				c.Ret = QuoteToGraphic(s)
			}
		},
	}, {
		Name: "QuoteRune",
		Params: []internal.Param{
			{Kind: adt.IntKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			r := c.Rune(0)
			if c.Do() {
				c.Ret = QuoteRune(r)
			}
		},
	}, {
		Name: "QuoteRuneToASCII",
		Params: []internal.Param{
			{Kind: adt.IntKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			r := c.Rune(0)
			if c.Do() {
				c.Ret = QuoteRuneToASCII(r)
			}
		},
	}, {
		Name: "QuoteRuneToGraphic",
		Params: []internal.Param{
			{Kind: adt.IntKind},
		},
		Result: adt.StringKind,
		Func: func(c *internal.CallCtxt) {
			r := c.Rune(0)
			if c.Do() {
				c.Ret = QuoteRuneToGraphic(r)
			}
		},
	}, {
		Name: "IsPrint",
		Params: []internal.Param{
			{Kind: adt.IntKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			r := c.Rune(0)
			if c.Do() {
				c.Ret = IsPrint(r)
			}
		},
	}, {
		Name: "IsGraphic",
		Params: []internal.Param{
			{Kind: adt.IntKind},
		},
		Result: adt.BoolKind,
		Func: func(c *internal.CallCtxt) {
			r := c.Rune(0)
			if c.Do() {
				c.Ret = IsGraphic(r)
			}
		},
	}},
}
