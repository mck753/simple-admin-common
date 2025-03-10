// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package i18n

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslator(t *testing.T) {
	l := &Translator{}
	l.NewBundle(LocaleFS)
	l.NewTranslator()
	res := l.Trans(context.WithValue(context.Background(), "lang", "zh"), "common.success")
	assert.Equal(t, "成功", res)
}

func TestNewTranslatorByPaths(t *testing.T) {
	l := NewTranslatorByPaths("locale/zh.json", "locale/en.json")

	res := l.Trans(context.WithValue(context.Background(), "lang", "zh"), "common.success")
	assert.Equal(t, "成功", res)

	res = l.Trans(context.WithValue(context.Background(), "lang", "en"), "common.success")
	assert.Equal(t, "Successfully", res)
}
