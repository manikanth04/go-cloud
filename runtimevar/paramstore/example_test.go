// Copyright 2018 The Go Cloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package paramstore_test

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/paramstore"
)

// MyConfig is a sample configuration struct.
type MyConfig struct {
	Server string
	Port   int
}

func ExampleNewVariable() {
	// Establish an AWS session.
	session, err := session.NewSession(&aws.Config{})

	// Create a decoder for decoding JSON strings into MyConfig.
	decoder := runtimevar.NewDecoder(MyConfig{}, runtimevar.JSONDecode)

	// Construct a *runtimevar.Variable that watches the variable.
	// For this example, the Parameter Store variable being referenced
	// should have a JSON string that decodes into MyConfig.
	v, err := paramstore.NewVariable(session, "myconfig", decoder, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer v.Close()

	// You can now read the current value of the variable from v.
	snapshot, err := v.Watch(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// The resulting runtimevar.Snapshot.Value will be of type MyConfig.
	log.Printf("Snapshot.Value: %#v", snapshot.Value.(MyConfig))
}
