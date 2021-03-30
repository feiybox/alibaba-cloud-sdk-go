package green

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

// UpdateBizTypeSetting invokes the green.UpdateBizTypeSetting API synchronously
func (client *Client) UpdateBizTypeSetting(request *UpdateBizTypeSettingRequest) (response *UpdateBizTypeSettingResponse, err error) {
	response = CreateUpdateBizTypeSettingResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateBizTypeSettingWithChan invokes the green.UpdateBizTypeSetting API asynchronously
func (client *Client) UpdateBizTypeSettingWithChan(request *UpdateBizTypeSettingRequest) (<-chan *UpdateBizTypeSettingResponse, <-chan error) {
	responseChan := make(chan *UpdateBizTypeSettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateBizTypeSetting(request)
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

// UpdateBizTypeSettingWithCallback invokes the green.UpdateBizTypeSetting API asynchronously
func (client *Client) UpdateBizTypeSettingWithCallback(request *UpdateBizTypeSettingRequest, callback func(response *UpdateBizTypeSettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateBizTypeSettingResponse
		var err error
		defer close(result)
		response, err = client.UpdateBizTypeSetting(request)
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

// UpdateBizTypeSettingRequest is the request struct for api UpdateBizTypeSetting
type UpdateBizTypeSettingRequest struct {
	*requests.RpcRequest
	Antispam     string `position:"Query" name:"Antispam"`
	Porn         string `position:"Query" name:"Porn"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	Terrorism    string `position:"Query" name:"Terrorism"`
	BizTypeName  string `position:"Query" name:"BizTypeName"`
	Live         string `position:"Query" name:"Live"`
	Ad           string `position:"Query" name:"Ad"`
	ResourceType string `position:"Query" name:"ResourceType"`
}

// UpdateBizTypeSettingResponse is the response struct for api UpdateBizTypeSetting
type UpdateBizTypeSettingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateBizTypeSettingRequest creates a request to invoke UpdateBizTypeSetting API
func CreateUpdateBizTypeSettingRequest() (request *UpdateBizTypeSettingRequest) {
	request = &UpdateBizTypeSettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "UpdateBizTypeSetting", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateBizTypeSettingResponse creates a response to parse from UpdateBizTypeSetting response
func CreateUpdateBizTypeSettingResponse() (response *UpdateBizTypeSettingResponse) {
	response = &UpdateBizTypeSettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
