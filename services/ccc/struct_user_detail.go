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

// UserDetail is a nested struct in ccc response
type UserDetail struct {
	DisplayName                string                      `json:"DisplayName" xml:"DisplayName"`
	LoginName                  string                      `json:"LoginName" xml:"LoginName"`
	Email                      string                      `json:"Email" xml:"Email"`
	WorkMode                   string                      `json:"WorkMode" xml:"WorkMode"`
	Mobile                     string                      `json:"Mobile" xml:"Mobile"`
	UserId                     string                      `json:"UserId" xml:"UserId"`
	DisplayId                  string                      `json:"DisplayId" xml:"DisplayId"`
	RoleName                   string                      `json:"RoleName" xml:"RoleName"`
	RoleId                     string                      `json:"RoleId" xml:"RoleId"`
	PrimaryAccount             bool                        `json:"PrimaryAccount" xml:"PrimaryAccount"`
	RamId                      int64                       `json:"RamId" xml:"RamId"`
	Extension                  string                      `json:"Extension" xml:"Extension"`
	DeviceId                   string                      `json:"DeviceId" xml:"DeviceId"`
	DeviceExt                  string                      `json:"DeviceExt" xml:"DeviceExt"`
	DeviceState                string                      `json:"DeviceState" xml:"DeviceState"`
	PersonalOutboundNumberList []PhoneNumber               `json:"PersonalOutboundNumberList" xml:"PersonalOutboundNumberList"`
	SkillLevelList             []UserSkillLevelInListUsers `json:"SkillLevelList" xml:"SkillLevelList"`
}
