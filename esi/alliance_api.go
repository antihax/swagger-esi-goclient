/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.2.7.dev1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type AllianceApiService service


/**
 * List all alliances
 * List all active player alliances  ---  Alternate route: &#x60;/v1/alliances/&#x60;  Alternate route: &#x60;/legacy/alliances/&#x60;  Alternate route: &#x60;/dev/alliances/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param datasource(nil) The server name you would like data from 
 * @return []int32
 */
func (a AllianceApiService) GetAlliances(datasource interface{}) ([]int32, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.client.Config.BasePath + "/alliances/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.client.Config.DefaultHeader {
		localVarHeaderParams[key] = a.client.Config.DefaultHeader[key]
	}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil,  nil, err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.client.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]int32)
	localVarHttpResponse, err := a.client.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetAlliances", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * Get alliance information
 * Public information about an alliance  ---  Alternate route: &#x60;/v2/alliances/{alliance_id}/&#x60;  Alternate route: &#x60;/dev/alliances/{alliance_id}/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param allianceId An Eve alliance ID 
 * @param datasource(nil) The server name you would like data from 
 * @return *GetAlliancesAllianceIdOk
 */
func (a AllianceApiService) GetAlliancesAllianceId(allianceId int32, datasource interface{}) (*GetAlliancesAllianceIdOk, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.client.Config.BasePath + "/alliances/{alliance_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"alliance_id"+"}", fmt.Sprintf("%v", allianceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.client.Config.DefaultHeader {
		localVarHeaderParams[key] = a.client.Config.DefaultHeader[key]
	}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil,  nil, err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.client.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GetAlliancesAllianceIdOk)
	localVarHttpResponse, err := a.client.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetAlliancesAllianceId", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * List alliance&#39;s corporations
 * List all current member corporations of an alliance  ---  Alternate route: &#x60;/v1/alliances/{alliance_id}/corporations/&#x60;  Alternate route: &#x60;/legacy/alliances/{alliance_id}/corporations/&#x60;  Alternate route: &#x60;/dev/alliances/{alliance_id}/corporations/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param allianceId An EVE alliance ID 
 * @param datasource(nil) The server name you would like data from 
 * @return []int32
 */
func (a AllianceApiService) GetAlliancesAllianceIdCorporations(allianceId int32, datasource interface{}) ([]int32, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.client.Config.BasePath + "/alliances/{alliance_id}/corporations/"
	localVarPath = strings.Replace(localVarPath, "{"+"alliance_id"+"}", fmt.Sprintf("%v", allianceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.client.Config.DefaultHeader {
		localVarHeaderParams[key] = a.client.Config.DefaultHeader[key]
	}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil,  nil, err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.client.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]int32)
	localVarHttpResponse, err := a.client.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetAlliancesAllianceIdCorporations", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * Get alliance icon
 * Get the icon urls for a alliance  ---  Alternate route: &#x60;/v1/alliances/{alliance_id}/icons/&#x60;  Alternate route: &#x60;/legacy/alliances/{alliance_id}/icons/&#x60;  Alternate route: &#x60;/dev/alliances/{alliance_id}/icons/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param allianceId An EVE alliance ID 
 * @param datasource(nil) The server name you would like data from 
 * @return *GetAlliancesAllianceIdIconsOk
 */
func (a AllianceApiService) GetAlliancesAllianceIdIcons(allianceId int32, datasource interface{}) (*GetAlliancesAllianceIdIconsOk, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.client.Config.BasePath + "/alliances/{alliance_id}/icons/"
	localVarPath = strings.Replace(localVarPath, "{"+"alliance_id"+"}", fmt.Sprintf("%v", allianceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.client.Config.DefaultHeader {
		localVarHeaderParams[key] = a.client.Config.DefaultHeader[key]
	}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil,  nil, err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.client.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GetAlliancesAllianceIdIconsOk)
	localVarHttpResponse, err := a.client.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetAlliancesAllianceIdIcons", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Get alliance names
 * Resolve a set of alliance IDs to alliance names  ---  Alternate route: &#x60;/v1/alliances/names/&#x60;  Alternate route: &#x60;/legacy/alliances/names/&#x60;  Alternate route: &#x60;/dev/alliances/names/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param allianceIds A comma separated list of alliance IDs 
 * @param datasource(nil) The server name you would like data from 
 * @return []GetAlliancesNames200Ok
 */
func (a AllianceApiService) GetAlliancesNames(allianceIds []int64, datasource interface{}) ([]GetAlliancesNames200Ok, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.client.Config.BasePath + "/alliances/names/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.client.Config.DefaultHeader {
		localVarHeaderParams[key] = a.client.Config.DefaultHeader[key]
	}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil,  nil, err
	}
		localVarQueryParams.Add("alliance_ids", a.client.parameterToString(allianceIds, "csv"))
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := a.client.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]GetAlliancesNames200Ok)
	localVarHttpResponse, err := a.client.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetAlliancesNames", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

