package resourcemanager

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

// DeleteServiceLinkedRole invokes the resourcemanager.DeleteServiceLinkedRole API synchronously
// api document: https://help.aliyun.com/api/resourcemanager/deleteservicelinkedrole.html
func (client *Client) DeleteServiceLinkedRole(request *DeleteServiceLinkedRoleRequest) (response *DeleteServiceLinkedRoleResponse, err error) {
	response = CreateDeleteServiceLinkedRoleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteServiceLinkedRoleWithChan invokes the resourcemanager.DeleteServiceLinkedRole API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/deleteservicelinkedrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteServiceLinkedRoleWithChan(request *DeleteServiceLinkedRoleRequest) (<-chan *DeleteServiceLinkedRoleResponse, <-chan error) {
	responseChan := make(chan *DeleteServiceLinkedRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteServiceLinkedRole(request)
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

// DeleteServiceLinkedRoleWithCallback invokes the resourcemanager.DeleteServiceLinkedRole API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/deleteservicelinkedrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteServiceLinkedRoleWithCallback(request *DeleteServiceLinkedRoleRequest, callback func(response *DeleteServiceLinkedRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteServiceLinkedRoleResponse
		var err error
		defer close(result)
		response, err = client.DeleteServiceLinkedRole(request)
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

// DeleteServiceLinkedRoleRequest is the request struct for api DeleteServiceLinkedRole
type DeleteServiceLinkedRoleRequest struct {
	*requests.RpcRequest
	RoleName string `position:"Query" name:"RoleName"`
}

// DeleteServiceLinkedRoleResponse is the response struct for api DeleteServiceLinkedRole
type DeleteServiceLinkedRoleResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DeletionTaskId string `json:"DeletionTaskId" xml:"DeletionTaskId"`
}

// CreateDeleteServiceLinkedRoleRequest creates a request to invoke DeleteServiceLinkedRole API
func CreateDeleteServiceLinkedRoleRequest() (request *DeleteServiceLinkedRoleRequest) {
	request = &DeleteServiceLinkedRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "DeleteServiceLinkedRole", "resourcemanager", "openAPI")
	return
}

// CreateDeleteServiceLinkedRoleResponse creates a response to parse from DeleteServiceLinkedRole response
func CreateDeleteServiceLinkedRoleResponse() (response *DeleteServiceLinkedRoleResponse) {
	response = &DeleteServiceLinkedRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}