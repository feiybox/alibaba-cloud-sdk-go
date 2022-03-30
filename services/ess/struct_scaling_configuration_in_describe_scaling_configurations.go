package ess

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

// ScalingConfigurationInDescribeScalingConfigurations is a nested struct in ess response
type ScalingConfigurationInDescribeScalingConfigurations struct {
	DeploymentSetId                 string                              `json:"DeploymentSetId" xml:"DeploymentSetId"`
	CreationTime                    string                              `json:"CreationTime" xml:"CreationTime"`
	ScalingConfigurationName        string                              `json:"ScalingConfigurationName" xml:"ScalingConfigurationName"`
	SystemDiskDescription           string                              `json:"SystemDiskDescription" xml:"SystemDiskDescription"`
	KeyPairName                     string                              `json:"KeyPairName" xml:"KeyPairName"`
	SecurityGroupId                 string                              `json:"SecurityGroupId" xml:"SecurityGroupId"`
	PrivatePoolOptionsId            string                              `json:"PrivatePoolOptions.Id" xml:"PrivatePoolOptions.Id"`
	SystemDiskAutoSnapshotPolicyId  string                              `json:"SystemDiskAutoSnapshotPolicyId" xml:"SystemDiskAutoSnapshotPolicyId"`
	SpotStrategy                    string                              `json:"SpotStrategy" xml:"SpotStrategy"`
	ScalingGroupId                  string                              `json:"ScalingGroupId" xml:"ScalingGroupId"`
	Affinity                        string                              `json:"Affinity" xml:"Affinity"`
	Tenancy                         string                              `json:"Tenancy" xml:"Tenancy"`
	SystemDiskSize                  int                                 `json:"SystemDiskSize" xml:"SystemDiskSize"`
	Ipv6AddressCount                int                                 `json:"Ipv6AddressCount" xml:"Ipv6AddressCount"`
	SpotDuration                    int                                 `json:"SpotDuration" xml:"SpotDuration"`
	LifecycleState                  string                              `json:"LifecycleState" xml:"LifecycleState"`
	InstanceName                    string                              `json:"InstanceName" xml:"InstanceName"`
	SecurityEnhancementStrategy     string                              `json:"SecurityEnhancementStrategy" xml:"SecurityEnhancementStrategy"`
	UserData                        string                              `json:"UserData" xml:"UserData"`
	PrivatePoolOptionsMatchCriteria string                              `json:"PrivatePoolOptions.MatchCriteria" xml:"PrivatePoolOptions.MatchCriteria"`
	DedicatedHostId                 string                              `json:"DedicatedHostId" xml:"DedicatedHostId"`
	InstanceGeneration              string                              `json:"InstanceGeneration" xml:"InstanceGeneration"`
	HpcClusterId                    string                              `json:"HpcClusterId" xml:"HpcClusterId"`
	PasswordInherit                 bool                                `json:"PasswordInherit" xml:"PasswordInherit"`
	Memory                          int                                 `json:"Memory" xml:"Memory"`
	ImageId                         string                              `json:"ImageId" xml:"ImageId"`
	ImageFamily                     string                              `json:"ImageFamily" xml:"ImageFamily"`
	LoadBalancerWeight              int                                 `json:"LoadBalancerWeight" xml:"LoadBalancerWeight"`
	SystemDiskCategory              string                              `json:"SystemDiskCategory" xml:"SystemDiskCategory"`
	HostName                        string                              `json:"HostName" xml:"HostName"`
	SystemDiskName                  string                              `json:"SystemDiskName" xml:"SystemDiskName"`
	InternetMaxBandwidthOut         int                                 `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
	InternetMaxBandwidthIn          int                                 `json:"InternetMaxBandwidthIn" xml:"InternetMaxBandwidthIn"`
	InstanceType                    string                              `json:"InstanceType" xml:"InstanceType"`
	InstanceDescription             string                              `json:"InstanceDescription" xml:"InstanceDescription"`
	IoOptimized                     string                              `json:"IoOptimized" xml:"IoOptimized"`
	RamRoleName                     string                              `json:"RamRoleName" xml:"RamRoleName"`
	SystemDiskPerformanceLevel      string                              `json:"SystemDiskPerformanceLevel" xml:"SystemDiskPerformanceLevel"`
	Cpu                             int                                 `json:"Cpu" xml:"Cpu"`
	ResourceGroupId                 string                              `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ZoneId                          string                              `json:"ZoneId" xml:"ZoneId"`
	InternetChargeType              string                              `json:"InternetChargeType" xml:"InternetChargeType"`
	ImageName                       string                              `json:"ImageName" xml:"ImageName"`
	ScalingConfigurationId          string                              `json:"ScalingConfigurationId" xml:"ScalingConfigurationId"`
	CreditSpecification             string                              `json:"CreditSpecification" xml:"CreditSpecification"`
	SpotInterruptionBehavior        string                              `json:"SpotInterruptionBehavior" xml:"SpotInterruptionBehavior"`
	SystemDiskCategories            SystemDiskCategories                `json:"SystemDiskCategories" xml:"SystemDiskCategories"`
	WeightedCapacities              WeightedCapacities                  `json:"WeightedCapacities" xml:"WeightedCapacities"`
	InstanceTypes                   InstanceTypes                       `json:"InstanceTypes" xml:"InstanceTypes"`
	SecurityGroupIds                SecurityGroupIds                    `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
	SchedulerOptions                SchedulerOptions                    `json:"SchedulerOptions" xml:"SchedulerOptions"`
	DataDisks                       DataDisks                           `json:"DataDisks" xml:"DataDisks"`
	Tags                            TagsInDescribeScalingConfigurations `json:"Tags" xml:"Tags"`
	SpotPriceLimit                  SpotPriceLimit                      `json:"SpotPriceLimit" xml:"SpotPriceLimit"`
	InstancePatternInfos            InstancePatternInfos                `json:"InstancePatternInfos" xml:"InstancePatternInfos"`
}
