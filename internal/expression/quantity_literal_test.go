// Copyright (c) 2020-2021, Volker Schmidt (volker@volsch.eu)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source Unit must retain the above copyright notice, this
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

package expression

import (
	"github.com/healthiop/hipath/hipathsys"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuantityLiteralUCUM(t *testing.T) {
	evaluator, err := ParseQuantityLiteral("-17.4", "'ms'")

	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, evaluator, "evaluator expected") {
		res, err := evaluator.Evaluate(nil, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.NotNil(t, res, "res expected")
		if assert.Implements(t, (*hipathsys.QuantityAccessor)(nil), res) {
			quantityAccessor := res.(hipathsys.QuantityAccessor)
			if assert.NotNil(t, quantityAccessor.Value()) {
				assert.Equal(t, -17.4, quantityAccessor.Value().Float64())
			}
			if assert.NotNil(t, quantityAccessor.Unit()) {
				assert.Equal(t, "ms", quantityAccessor.Unit().String())
			}
		}
	}
}

func TestQuantityLiteralNoUCUM(t *testing.T) {
	evaluator, err := ParseQuantityLiteral("-17.4", "milliseconds")

	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, evaluator, "evaluator expected") {
		res, err := evaluator.Evaluate(nil, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.NotNil(t, res, "res expected")
		if assert.Implements(t, (*hipathsys.QuantityAccessor)(nil), res) {
			quantityAccessor := res.(hipathsys.QuantityAccessor)
			if assert.NotNil(t, quantityAccessor.Value()) {
				assert.Equal(t, -17.4, quantityAccessor.Value().Float64())
			}
			if assert.NotNil(t, quantityAccessor.Unit()) {
				assert.Equal(t, "milliseconds", quantityAccessor.Unit().String())
			}
		}
	}
}

func TestQuantityLiteralUnitEmpty(t *testing.T) {
	evaluator, err := ParseQuantityLiteral("-17.4", "")

	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, evaluator, "evaluator expected") {
		res, err := evaluator.Evaluate(nil, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.NotNil(t, res, "res expected")
		if assert.Implements(t, (*hipathsys.QuantityAccessor)(nil), res) {
			quantityAccessor := res.(hipathsys.QuantityAccessor)
			if assert.NotNil(t, quantityAccessor.Value()) {
				assert.Equal(t, -17.4, quantityAccessor.Value().Float64())
			}
			assert.Nil(t, quantityAccessor.Unit())
		}
	}
}

func TestQuantityLiteralUnitEmptyStringLiteral(t *testing.T) {
	evaluator, err := ParseQuantityLiteral("-17.4", "''")

	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, evaluator, "evaluator expected") {
		res, err := evaluator.Evaluate(nil, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.NotNil(t, res, "res expected")
		if assert.Implements(t, (*hipathsys.QuantityAccessor)(nil), res) {
			quantityAccessor := res.(hipathsys.QuantityAccessor)
			if assert.NotNil(t, quantityAccessor.Value()) {
				assert.Equal(t, -17.4, quantityAccessor.Value().Float64())
			}
			assert.Nil(t, quantityAccessor.Unit())
		}
	}
}

func TestQuantityLiteralValueInvalid(t *testing.T) {
	evaluator, err := ParseQuantityLiteral("-17.u", " 'test'")

	assert.Error(t, err, "error expected")
	assert.Nil(t, evaluator, "no evaluator expected")
}

func TestQuantityLiteralUnitInvalid(t *testing.T) {
	evaluator, err := ParseQuantityLiteral("-17.4", " 'test'")

	assert.Error(t, err, "error expected")
	assert.Nil(t, evaluator, "no evaluator expected")
}
