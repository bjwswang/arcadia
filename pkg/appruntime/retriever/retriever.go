/*
Copyright 2023 KubeAGI.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package retriever

import "encoding/json"

type Reference struct {
	// Question row
	Question string `json:"question" example:"q: 旷工最小计算单位为多少天？"`
	// Answer row
	Answer string `json:"answer" example:"旷工最小计算单位为 0.5 天。"`
	// vector search score
	Score float32 `json:"score" example:"0.34"`
	// the qa file fullpath
	QAFilePath string `json:"qa_file_path" example:"dataset/dataset-playground/v1/qa.csv"`
	// line number in the qa file
	QALineNumber int `json:"qa_line_number" example:"7"`
	// source file name, only file name, not full path
	FileName string `json:"file_name" example:"员工考勤管理制度-2023.pdf"`
	// page number in the source file
	PageNumber int `json:"page_number" example:"1"`
	// related content in the source file or in webpage
	Content string `json:"content" example:"旷工最小计算单位为0.5天，不足0.5天以0.5天计算，超过0.5天不满1天以1天计算，以此类推。"`
	// Title of the webpage
	Title string `json:"title,omitempty" example:"开始使用 Microsoft 帐户 – Microsoft"`
	// URL of the webpage
	URL string `json:"url,omitempty" example:"https://www.microsoft.com/zh-cn/welcome"`
}

func (reference Reference) String() string {
	bytes, err := json.Marshal(&reference)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func AddReferencesToArgs(args map[string]any, refs []Reference) map[string]any {
	if len(refs) == 0 {
		return args
	}
	old, exist := args["_references"]
	if exist {
		oldRefs := old.([]Reference)
		args["_references"] = append(oldRefs, refs...)
		return args
	}
	args["_references"] = refs
	return args
}
