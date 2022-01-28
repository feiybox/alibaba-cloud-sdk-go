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

// ListIntervalAgentReport invokes the ccc.ListIntervalAgentReport API synchronously
func (client *Client) ListIntervalAgentReport(request *ListIntervalAgentReportRequest) (response *ListIntervalAgentReportResponse, err error) {
	response = CreateListIntervalAgentReportResponse()
	err = client.DoAction(request, response)
	return
}

// ListIntervalAgentReportWithChan invokes the ccc.ListIntervalAgentReport API asynchronously
func (client *Client) ListIntervalAgentReportWithChan(request *ListIntervalAgentReportRequest) (<-chan *ListIntervalAgentReportResponse, <-chan error) {
	responseChan := make(chan *ListIntervalAgentReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListIntervalAgentReport(request)
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

// ListIntervalAgentReportWithCallback invokes the ccc.ListIntervalAgentReport API asynchronously
func (client *Client) ListIntervalAgentReportWithCallback(request *ListIntervalAgentReportRequest, callback func(response *ListIntervalAgentReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListIntervalAgentReportResponse
		var err error
		defer close(result)
		response, err = client.ListIntervalAgentReport(request)
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

// ListIntervalAgentReportRequest is the request struct for api ListIntervalAgentReport
type ListIntervalAgentReportRequest struct {
	*requests.RpcRequest
	AgentId    string           `position:"Query" name:"AgentId"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	Interval   string           `position:"Query" name:"Interval"`
}

// ListIntervalAgentReportResponse is the response struct for api ListIntervalAgentReport
type ListIntervalAgentReportResponse struct {
	*responses.BaseResponse
	Code           string     `json:"Code" xml:"Code"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string     `json:"Message" xml:"Message"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Data           []DataItem `json:"Data" xml:"Data"`
}

// CreateListIntervalAgentReportRequest creates a request to invoke ListIntervalAgentReport API
func CreateListIntervalAgentReportRequest() (request *ListIntervalAgentReportRequest) {
	request = &ListIntervalAgentReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListIntervalAgentReport", "", "")
	request.Method = requests.POST
	return
}

// CreateListIntervalAgentReportResponse creates a response to parse from ListIntervalAgentReport response
func CreateListIntervalAgentReportResponse() (response *ListIntervalAgentReportResponse) {
	response = &ListIntervalAgentReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
