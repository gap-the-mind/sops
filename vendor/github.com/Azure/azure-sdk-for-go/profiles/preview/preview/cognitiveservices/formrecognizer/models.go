// +build go1.9

// Copyright 2019 Microsoft Corporation
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

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package formrecognizer

import original "github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/v1.0/formrecognizer"

type Status = original.Status

const (
	Failure        Status = original.Failure
	PartialSuccess Status = original.PartialSuccess
	Success        Status = original.Success
)

type Status1 = original.Status1

const (
	Created Status1 = original.Created
	Invalid Status1 = original.Invalid
	Ready   Status1 = original.Ready
)

type Status2 = original.Status2

const (
	Status2Failure        Status2 = original.Status2Failure
	Status2PartialSuccess Status2 = original.Status2PartialSuccess
	Status2Success        Status2 = original.Status2Success
)

type AnalyzeResult = original.AnalyzeResult
type BaseClient = original.BaseClient
type ErrorInformation = original.ErrorInformation
type ErrorResponse = original.ErrorResponse
type ExtractedKeyValuePair = original.ExtractedKeyValuePair
type ExtractedPage = original.ExtractedPage
type ExtractedTable = original.ExtractedTable
type ExtractedTableColumn = original.ExtractedTableColumn
type ExtractedToken = original.ExtractedToken
type FormDocumentReport = original.FormDocumentReport
type FormOperationError = original.FormOperationError
type InnerError = original.InnerError
type KeysResult = original.KeysResult
type ModelResult = original.ModelResult
type ModelsResult = original.ModelsResult
type TrainRequest = original.TrainRequest
type TrainResult = original.TrainResult

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleStatus1Values() []Status1 {
	return original.PossibleStatus1Values()
}
func PossibleStatus2Values() []Status2 {
	return original.PossibleStatus2Values()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}