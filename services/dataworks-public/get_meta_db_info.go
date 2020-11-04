package dataworks_public

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

// GetMetaDBInfo invokes the dataworks_public.GetMetaDBInfo API synchronously
func (client *Client) GetMetaDBInfo(request *GetMetaDBInfoRequest) (response *GetMetaDBInfoResponse, err error) {
	response = CreateGetMetaDBInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetMetaDBInfoWithChan invokes the dataworks_public.GetMetaDBInfo API asynchronously
func (client *Client) GetMetaDBInfoWithChan(request *GetMetaDBInfoRequest) (<-chan *GetMetaDBInfoResponse, <-chan error) {
	responseChan := make(chan *GetMetaDBInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMetaDBInfo(request)
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

// GetMetaDBInfoWithCallback invokes the dataworks_public.GetMetaDBInfo API asynchronously
func (client *Client) GetMetaDBInfoWithCallback(request *GetMetaDBInfoRequest, callback func(response *GetMetaDBInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMetaDBInfoResponse
		var err error
		defer close(result)
		response, err = client.GetMetaDBInfo(request)
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

// GetMetaDBInfoRequest is the request struct for api GetMetaDBInfo
type GetMetaDBInfoRequest struct {
	*requests.RpcRequest
	DataSourceType string `position:"Query" name:"DataSourceType"`
	DatabaseName   string `position:"Query" name:"DatabaseName"`
	ClusterId      string `position:"Query" name:"ClusterId"`
	AppGuid        string `position:"Query" name:"AppGuid"`
}

// GetMetaDBInfoResponse is the response struct for api GetMetaDBInfo
type GetMetaDBInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetMetaDBInfoRequest creates a request to invoke GetMetaDBInfo API
func CreateGetMetaDBInfoRequest() (request *GetMetaDBInfoRequest) {
	request = &GetMetaDBInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetMetaDBInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateGetMetaDBInfoResponse creates a response to parse from GetMetaDBInfo response
func CreateGetMetaDBInfoResponse() (response *GetMetaDBInfoResponse) {
	response = &GetMetaDBInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}