/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/LilShah/dummy-operator-v2/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"testing"
	//+kubebuilder:scaffold:imports
)

func getDummy(name string, namespace string, message string) *v1alpha1.Dummy {
	return &v1alpha1.Dummy{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1alpha1.DummySpec{
			Message: message,
		},
	}
}

const (
	testDummyName      = "dummy"
	testDummyNamespace = "default"
)

func TestDummyController(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"Hello from the other side", "Hello from the other side"},
		{"", ""},
		{"Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long", "Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long Long"},
		{"!@#$%^&*()_+|\"``~", "!@#$%^&*()_+|\"``~"},
	}

	for _, test := range tests {
		dummy := getDummy("dummy", "default", test.input)
		scheme := runtime.NewScheme()
		utilruntime.Must(v1alpha1.AddToScheme(scheme))

		fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(dummy).Build()
		r := &DummyReconciler{
			Client: fakeClient,
			Scheme: scheme,
			Log:    ctrl.Log.WithName("dummy"),
		}

		_, err := r.Reconcile(
			context.TODO(),
			reconcile.Request{
				NamespacedName: types.NamespacedName{
					Name:      dummy.Name,
					Namespace: dummy.Namespace,
				},
			},
		)
		require.NoError(t, err)

		err = r.Get(context.TODO(), types.NamespacedName{Name: dummy.Name, Namespace: dummy.Namespace}, dummy)
		assert.NoError(t, err)

		assert.Equal(t, test.expected, dummy.Spec.Message)
		assert.Equal(t, test.expected, dummy.Status.SpecEcho)
	}
}
