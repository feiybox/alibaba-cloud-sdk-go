package domain

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

// GetOperationOssUploadPolicy invokes the domain.GetOperationOssUploadPolicy API synchronously
func (client *Client) GetOperationOssUploadPolicy(request *GetOperationOssUploadPolicyRequest) (response *GetOperationOssUploadPolicyResponse, err error) {
	response = CreateGetOperationOssUploadPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// GetOperationOssUploadPolicyWithChan invokes the domain.GetOperationOssUploadPolicy API asynchronously
func (client *Client) GetOperationOssUploadPolicyWithChan(request *GetOperationOssUploadPolicyRequest) (<-chan *GetOperationOssUploadPolicyResponse, <-chan error) {
	responseChan := make(chan *GetOperationOssUploadPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOperationOssUploadPolicy(request)
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

// GetOperationOssUploadPolicyWithCallback invokes the domain.GetOperationOssUploadPolicy API asynchronously
func (client *Client) GetOperationOssUploadPolicyWithCallback(request *GetOperationOssUploadPolicyRequest, callback func(response *GetOperationOssUploadPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOperationOssUploadPolicyResponse
		var err error
		defer close(result)
		response, err = client.GetOperationOssUploadPolicy(request)
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

// GetOperationOssUploadPolicyRequest is the request struct for api GetOperationOssUploadPolicy
type GetOperationOssUploadPolicyRequest struct {
	*requests.RpcRequest
	AuditType requests.Integer `position:"Query" name:"AuditType"`
	Lang      string           `position:"Query" name:"Lang"`
}

// GetOperationOssUploadPolicyResponse is the response struct for api GetOperationOssUploadPolicy
type GetOperationOssUploadPolicyResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Accessid      string `json:"Accessid" xml:"Accessid"`
	EncodedPolicy string `json:"EncodedPolicy" xml:"EncodedPolicy"`
	Signature     string `json:"Signature" xml:"Signature"`
	FileDir       string `json:"FileDir" xml:"FileDir"`
	Host          string `json:"Host" xml:"Host"`
	ExpireTime    string `json:"ExpireTime" xml:"ExpireTime"`
}

// CreateGetOperationOssUploadPolicyRequest creates a request to invoke GetOperationOssUploadPolicy API
func CreateGetOperationOssUploadPolicyRequest() (request *GetOperationOssUploadPolicyRequest) {
	request = &GetOperationOssUploadPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "GetOperationOssUploadPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOperationOssUploadPolicyResponse creates a response to parse from GetOperationOssUploadPolicy response
func CreateGetOperationOssUploadPolicyResponse() (response *GetOperationOssUploadPolicyResponse) {
	response = &GetOperationOssUploadPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
