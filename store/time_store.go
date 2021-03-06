// Copyright 2015 Google Inc. All Rights Reserved.
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

// Package cache implements a cache for time series data.
package store

import "time"

// TimeStore represents a timeseries storage.
// Implementations are expected to be thread safe.
type TimeStore interface {
	// Put stores 'data' with 'timestamp'. Returns error upon failure.
	// Ignores zero timestamps.
	Put(timestamp time.Time, data interface{}) error
	// Get returns a slice of elements that were previously stored with timestamps
	// between 'start' and 'end'. 'start' is expected to be before 'end'.
	// If 'start' is zero, it returns all the elements up until 'end'.
	// If 'end' is zero, it returns all the elements from 'start'.
	// If both 'start' and 'end' are zero, it returns all the elements in the cache.
	Get(start, end time.Time) []interface{}
	// Delete removes all elements that were previously stored with timestamps
	// between 'start' and 'end'.
	Delete(start, end time.Time) error
}
