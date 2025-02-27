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

// Code generated by protoc-gen-go-apiclientconnect. DO NOT EDIT.

package registryv1alpha1apiclientconnect

import (
	context "context"
	registryv1alpha1connect "github.com/bufbuild/buf/private/gen/proto/connect/buf/alpha/registry/v1alpha1/registryv1alpha1connect"
	v1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/image/v1"
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
	connect_go "github.com/bufbuild/connect-go"
	zap "go.uber.org/zap"
	pluginpb "google.golang.org/protobuf/types/pluginpb"
)

type generateServiceClient struct {
	logger          *zap.Logger
	client          registryv1alpha1connect.GenerateServiceClient
	contextModifier func(context.Context) context.Context
}

// GeneratePlugins generates an array of files given the provided
// module reference and plugin version and option tuples. No attempt
// is made at merging insertion points.
func (s *generateServiceClient) GeneratePlugins(
	ctx context.Context,
	image *v1.Image,
	plugins []*v1alpha1.PluginReference,
	includeImports bool,
	includeWellKnownTypes bool,
) (responses []*pluginpb.CodeGeneratorResponse, runtimeLibraries []*v1alpha1.RuntimeLibrary, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GeneratePlugins(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.GeneratePluginsRequest{
				Image:                 image,
				Plugins:               plugins,
				IncludeImports:        includeImports,
				IncludeWellKnownTypes: includeWellKnownTypes,
			}),
	)
	if err != nil {
		return nil, nil, err
	}
	return response.Msg.Responses, response.Msg.RuntimeLibraries, nil
}

// GenerateTemplate generates an array of files given the provided
// module reference and template version.
func (s *generateServiceClient) GenerateTemplate(
	ctx context.Context,
	image *v1.Image,
	templateOwner string,
	templateName string,
	templateVersion string,
	includeImports bool,
	includeWellKnownTypes bool,
) (files []*v1alpha1.File, runtimeLibraries []*v1alpha1.RuntimeLibrary, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GenerateTemplate(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.GenerateTemplateRequest{
				Image:                 image,
				TemplateOwner:         templateOwner,
				TemplateName:          templateName,
				TemplateVersion:       templateVersion,
				IncludeImports:        includeImports,
				IncludeWellKnownTypes: includeWellKnownTypes,
			}),
	)
	if err != nil {
		return nil, nil, err
	}
	return response.Msg.Files, response.Msg.RuntimeLibraries, nil
}
