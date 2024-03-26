// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/containers/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *CommitOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *CommitOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithAuthor set field Author to given value
func (o *CommitOptions) WithAuthor(value string) *CommitOptions {
	o.Author = &value
	return o
}

// GetAuthor returns value of field Author
func (o *CommitOptions) GetAuthor() string {
	if o.Author == nil {
		var z string
		return z
	}
	return *o.Author
}

// WithChanges set field Changes to given value
func (o *CommitOptions) WithChanges(value []string) *CommitOptions {
	o.Changes = value
	return o
}

// GetChanges returns value of field Changes
func (o *CommitOptions) GetChanges() []string {
	if o.Changes == nil {
		var z []string
		return z
	}
	return o.Changes
}

// WithComment set field Comment to given value
func (o *CommitOptions) WithComment(value string) *CommitOptions {
	o.Comment = &value
	return o
}

// GetComment returns value of field Comment
func (o *CommitOptions) GetComment() string {
	if o.Comment == nil {
		var z string
		return z
	}
	return *o.Comment
}

// WithFormat set field Format to given value
func (o *CommitOptions) WithFormat(value string) *CommitOptions {
	o.Format = &value
	return o
}

// GetFormat returns value of field Format
func (o *CommitOptions) GetFormat() string {
	if o.Format == nil {
		var z string
		return z
	}
	return *o.Format
}

// WithPause set field Pause to given value
func (o *CommitOptions) WithPause(value bool) *CommitOptions {
	o.Pause = &value
	return o
}

// GetPause returns value of field Pause
func (o *CommitOptions) GetPause() bool {
	if o.Pause == nil {
		var z bool
		return z
	}
	return *o.Pause
}

// WithStream set field Stream to given value
func (o *CommitOptions) WithStream(value bool) *CommitOptions {
	o.Stream = &value
	return o
}

// GetStream returns value of field Stream
func (o *CommitOptions) GetStream() bool {
	if o.Stream == nil {
		var z bool
		return z
	}
	return *o.Stream
}

// WithSquash set field Squash to given value
func (o *CommitOptions) WithSquash(value bool) *CommitOptions {
	o.Squash = &value
	return o
}

// GetSquash returns value of field Squash
func (o *CommitOptions) GetSquash() bool {
	if o.Squash == nil {
		var z bool
		return z
	}
	return *o.Squash
}

// WithRepo set field Repo to given value
func (o *CommitOptions) WithRepo(value string) *CommitOptions {
	o.Repo = &value
	return o
}

// GetRepo returns value of field Repo
func (o *CommitOptions) GetRepo() string {
	if o.Repo == nil {
		var z string
		return z
	}
	return *o.Repo
}

// WithTag set field Tag to given value
func (o *CommitOptions) WithTag(value string) *CommitOptions {
	o.Tag = &value
	return o
}

// GetTag returns value of field Tag
func (o *CommitOptions) GetTag() string {
	if o.Tag == nil {
		var z string
		return z
	}
	return *o.Tag
}