package ccc

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

// GetTurnServerList invokes the ccc.GetTurnServerList API synchronously
func (client *Client) GetTurnServerList(request *GetTurnServerListRequest) (response *GetTurnServerListResponse, err error) {
	response = CreateGetTurnServerListResponse()
	err = client.DoAction(request, response)
	return
}

// GetTurnServerListWithChan invokes the ccc.GetTurnServerList API asynchronously
func (client *Client) GetTurnServerListWithChan(request *GetTurnServerListRequest) (<-chan *GetTurnServerListResponse, <-chan error) {
	responseChan := make(chan *GetTurnServerListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTurnServerList(request)
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

// GetTurnServerListWithCallback invokes the ccc.GetTurnServerList API asynchronously
func (client *Client) GetTurnServerListWithCallback(request *GetTurnServerListRequest, callback func(response *GetTurnServerListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTurnServerListResponse
		var err error
		defer close(result)
		response, err = client.GetTurnServerList(request)
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

// GetTurnServerListRequest is the request struct for api GetTurnServerList
type GetTurnServerListRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetTurnServerListResponse is the response struct for api GetTurnServerList
type GetTurnServerListResponse struct {
	*responses.BaseResponse
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	Data           string   `json:"Data" xml:"Data"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
}

// CreateGetTurnServerListRequest creates a request to invoke GetTurnServerList API
func CreateGetTurnServerListRequest() (request *GetTurnServerListRequest) {
	request = &GetTurnServerListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "GetTurnServerList", "", "")
	request.Method = requests.POST
	return
}

// CreateGetTurnServerListResponse creates a response to parse from GetTurnServerList response
func CreateGetTurnServerListResponse() (response *GetTurnServerListResponse) {
	response = &GetTurnServerListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
