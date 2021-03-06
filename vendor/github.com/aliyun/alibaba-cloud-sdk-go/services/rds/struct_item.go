package rds

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

// Item is a nested struct in rds response
type Item struct {
	CrossBackupSetSize        int64                                      `json:"CrossBackupSetSize" xml:"CrossBackupSetSize"`
	Category                  string                                     `json:"Category" xml:"Category"`
	DBInstanceDescription     string                                     `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
	Engine                    string                                     `json:"Engine" xml:"Engine"`
	CrossLogBackupId          int                                        `json:"CrossLogBackupId" xml:"CrossLogBackupId"`
	CrossDownloadLink         string                                     `json:"CrossDownloadLink" xml:"CrossDownloadLink"`
	LogFileName               string                                     `json:"LogFileName" xml:"LogFileName"`
	ReportTime                string                                     `json:"ReportTime" xml:"ReportTime"`
	BackupSetStatus           int                                        `json:"BackupSetStatus" xml:"BackupSetStatus"`
	DBInstanceStatusDesc      string                                     `json:"DBInstanceStatusDesc" xml:"DBInstanceStatusDesc"`
	BackupEnabled             string                                     `json:"BackupEnabled" xml:"BackupEnabled"`
	RegionId                  string                                     `json:"RegionId" xml:"RegionId"`
	CrossBackupRegion         string                                     `json:"CrossBackupRegion" xml:"CrossBackupRegion"`
	Duration                  int                                        `json:"Duration" xml:"Duration"`
	CrossBackupId             int                                        `json:"CrossBackupId" xml:"CrossBackupId"`
	BackupStartTime           string                                     `json:"BackupStartTime" xml:"BackupStartTime"`
	LinkExpiredTime           string                                     `json:"LinkExpiredTime" xml:"LinkExpiredTime"`
	DBInstanceId              string                                     `json:"DBInstanceId" xml:"DBInstanceId"`
	BackupSetScale            int                                        `json:"BackupSetScale" xml:"BackupSetScale"`
	DBInstanceStorageType     string                                     `json:"DBInstanceStorageType" xml:"DBInstanceStorageType"`
	AutoRenew                 string                                     `json:"AutoRenew" xml:"AutoRenew"`
	CrossLogBackupSize        int64                                      `json:"CrossLogBackupSize" xml:"CrossLogBackupSize"`
	BackupEnabledTime         string                                     `json:"BackupEnabledTime" xml:"BackupEnabledTime"`
	CrossBackupSetFile        string                                     `json:"CrossBackupSetFile" xml:"CrossBackupSetFile"`
	CrossBackupDownloadLink   string                                     `json:"CrossBackupDownloadLink" xml:"CrossBackupDownloadLink"`
	LogEndTime                string                                     `json:"LogEndTime" xml:"LogEndTime"`
	EngineVersion             string                                     `json:"EngineVersion" xml:"EngineVersion"`
	CrossBackupSetLocation    string                                     `json:"CrossBackupSetLocation" xml:"CrossBackupSetLocation"`
	LogBackupEnabled          string                                     `json:"LogBackupEnabled" xml:"LogBackupEnabled"`
	LogBackupEnabledTime      string                                     `json:"LogBackupEnabledTime" xml:"LogBackupEnabledTime"`
	BackupEndTime             string                                     `json:"BackupEndTime" xml:"BackupEndTime"`
	BackupMethod              string                                     `json:"BackupMethod" xml:"BackupMethod"`
	RelService                string                                     `json:"RelService" xml:"RelService"`
	BackupType                string                                     `json:"BackupType" xml:"BackupType"`
	LogBeginTime              string                                     `json:"LogBeginTime" xml:"LogBeginTime"`
	Retention                 int                                        `json:"Retention" xml:"Retention"`
	DBInstanceStatus          string                                     `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
	RetentType                int                                        `json:"RetentType" xml:"RetentType"`
	CrossIntranetDownloadLink string                                     `json:"CrossIntranetDownloadLink" xml:"CrossIntranetDownloadLink"`
	InstanceId                int                                        `json:"InstanceId" xml:"InstanceId"`
	ConsistentTime            string                                     `json:"ConsistentTime" xml:"ConsistentTime"`
	CrossBackupType           string                                     `json:"CrossBackupType" xml:"CrossBackupType"`
	Status                    string                                     `json:"Status" xml:"Status"`
	RelServiceId              string                                     `json:"RelServiceId" xml:"RelServiceId"`
	LockMode                  string                                     `json:"LockMode" xml:"LockMode"`
	RestoreRegions            RestoreRegions                             `json:"RestoreRegions" xml:"RestoreRegions"`
	LatencyTopNItems          LatencyTopNItemsInDescribeSQLLogReportList `json:"LatencyTopNItems" xml:"LatencyTopNItems"`
	QPSTopNItems              QPSTopNItemsInDescribeSQLLogReportList     `json:"QPSTopNItems" xml:"QPSTopNItems"`
}
