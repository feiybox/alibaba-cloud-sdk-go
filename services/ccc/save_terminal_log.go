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

// SaveTerminalLog invokes the ccc.SaveTerminalLog API synchronously
func (client *Client) SaveTerminalLog(request *SaveTerminalLogRequest) (response *SaveTerminalLogResponse, err error) {
	response = CreateSaveTerminalLogResponse()
	err = client.DoAction(request, response)
	return
}

// SaveTerminalLogWithChan invokes the ccc.SaveTerminalLog API asynchronously
func (client *Client) SaveTerminalLogWithChan(request *SaveTerminalLogRequest) (<-chan *SaveTerminalLogResponse, <-chan error) {
	responseChan := make(chan *SaveTerminalLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveTerminalLog(request)
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

// SaveTerminalLogWithCallback invokes the ccc.SaveTerminalLog API asynchronously
func (client *Client) SaveTerminalLogWithCallback(request *SaveTerminalLogRequest, callback func(response *SaveTerminalLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveTerminalLogResponse
		var err error
		defer close(result)
		response, err = client.SaveTerminalLog(request)
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

// SaveTerminalLogRequest is the request struct for api SaveTerminalLog
type SaveTerminalLogRequest struct {
	*requests.RpcRequest
	CallId          string           `position:"Query" name:"CallId"`
	Content         string           `position:"Query" name:"Content"`
	UniqueRequestId string           `position:"Query" name:"UniqueRequestId"`
	JobId           string           `position:"Query" name:"JobId"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	AppName         string           `position:"Query" name:"AppName"`
	DataType        requests.Integer `position:"Query" name:"DataType"`
	Status          string           `position:"Query" name:"Status"`
	MethodName      string           `position:"Query" name:"MethodName"`
}

// SaveTerminalLogResponse is the response struct for api SaveTerminalLog
type SaveTerminalLogResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	TimeStamp      int64  `json:"TimeStamp" xml:"TimeStamp"`
}

// CreateSaveTerminalLogRequest creates a request to invoke SaveTerminalLog API
func CreateSaveTerminalLogRequest() (request *SaveTerminalLogRequest) {
	request = &SaveTerminalLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "SaveTerminalLog", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveTerminalLogResponse creates a response to parse from SaveTerminalLog response
func CreateSaveTerminalLogResponse() (response *SaveTerminalLogResponse) {
	response = &SaveTerminalLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
