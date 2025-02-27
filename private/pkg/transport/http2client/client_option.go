// Copyright 2020-2022 Buf Technologies, Inc.
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

package http2client

import "crypto/tls"

// ClientOption is an option to modify the *http.Client.
type ClientOption func(*clientOption)

type clientOption struct {
	tlsConfig     *tls.Config
	useH2C        bool
	observability bool
}

// WithTLSConfig returns a new ClientOption to use the tls.Config.
//
// The default is to use no TLS.
func WithTLSConfig(tlsConfig *tls.Config) ClientOption {
	return func(option *clientOption) {
		option.tlsConfig = tlsConfig
	}
}

// WithH2C returns a new ClientOption that allows dialing
// h2c (cleartext) servers.
func WithH2C() ClientOption {
	return func(option *clientOption) {
		option.useH2C = true
	}
}

// WithObservability returns a new ClientOption to use
// OpenCensus tracing and metrics.
//
// The default is to use no observability.
func WithObservability() ClientOption {
	return func(option *clientOption) {
		option.observability = true
	}
}
