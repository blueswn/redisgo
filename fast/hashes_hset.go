// Copyright 2019 blueswn
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package fast

// Set redis strings set value
//func (hashes *Hashes) HSet(field string, value string) error {
//	conn := hashes.Conn()
//	if hashes.isAutoClose {
//		defer conn.Close()
//	}
//
//	key := hashes.statement.GetKey()
//	if key == "" {
//		return ErrKeyEmpty
//	}
//
//	_, err := conn.Do("HSET",  key, field, value)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

