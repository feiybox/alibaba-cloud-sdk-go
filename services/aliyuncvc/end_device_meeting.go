package aliyuncvc

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

// EndDeviceMeeting invokes the aliyuncvc.EndDeviceMeeting API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/enddevicemeeting.html
func (client *Client) EndDeviceMeeting(request *EndDeviceMeetingRequest) (response *EndDeviceMeetingResponse, err error) {
	response = CreateEndDeviceMeetingResponse()
	err = client.DoAction(request, response)
	return
}

// EndDeviceMeetingWithChan invokes the aliyuncvc.EndDeviceMeeting API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/enddevicemeeting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EndDeviceMeetingWithChan(request *EndDeviceMeetingRequest) (<-chan *EndDeviceMeetingResponse, <-chan error) {
	responseChan := make(chan *EndDeviceMeetingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EndDeviceMeeting(request)
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

// EndDeviceMeetingWithCallback invokes the aliyuncvc.EndDeviceMeeting API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/enddevicemeeting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EndDeviceMeetingWithCallback(request *EndDeviceMeetingRequest, callback func(response *EndDeviceMeetingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EndDeviceMeetingResponse
		var err error
		defer close(result)
		response, err = client.EndDeviceMeeting(request)
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

// EndDeviceMeetingRequest is the request struct for api EndDeviceMeeting
type EndDeviceMeetingRequest struct {
	*requests.RpcRequest
	MeetingUUID string `position:"Body" name:"MeetingUUID"`
	SN          string `position:"Body" name:"SN"`
}

// EndDeviceMeetingResponse is the response struct for api EndDeviceMeeting
type EndDeviceMeetingResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEndDeviceMeetingRequest creates a request to invoke EndDeviceMeeting API
func CreateEndDeviceMeetingRequest() (request *EndDeviceMeetingRequest) {
	request = &EndDeviceMeetingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "EndDeviceMeeting", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEndDeviceMeetingResponse creates a response to parse from EndDeviceMeeting response
func CreateEndDeviceMeetingResponse() (response *EndDeviceMeetingResponse) {
	response = &EndDeviceMeetingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
