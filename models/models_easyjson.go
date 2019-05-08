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

func easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels(in *jlexer.Lexer, out *UpdateInfo) {
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
		case "userId":
			out.UserId = string(in.String())
		case "timestamp":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
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
func easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels(out *jwriter.Writer, in UpdateInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserId))
	}
	{
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Timestamp).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels(l, v)
}
func easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels1(in *jlexer.Lexer, out *MemberInfo) {
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
		case "userId":
			out.UserId = string(in.String())
		case "privilege":
			out.Privilege = string(in.String())
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
func easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels1(out *jwriter.Writer, in MemberInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserId))
	}
	{
		const prefix string = ",\"privilege\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Privilege))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MemberInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MemberInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MemberInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MemberInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels1(l, v)
}
func easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels2(in *jlexer.Lexer, out *DocumentsInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(DocumentsInfo, 0, 1)
			} else {
				*out = DocumentsInfo{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 DocumentInfo
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels2(out *jwriter.Writer, in DocumentsInfo) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v DocumentsInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DocumentsInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DocumentsInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DocumentsInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels2(l, v)
}
func easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels3(in *jlexer.Lexer, out *DocumentMeta) {
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
		case "size":
			out.Size = int64(in.Int64())
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
func easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels3(out *jwriter.Writer, in DocumentMeta) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"size\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Size))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DocumentMeta) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DocumentMeta) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DocumentMeta) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DocumentMeta) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels3(l, v)
}
func easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels4(in *jlexer.Lexer, out *DocumentInfoFull) {
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
		case "membersInfo":
			if in.IsNull() {
				in.Skip()
				out.MembersInfo = nil
			} else {
				in.Delim('[')
				if out.MembersInfo == nil {
					if !in.IsDelim(']') {
						out.MembersInfo = make([]MemberInfo, 0, 2)
					} else {
						out.MembersInfo = []MemberInfo{}
					}
				} else {
					out.MembersInfo = (out.MembersInfo)[:0]
				}
				for !in.IsDelim(']') {
					var v4 MemberInfo
					(v4).UnmarshalEasyJSON(in)
					out.MembersInfo = append(out.MembersInfo, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "updateHistory":
			if in.IsNull() {
				in.Skip()
				out.UpdateHistory = nil
			} else {
				in.Delim('[')
				if out.UpdateHistory == nil {
					if !in.IsDelim(']') {
						out.UpdateHistory = make([]UpdateInfo, 0, 1)
					} else {
						out.UpdateHistory = []UpdateInfo{}
					}
				} else {
					out.UpdateHistory = (out.UpdateHistory)[:0]
				}
				for !in.IsDelim(']') {
					var v5 UpdateInfo
					(v5).UnmarshalEasyJSON(in)
					out.UpdateHistory = append(out.UpdateHistory, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "id":
			out.Id = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "ownerId":
			out.OwnerId = string(in.String())
		case "subject":
			out.Subject = string(in.String())
		case "meta":
			(out.Meta).UnmarshalEasyJSON(in)
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
func easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels4(out *jwriter.Writer, in DocumentInfoFull) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"membersInfo\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.MembersInfo == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.MembersInfo {
				if v6 > 0 {
					out.RawByte(',')
				}
				(v7).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"updateHistory\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.UpdateHistory == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.UpdateHistory {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"ownerId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OwnerId))
	}
	{
		const prefix string = ",\"subject\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Subject))
	}
	{
		const prefix string = ",\"meta\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Meta).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DocumentInfoFull) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DocumentInfoFull) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DocumentInfoFull) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DocumentInfoFull) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels4(l, v)
}
func easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels5(in *jlexer.Lexer, out *DocumentInfo) {
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
			out.Id = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "ownerId":
			out.OwnerId = string(in.String())
		case "subject":
			out.Subject = string(in.String())
		case "meta":
			(out.Meta).UnmarshalEasyJSON(in)
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
func easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels5(out *jwriter.Writer, in DocumentInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"ownerId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OwnerId))
	}
	{
		const prefix string = ",\"subject\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Subject))
	}
	{
		const prefix string = ",\"meta\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Meta).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DocumentInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DocumentInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComStudtoolDocumentsServiceModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DocumentInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DocumentInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComStudtoolDocumentsServiceModels5(l, v)
}
