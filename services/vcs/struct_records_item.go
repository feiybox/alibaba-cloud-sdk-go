package vcs

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

// RecordsItem is a nested struct in vcs response
type RecordsItem struct {
	LeftUpX         string        `json:"LeftUpX" xml:"LeftUpX"`
	RightBottomY    string        `json:"RightBottomY" xml:"RightBottomY"`
	LeftUpY         string        `json:"LeftUpY" xml:"LeftUpY"`
	FirstAppearTime string        `json:"FirstAppearTime" xml:"FirstAppearTime"`
	PersonId        string        `json:"PersonId" xml:"PersonId"`
	Score           string        `json:"Score" xml:"Score"`
	ShotTime        string        `json:"ShotTime" xml:"ShotTime"`
	TagCode         string        `json:"TagCode" xml:"TagCode"`
	GbId            string        `json:"GbId" xml:"GbId"`
	TagMetric       string        `json:"TagMetric" xml:"TagMetric"`
	DateTime        string        `json:"DateTime" xml:"DateTime"`
	TagValue        string        `json:"TagValue" xml:"TagValue"`
	PicUrl          string        `json:"PicUrl" xml:"PicUrl"`
	RightBottomX    string        `json:"RightBottomX" xml:"RightBottomX"`
	MonitorPicUrl   string        `json:"MonitorPicUrl" xml:"MonitorPicUrl"`
	TagList         []TagListItem `json:"TagList" xml:"TagList"`
}
