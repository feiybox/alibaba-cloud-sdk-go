package reid

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

// DescribeCustomerFlowByLocation invokes the reid.DescribeCustomerFlowByLocation API synchronously
// api document: https://help.aliyun.com/api/reid/describecustomerflowbylocation.html
func (client *Client) DescribeCustomerFlowByLocation(request *DescribeCustomerFlowByLocationRequest) (response *DescribeCustomerFlowByLocationResponse, err error) {
	response = CreateDescribeCustomerFlowByLocationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCustomerFlowByLocationWithChan invokes the reid.DescribeCustomerFlowByLocation API asynchronously
// api document: https://help.aliyun.com/api/reid/describecustomerflowbylocation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCustomerFlowByLocationWithChan(request *DescribeCustomerFlowByLocationRequest) (<-chan *DescribeCustomerFlowByLocationResponse, <-chan error) {
	responseChan := make(chan *DescribeCustomerFlowByLocationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCustomerFlowByLocation(request)
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

// DescribeCustomerFlowByLocationWithCallback invokes the reid.DescribeCustomerFlowByLocation API asynchronously
// api document: https://help.aliyun.com/api/reid/describecustomerflowbylocation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCustomerFlowByLocationWithCallback(request *DescribeCustomerFlowByLocationRequest, callback func(response *DescribeCustomerFlowByLocationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCustomerFlowByLocationResponse
		var err error
		defer close(result)
		response, err = client.DescribeCustomerFlowByLocation(request)
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

// DescribeCustomerFlowByLocationRequest is the request struct for api DescribeCustomerFlowByLocation
type DescribeCustomerFlowByLocationRequest struct {
	*requests.RpcRequest
	StartDate         string           `position:"Body" name:"StartDate"`
	StoreId           requests.Integer `position:"Body" name:"StoreId"`
	MinCount          requests.Integer `position:"Body" name:"MinCount"`
	ParentAmount      requests.Integer `position:"Body" name:"ParentAmount"`
	MaxCount          requests.Integer `position:"Body" name:"MaxCount"`
	EndDate           string           `position:"Body" name:"EndDate"`
	LocationId        requests.Integer `position:"Body" name:"LocationId"`
	ParentLocationIds string           `position:"Body" name:"ParentLocationIds"`
}

// DescribeCustomerFlowByLocationResponse is the response struct for api DescribeCustomerFlowByLocation
type DescribeCustomerFlowByLocationResponse struct {
	*responses.BaseResponse
	LocationId        int64             `json:"LocationId" xml:"LocationId"`
	ParentLocationIds string            `json:"ParentLocationIds" xml:"ParentLocationIds"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Success           bool              `json:"Success" xml:"Success"`
	ErrorCode         string            `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage      string            `json:"ErrorMessage" xml:"ErrorMessage"`
	LocationName      string            `json:"LocationName" xml:"LocationName"`
	Percent           float64           `json:"Percent" xml:"Percent"`
	StoreId           int64             `json:"StoreId" xml:"StoreId"`
	Count             int64             `json:"Count" xml:"Count"`
	CustomerFlowItems CustomerFlowItems `json:"CustomerFlowItems" xml:"CustomerFlowItems"`
}

// CreateDescribeCustomerFlowByLocationRequest creates a request to invoke DescribeCustomerFlowByLocation API
func CreateDescribeCustomerFlowByLocationRequest() (request *DescribeCustomerFlowByLocationRequest) {
	request = &DescribeCustomerFlowByLocationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "DescribeCustomerFlowByLocation", "1.1.2", "openAPI")
	return
}

// CreateDescribeCustomerFlowByLocationResponse creates a response to parse from DescribeCustomerFlowByLocation response
func CreateDescribeCustomerFlowByLocationResponse() (response *DescribeCustomerFlowByLocationResponse) {
	response = &DescribeCustomerFlowByLocationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
