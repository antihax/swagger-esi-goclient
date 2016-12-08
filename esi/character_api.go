/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.1.dev2
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

type CharacterApiService service

/**
 * Get character&#39;s public information
 * Public information about a character  ---  Alternate route: &#x60;/v3/characters/{character_id}/&#x60;  Alternate route: &#x60;/legacy/characters/{character_id}/&#x60;  Alternate route: &#x60;/dev/characters/{character_id}/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param characterId An EVE character ID
 * @param datasource(string) The server name you would like data from
 * @return *GetCharactersCharacterIdOk
 */
func (a CharacterApiService) GetCharactersCharacterId(characterId int32, datasource interface{}) (*GetCharactersCharacterIdOk, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/characters/{character_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
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
	var successPayload = new(GetCharactersCharacterIdOk)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, err
	}

	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return nil, err
	}

	return successPayload, err
}

/**
 * Get corporation history
 * Get a list of all the corporations a character has been a member of  ---  Alternate route: &#x60;/v1/characters/{character_id}/corporationhistory/&#x60;  Alternate route: &#x60;/legacy/characters/{character_id}/corporationhistory/&#x60;  Alternate route: &#x60;/dev/characters/{character_id}/corporationhistory/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param characterId An EVE character ID
 * @param datasource(string) The server name you would like data from
 * @return []GetCharactersCharacterIdCorporationhistory200Ok
 */
func (a CharacterApiService) GetCharactersCharacterIdCorporationhistory(characterId int32, datasource interface{}) ([]GetCharactersCharacterIdCorporationhistory200Ok, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/characters/{character_id}/corporationhistory/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
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
	var successPayload = new([]GetCharactersCharacterIdCorporationhistory200Ok)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return *successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return *successPayload, err
	}

	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return nil, err
	}

	return *successPayload, err
}

/**
 * Get character portraits
 * Get portrait urls for a character  ---  Alternate route: &#x60;/v2/characters/{character_id}/portrait/&#x60;  Alternate route: &#x60;/dev/characters/{character_id}/portrait/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param characterId An EVE character ID
 * @param datasource(string) The server name you would like data from
 * @return *GetCharactersCharacterIdPortraitOk
 */
func (a CharacterApiService) GetCharactersCharacterIdPortrait(characterId int32, datasource interface{}) (*GetCharactersCharacterIdPortraitOk, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/characters/{character_id}/portrait/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
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
	var successPayload = new(GetCharactersCharacterIdPortraitOk)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, err
	}

	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return nil, err
	}

	return successPayload, err
}

/**
 * Get character names
 * Resolve a set of character IDs to character names  ---  Alternate route: &#x60;/v1/characters/names/&#x60;  Alternate route: &#x60;/legacy/characters/names/&#x60;  Alternate route: &#x60;/dev/characters/names/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param characterIds A comma separated list of character IDs
 * @param datasource(string) The server name you would like data from
 * @return []GetCharactersNames200Ok
 */
func (a CharacterApiService) GetCharactersNames(characterIds []int64, datasource interface{}) ([]GetCharactersNames200Ok, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/characters/names/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, err
	}
	localVarQueryParams.Add("character_ids", a.client.parameterToString(characterIds, "csv"))
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
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
	var successPayload = new([]GetCharactersNames200Ok)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return *successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return *successPayload, err
	}

	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return nil, err
	}

	return *successPayload, err
}

/**
 * Calculate a CSPA charge cost
 * Takes a source character ID in the url and a set of target character ID&#39;s in the body, returns a CSPA charge cost  ---  Alternate route: &#x60;/v3/characters/{character_id}/cspa/&#x60;  Alternate route: &#x60;/legacy/characters/{character_id}/cspa/&#x60;  Alternate route: &#x60;/dev/characters/{character_id}/cspa/&#x60;
 *
 * @param characterId An EVE character ID
 * @param characters The target characters to calculate the charge for
 * @param datasource(string) The server name you would like data from
 * @return *PostCharactersCharacterIdCspaCreated
 */
func (a CharacterApiService) PostCharactersCharacterIdCspa(ts TokenSource, characterId int32, characters PostCharactersCharacterIdCspaCharacters, datasource interface{}) (*PostCharactersCharacterIdCspaCreated, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/characters/{character_id}/cspa/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
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
	// body params
	localVarPostBody = &characters
	var successPayload = new(PostCharactersCharacterIdCspaCreated)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return successPayload, err
	}

	if ts != nil {
		if t, err := ts.Token(); err != nil {
			return successPayload, err
		} else if t != nil {
			t.SetAuthHeader(r)
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, err
	}

	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return nil, err
	}

	return successPayload, err
}
