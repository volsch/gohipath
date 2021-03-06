// Copyright (c) 2020-2021, Volker Schmidt (volker@volsch.eu)
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

package internal

import (
	"github.com/healthiop/hipath/internal/expression"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVisitEqualityExpression(t *testing.T) {
	args := make([]interface{}, 3)
	args[0] = expression.NewEmptyLiteral()
	args[1] = "x"
	args[2] = expression.NewEmptyLiteral()
	res, err := visitEqualityExpression(nil, args)

	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no evaluator expected")
}

func TestVisitInequalityExpression(t *testing.T) {
	args := make([]interface{}, 3)
	args[0] = expression.NewEmptyLiteral()
	args[1] = "x"
	args[2] = expression.NewEmptyLiteral()
	res, err := visitInequalityExpression(nil, args)

	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no evaluator expected")
}

func TestVisitMembershipExpression(t *testing.T) {
	args := make([]interface{}, 3)
	args[0] = expression.NewEmptyLiteral()
	args[1] = "x"
	args[2] = expression.NewEmptyLiteral()
	res, err := visitMembershipExpression(nil, args)

	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no evaluator expected")
}

func TestVisitBooleanExpression(t *testing.T) {
	args := make([]interface{}, 3)
	args[0] = expression.NewEmptyLiteral()
	args[1] = "x"
	args[2] = expression.NewEmptyLiteral()
	res, err := visitBooleanExpression(nil, args)

	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no evaluator expected")
}

func TestVisitTypeExpression(t *testing.T) {
	args := make([]interface{}, 3)
	args[0] = expression.NewEmptyLiteral()
	args[1] = "x"
	args[2] = "String"
	res, err := visitTypeExpression(nil, args)

	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no evaluator expected")
}
