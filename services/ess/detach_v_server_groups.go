package ess

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

// DetachVServerGroups invokes the ess.DetachVServerGroups API synchronously
// api document: https://help.aliyun.com/api/ess/detachvservergroups.html
func (client *Client) DetachVServerGroups(request *DetachVServerGroupsRequest) (response *DetachVServerGroupsResponse, err error) {
	response = CreateDetachVServerGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DetachVServerGroupsWithChan invokes the ess.DetachVServerGroups API asynchronously
// api document: https://help.aliyun.com/api/ess/detachvservergroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetachVServerGroupsWithChan(request *DetachVServerGroupsRequest) (<-chan *DetachVServerGroupsResponse, <-chan error) {
	responseChan := make(chan *DetachVServerGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachVServerGroups(request)
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

// DetachVServerGroupsWithCallback invokes the ess.DetachVServerGroups API asynchronously
// api document: https://help.aliyun.com/api/ess/detachvservergroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetachVServerGroupsWithCallback(request *DetachVServerGroupsRequest, callback func(response *DetachVServerGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachVServerGroupsResponse
		var err error
		defer close(result)
		response, err = client.DetachVServerGroups(request)
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

// DetachVServerGroupsRequest is the request struct for api DetachVServerGroups
type DetachVServerGroupsRequest struct {
	*requests.RpcRequest
	ScalingGroupId       string                             `position:"Query" name:"ScalingGroupId"`
	ResourceOwnerAccount string                             `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer                   `position:"Query" name:"OwnerId"`
	ForceDetach          requests.Boolean                   `position:"Query" name:"ForceDetach"`
	VServerGroup         *[]DetachVServerGroupsVServerGroup `position:"Query" name:"VServerGroup"  type:"Repeated"`
}

// DetachVServerGroupsVServerGroup is a repeated param struct in DetachVServerGroupsRequest
type DetachVServerGroupsVServerGroup struct {
	LoadBalancerId        string    `name:"LoadBalancerId"`
	VServerGroupAttribute *[]string `name:"VServerGroupAttribute" type:"Repeated"`
}

// DetachVServerGroupsResponse is the response struct for api DetachVServerGroups
type DetachVServerGroupsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachVServerGroupsRequest creates a request to invoke DetachVServerGroups API
func CreateDetachVServerGroupsRequest() (request *DetachVServerGroupsRequest) {
	request = &DetachVServerGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DetachVServerGroups", "ess", "openAPI")
	return
}

// CreateDetachVServerGroupsResponse creates a response to parse from DetachVServerGroups response
func CreateDetachVServerGroupsResponse() (response *DetachVServerGroupsResponse) {
	response = &DetachVServerGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
