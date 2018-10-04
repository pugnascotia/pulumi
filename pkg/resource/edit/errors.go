// Copyright 2016-2018, Pulumi Corporation.
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

package edit

import (
	"fmt"

	"github.com/pulumi/pulumi/pkg/resource"
)

// AmbiguousResourceError is returned by LocateResource if the given URN could possibly refer to more than one resource
// in the current snapshot. The `Resources` field contains all such resources in the snapshot.
type AmbiguousResourceError struct {
	URN       resource.URN
	Resources []*resource.State
}

func (a AmbiguousResourceError) Error() string {
	return fmt.Sprintf("URN %s is ambiguous among %d resources in the snapshot", a.URN, len(a.Resources))
}

// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it.
type ResourceHasDependenciesError struct {
	Condemned    *resource.State
	Dependencies []*resource.State
}

func (r ResourceHasDependenciesError) Error() string {
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)
}