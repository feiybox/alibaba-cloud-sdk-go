package outboundbot

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

// SubmitScriptReview invokes the outboundbot.SubmitScriptReview API synchronously
func (client *Client) SubmitScriptReview(request *SubmitScriptReviewRequest) (response *SubmitScriptReviewResponse, err error) {
	response = CreateSubmitScriptReviewResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitScriptReviewWithChan invokes the outboundbot.SubmitScriptReview API asynchronously
func (client *Client) SubmitScriptReviewWithChan(request *SubmitScriptReviewRequest) (<-chan *SubmitScriptReviewResponse, <-chan error) {
	responseChan := make(chan *SubmitScriptReviewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitScriptReview(request)
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

// SubmitScriptReviewWithCallback invokes the outboundbot.SubmitScriptReview API asynchronously
func (client *Client) SubmitScriptReviewWithCallback(request *SubmitScriptReviewRequest, callback func(response *SubmitScriptReviewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitScriptReviewResponse
		var err error
		defer close(result)
		response, err = client.SubmitScriptReview(request)
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

// SubmitScriptReviewRequest is the request struct for api SubmitScriptReview
type SubmitScriptReviewRequest struct {
	*requests.RpcRequest
	Description string `position:"Query" name:"Description"`
	ScriptId    string `position:"Query" name:"ScriptId"`
	InstanceId  string `position:"Query" name:"InstanceId"`
}

// SubmitScriptReviewResponse is the response struct for api SubmitScriptReview
type SubmitScriptReviewResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateSubmitScriptReviewRequest creates a request to invoke SubmitScriptReview API
func CreateSubmitScriptReviewRequest() (request *SubmitScriptReviewRequest) {
	request = &SubmitScriptReviewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "SubmitScriptReview", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitScriptReviewResponse creates a response to parse from SubmitScriptReview response
func CreateSubmitScriptReviewResponse() (response *SubmitScriptReviewResponse) {
	response = &SubmitScriptReviewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
