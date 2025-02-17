// Copyright 2022 EMQ Technologies Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package context

import (
	"fmt"
	"github.com/lf-edge/ekuiper/internal"
	"github.com/lf-edge/ekuiper/pkg/message"
)

func (c *DefaultContext) Decode(data []byte) (map[string]interface{}, error) {
	v := c.Value(internal.DecodeKey)
	f, ok := v.(message.Converter)
	if ok {
		t, err := f.Decode(data)
		if err != nil {
			return nil, fmt.Errorf("decode failed: %v", err)
		}
		if result, ok := t.(map[string]interface{}); ok {
			return result, nil
		} else {
			return nil, fmt.Errorf("only map[string]interface{} is supported but got: %v", t)
		}
	}
	return nil, fmt.Errorf("no decoder configured")
}
