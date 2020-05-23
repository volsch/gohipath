// Copyright (c) 2020, Volker Schmidt (volker@volsch.eu)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package pathsys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBooleanImplementsNegator(t *testing.T) {
	o := NewBoolean(false)
	assert.Implements(t, (*Negator)(nil), o)
}

func TestBooleanDataType(t *testing.T) {
	o := NewBoolean(false)
	dataType := o.DataType()
	assert.Equal(t, BooleanDataType, dataType)
}

func TestBooleanTypeInfo(t *testing.T) {
	o := NewBoolean(false)
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "System.Boolean", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "System.Any", i.FQBaseName().String())
		}
	}
}

func TestBooleanValue(t *testing.T) {
	o := NewBoolean(true)
	value := o.Bool()
	assert.Equal(t, true, value)
}

func TestParseBooleanTrue(t *testing.T) {
	o, err := ParseBoolean("true")

	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, o, "value expected") {
		assert.Equal(t, true, o.Bool())
		assert.Equal(t, "true", o.String())
	}
}

func TestParseBooleanFalse(t *testing.T) {
	o, err := ParseBoolean("false")

	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, o, "value expected") {
		assert.Equal(t, false, o.Bool())
		assert.Equal(t, "false", o.String())
	}
}

func TestParseBooleanInvalid(t *testing.T) {
	o, err := ParseBoolean("0")

	assert.Nil(t, o, "value unexpected")
	assert.NotNil(t, err, "error expected")
}

func TestBooleanNegateTrue(t *testing.T) {
	o := NewBoolean(true)
	n := o.Negate()
	assert.NotSame(t, o, n)
	assert.Equal(t, true, o.Bool())
	if assert.Implements(t, (*BooleanAccessor)(nil), n) {
		assert.Equal(t, false, n.(BooleanAccessor).Bool())
	}
}

func TestBooleanNegateFalse(t *testing.T) {
	o := NewBoolean(false)
	n := o.Negate()
	assert.NotSame(t, o, n)
	assert.Equal(t, false, o.Bool())
	if assert.Implements(t, (*BooleanAccessor)(nil), n) {
		assert.Equal(t, true, n.(BooleanAccessor).Bool())
	}
}

func TestBooleanEqualNil(t *testing.T) {
	assert.Equal(t, false, NewBoolean(false).Equal(nil))
}

func TestBooleanEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewBoolean(false).Equal(newAccessorMock()))
	assert.Equal(t, false, NewBoolean(false).Equivalent(newAccessorMock()))
}

func TestBooleanEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewBoolean(false).Equal(NewBoolean(false)))
	assert.Equal(t, true, NewBoolean(false).Equivalent(NewBoolean(false)))
}

func TestBooleanEqualNotEqual(t *testing.T) {
	assert.Equal(t, false, NewBoolean(false).Equal(NewBoolean(true)))
	assert.Equal(t, false, NewBoolean(false).Equivalent(NewBoolean(true)))
}