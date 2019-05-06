package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDedicatedHostTypes invokes the ecs.DescribeDedicatedHostTypes API synchronously
// api document: https://help.aliyun.com/api/ecs/describededicatedhosttypes.html
func (client *Client) DescribeDedicatedHostTypes(request *DescribeDedicatedHostTypesRequest) (response *DescribeDedicatedHostTypesResponse, err error) {
	response = CreateDescribeDedicatedHostTypesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDedicatedHostTypesWithChan invokes the ecs.DescribeDedicatedHostTypes API asynchronously
// api document: https://help.aliyun.com/api/ecs/describededicatedhosttypes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDedicatedHostTypesWithChan(request *DescribeDedicatedHostTypesRequest) (<-chan *DescribeDedicatedHostTypesResponse, <-chan error) {
	responseChan := make(chan *DescribeDedicatedHostTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDedicatedHostTypes(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeDedicatedHostTypesWithCallback invokes the ecs.DescribeDedicatedHostTypes API asynchronously
// api document: https://help.aliyun.com/api/ecs/describededicatedhosttypes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDedicatedHostTypesWithCallback(request *DescribeDedicatedHostTypesRequest, callback func(response *DescribeDedicatedHostTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDedicatedHostTypesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDedicatedHostTypes(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeDedicatedHostTypesRequest is the request struct for api DescribeDedicatedHostTypes
type DescribeDedicatedHostTypesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId             requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SupportedInstanceTypeFamily string           `position:"Query" name:"SupportedInstanceTypeFamily"`
	DedicatedHostType           string           `position:"Query" name:"DedicatedHostType"`
	ResourceOwnerAccount        string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string           `position:"Query" name:"OwnerAccount"`
	OwnerId                     requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDedicatedHostTypesResponse is the response struct for api DescribeDedicatedHostTypes
type DescribeDedicatedHostTypesResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	DedicatedHostTypes DedicatedHostTypes `json:"DedicatedHostTypes" xml:"DedicatedHostTypes"`
}

// CreateDescribeDedicatedHostTypesRequest creates a request to invoke DescribeDedicatedHostTypes API
func CreateDescribeDedicatedHostTypesRequest() (request *DescribeDedicatedHostTypesRequest) {
	request = &DescribeDedicatedHostTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDedicatedHostTypes", "ecs", "openAPI")
	return
}

// CreateDescribeDedicatedHostTypesResponse creates a response to parse from DescribeDedicatedHostTypes response
func CreateDescribeDedicatedHostTypesResponse() (response *DescribeDedicatedHostTypesResponse) {
	response = &DescribeDedicatedHostTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}