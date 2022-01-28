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

// ListSipTraces invokes the ccc.ListSipTraces API synchronously
func (client *Client) ListSipTraces(request *ListSipTracesRequest) (response *ListSipTracesResponse, err error) {
	response = CreateListSipTracesResponse()
	err = client.DoAction(request, response)
	return
}

// ListSipTracesWithChan invokes the ccc.ListSipTraces API asynchronously
func (client *Client) ListSipTracesWithChan(request *ListSipTracesRequest) (<-chan *ListSipTracesResponse, <-chan error) {
	responseChan := make(chan *ListSipTracesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSipTraces(request)
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

// ListSipTracesWithCallback invokes the ccc.ListSipTraces API asynchronously
func (client *Client) ListSipTracesWithCallback(request *ListSipTracesRequest, callback func(response *ListSipTracesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSipTracesResponse
		var err error
		defer close(result)
		response, err = client.ListSipTraces(request)
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

// ListSipTracesRequest is the request struct for api ListSipTraces
type ListSipTracesRequest struct {
	*requests.RpcRequest
	CallId     string `position:"Query" name:"CallId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListSipTracesResponse is the response struct for api ListSipTraces
type ListSipTracesResponse struct {
	*responses.BaseResponse
	Code           string       `json:"Code" xml:"Code"`
	HttpStatusCode int          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string       `json:"Message" xml:"Message"`
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Data           []CallTraces `json:"Data" xml:"Data"`
}

// CreateListSipTracesRequest creates a request to invoke ListSipTraces API
func CreateListSipTracesRequest() (request *ListSipTracesRequest) {
	request = &ListSipTracesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListSipTraces", "", "")
	request.Method = requests.POST
	return
}

// CreateListSipTracesResponse creates a response to parse from ListSipTraces response
func CreateListSipTracesResponse() (response *ListSipTracesResponse) {
	response = &ListSipTracesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
