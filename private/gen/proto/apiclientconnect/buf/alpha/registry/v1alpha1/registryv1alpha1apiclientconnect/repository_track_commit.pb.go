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
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
	connect_go "github.com/bufbuild/connect-go"
	zap "go.uber.org/zap"
)

type repositoryTrackCommitServiceClient struct {
	logger          *zap.Logger
	client          registryv1alpha1connect.RepositoryTrackCommitServiceClient
	contextModifier func(context.Context) context.Context
}

// GetRepositoryTrackCommitByRepositoryCommit returns the RepositoryTrackCommit associated given repository_commit on
// the given repository_track. Returns NOT_FOUND if the RepositoryTrackCommit does not exist.
func (s *repositoryTrackCommitServiceClient) GetRepositoryTrackCommitByRepositoryCommit(
	ctx context.Context,
	repositoryTrackId string,
	repositoryCommitId string,
) (repositoryTrackCommit *v1alpha1.RepositoryTrackCommit, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetRepositoryTrackCommitByRepositoryCommit(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.GetRepositoryTrackCommitByRepositoryCommitRequest{
				RepositoryTrackId:  repositoryTrackId,
				RepositoryCommitId: repositoryCommitId,
			}),
	)
	if err != nil {
		return nil, err
	}
	return response.Msg.RepositoryTrackCommit, nil
}

// ListRepositoryTrackCommitsByRepositoryTrack lists the RepositoryTrackCommitS associated with a repository track,
// ordered by their sequence id.
func (s *repositoryTrackCommitServiceClient) ListRepositoryTrackCommitsByRepositoryTrack(
	ctx context.Context,
	repositoryTrackId string,
	pageSize uint32,
	pageToken string,
	reverse bool,
) (repositoryTrackCommits []*v1alpha1.RepositoryTrackCommit, nextPageToken string, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListRepositoryTrackCommitsByRepositoryTrack(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.ListRepositoryTrackCommitsByRepositoryTrackRequest{
				RepositoryTrackId: repositoryTrackId,
				PageSize:          pageSize,
				PageToken:         pageToken,
				Reverse:           reverse,
			}),
	)
	if err != nil {
		return nil, "", err
	}
	return response.Msg.RepositoryTrackCommits, response.Msg.NextPageToken, nil
}

// GetRepositoryTrackCommitByReference returns the RepositoryTrackCommit associated with the given reference.
func (s *repositoryTrackCommitServiceClient) GetRepositoryTrackCommitByReference(
	ctx context.Context,
	repositoryOwner string,
	repositoryName string,
	track string,
	reference string,
) (repositoryTrackCommit *v1alpha1.RepositoryTrackCommit, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetRepositoryTrackCommitByReference(
		ctx,
		connect_go.NewRequest(
			&v1alpha1.GetRepositoryTrackCommitByReferenceRequest{
				RepositoryOwner: repositoryOwner,
				RepositoryName:  repositoryName,
				Track:           track,
				Reference:       reference,
			}),
	)
	if err != nil {
		return nil, err
	}
	return response.Msg.RepositoryTrackCommit, nil
}
