// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeBackendInternalModels(in *jlexer.Lexer, out *ProductShrinkedArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ProductShrinkedArr, 0, 8)
			} else {
				*out = ProductShrinkedArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *ProductShrinked
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(ProductShrinked)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*v1).UnmarshalJSON(data))
				}
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeBackendInternalModels(out *jwriter.Writer, in ProductShrinkedArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				out.Raw((*v3).MarshalJSON())
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ProductShrinkedArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeBackendInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProductShrinkedArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeBackendInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProductShrinkedArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeBackendInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProductShrinkedArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeBackendInternalModels(l, v)
}
func easyjsonD2b7633eDecodeBackendInternalModels1(in *jlexer.Lexer, out *ProductExtendedArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ProductExtendedArr, 0, 8)
			} else {
				*out = ProductExtendedArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 *ProductExtended
			if in.IsNull() {
				in.Skip()
				v4 = nil
			} else {
				if v4 == nil {
					v4 = new(ProductExtended)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*v4).UnmarshalJSON(data))
				}
			}
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeBackendInternalModels1(out *jwriter.Writer, in ProductExtendedArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				out.Raw((*v6).MarshalJSON())
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ProductExtendedArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeBackendInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProductExtendedArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeBackendInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProductExtendedArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeBackendInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProductExtendedArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeBackendInternalModels1(l, v)
}
func easyjsonD2b7633eDecodeBackendInternalModels2(in *jlexer.Lexer, out *ProductShrinked) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "shrinked":
			out.Shrinked = bool(in.Bool())
		case "barcode":
			out.Barcode = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeBackendInternalModels2(out *jwriter.Writer, in ProductShrinked) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"shrinked\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Shrinked))
	}
	{
		const prefix string = ",\"barcode\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Barcode))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProductShrinked) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeBackendInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProductShrinked) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeBackendInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProductShrinked) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeBackendInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProductShrinked) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeBackendInternalModels2(l, v)
}
func easyjsonD2b7633eDecodeBackendInternalModels3(in *jlexer.Lexer, out *ProductExtended) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "shrinked":
			out.Shrinked = bool(in.Bool())
		case "barcode":
			out.Barcode = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "contents":
			out.Contents = string(in.String())
		case "category_url":
			out.CategoryURL = string(in.String())
		case "mass":
			out.Mass = string(in.String())
		case "best_before":
			out.BestBefore = string(in.String())
		case "nutrition":
			out.Nutrition = string(in.String())
		case "manufacturer":
			out.Manufacturer = string(in.String())
		case "image":
			out.Image = string(in.String())
		case "ingredients":
			if m, ok := out.Ingredients.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Ingredients.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Ingredients = in.Interface()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeBackendInternalModels3(out *jwriter.Writer, in ProductExtended) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"shrinked\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Shrinked))
	}
	{
		const prefix string = ",\"barcode\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Barcode))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"contents\":"
		out.RawString(prefix)
		out.String(string(in.Contents))
	}
	{
		const prefix string = ",\"category_url\":"
		out.RawString(prefix)
		out.String(string(in.CategoryURL))
	}
	{
		const prefix string = ",\"mass\":"
		out.RawString(prefix)
		out.String(string(in.Mass))
	}
	{
		const prefix string = ",\"best_before\":"
		out.RawString(prefix)
		out.String(string(in.BestBefore))
	}
	{
		const prefix string = ",\"nutrition\":"
		out.RawString(prefix)
		out.String(string(in.Nutrition))
	}
	{
		const prefix string = ",\"manufacturer\":"
		out.RawString(prefix)
		out.String(string(in.Manufacturer))
	}
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix)
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"ingredients\":"
		out.RawString(prefix)
		if m, ok := in.Ingredients.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Ingredients.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Ingredients))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProductExtended) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeBackendInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProductExtended) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeBackendInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProductExtended) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeBackendInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProductExtended) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeBackendInternalModels3(l, v)
}
func easyjsonD2b7633eDecodeBackendInternalModels4(in *jlexer.Lexer, out *IngredientArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(IngredientArr, 0, 8)
			} else {
				*out = IngredientArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v7 *Ingredient
			if in.IsNull() {
				in.Skip()
				v7 = nil
			} else {
				if v7 == nil {
					v7 = new(Ingredient)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*v7).UnmarshalJSON(data))
				}
			}
			*out = append(*out, v7)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeBackendInternalModels4(out *jwriter.Writer, in IngredientArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in {
			if v8 > 0 {
				out.RawByte(',')
			}
			if v9 == nil {
				out.RawString("null")
			} else {
				out.Raw((*v9).MarshalJSON())
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v IngredientArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeBackendInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IngredientArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeBackendInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IngredientArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeBackendInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IngredientArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeBackendInternalModels4(l, v)
}
func easyjsonD2b7633eDecodeBackendInternalModels5(in *jlexer.Lexer, out *Ingredient) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "danger":
			out.Danger = int(in.Int())
		case "description":
			out.Description = string(in.String())
		case "wiki_link":
			out.WikiLink = string(in.String())
		case "groups":
			if in.IsNull() {
				in.Skip()
				out.Groups = nil
			} else {
				in.Delim('[')
				if out.Groups == nil {
					if !in.IsDelim(']') {
						out.Groups = make([]int64, 0, 8)
					} else {
						out.Groups = []int64{}
					}
				} else {
					out.Groups = (out.Groups)[:0]
				}
				for !in.IsDelim(']') {
					var v10 int64
					v10 = int64(in.Int64())
					out.Groups = append(out.Groups, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeBackendInternalModels5(out *jwriter.Writer, in Ingredient) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"danger\":"
		out.RawString(prefix)
		out.Int(int(in.Danger))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"wiki_link\":"
		out.RawString(prefix)
		out.String(string(in.WikiLink))
	}
	{
		const prefix string = ",\"groups\":"
		out.RawString(prefix)
		if in.Groups == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Groups {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Ingredient) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeBackendInternalModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Ingredient) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeBackendInternalModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Ingredient) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeBackendInternalModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Ingredient) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeBackendInternalModels5(l, v)
}
func easyjsonD2b7633eDecodeBackendInternalModels6(in *jlexer.Lexer, out *GroupArr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GroupArr, 0, 8)
			} else {
				*out = GroupArr{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v13 *Group
			if in.IsNull() {
				in.Skip()
				v13 = nil
			} else {
				if v13 == nil {
					v13 = new(Group)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*v13).UnmarshalJSON(data))
				}
			}
			*out = append(*out, v13)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeBackendInternalModels6(out *jwriter.Writer, in GroupArr) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v14, v15 := range in {
			if v14 > 0 {
				out.RawByte(',')
			}
			if v15 == nil {
				out.RawString("null")
			} else {
				out.Raw((*v15).MarshalJSON())
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GroupArr) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeBackendInternalModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GroupArr) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeBackendInternalModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GroupArr) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeBackendInternalModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GroupArr) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeBackendInternalModels6(l, v)
}
func easyjsonD2b7633eDecodeBackendInternalModels7(in *jlexer.Lexer, out *Group) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "about":
			out.About = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeBackendInternalModels7(out *jwriter.Writer, in Group) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Group) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeBackendInternalModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Group) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeBackendInternalModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Group) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeBackendInternalModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Group) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeBackendInternalModels7(l, v)
}
func easyjsonD2b7633eDecodeBackendInternalModels8(in *jlexer.Lexer, out *Msg) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message":
			out.Message = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeBackendInternalModels8(out *jwriter.Writer, in Msg) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Msg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeBackendInternalModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Msg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeBackendInternalModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Msg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeBackendInternalModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Msg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeBackendInternalModels8(l, v)
}
func easyjsonD2b7633eDecodeBackendInternalModels9(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int32(in.Int32())
		case "username":
			out.Username = string(in.String())
		case "password":
			out.Password = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeBackendInternalModels9(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.ID))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeBackendInternalModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeBackendInternalModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeBackendInternalModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeBackendInternalModels9(l, v)
}
