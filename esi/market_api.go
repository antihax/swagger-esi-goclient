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

type MarketApiService service

/**
 * List market prices
 * Return a list of prices  ---  Alternate route: &#x60;/v1/markets/prices/&#x60;  Alternate route: &#x60;/legacy/markets/prices/&#x60;  Alternate route: &#x60;/dev/markets/prices/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param datasource(string) The server name you would like data from
 * @return []GetMarketsPrices200Ok
 */
func (a MarketApiService) GetMarketsPrices(datasource interface{}) ([]GetMarketsPrices200Ok, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/markets/prices/"

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
	var successPayload = new([]GetMarketsPrices200Ok)

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
 * List historical market statistics in a region
 * Return a list of historical market statistics for the specified type in a region  ---  Alternate route: &#x60;/v1/markets/{region_id}/history/&#x60;  Alternate route: &#x60;/legacy/markets/{region_id}/history/&#x60;  Alternate route: &#x60;/dev/markets/{region_id}/history/&#x60;   ---  This route is cached for up to 300 seconds
 *
 * @param regionId Return statistics in this region
 * @param typeId Return statistics for this type
 * @param datasource(string) The server name you would like data from
 * @return []GetMarketsRegionIdHistory200Ok
 */
func (a MarketApiService) GetMarketsRegionIdHistory(regionId int64, typeId int64, datasource interface{}) ([]GetMarketsRegionIdHistory200Ok, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/markets/{region_id}/history/"
	localVarPath = strings.Replace(localVarPath, "{"+"region_id"+"}", fmt.Sprintf("%v", regionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, err
	}
	localVarQueryParams.Add("type_id", a.client.parameterToString(typeId, ""))
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
	var successPayload = new([]GetMarketsRegionIdHistory200Ok)

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
 * List orders in a region
 * Return a list of orders in a region  ---  Alternate route: &#x60;/v1/markets/{region_id}/orders/&#x60;  Alternate route: &#x60;/legacy/markets/{region_id}/orders/&#x60;  Alternate route: &#x60;/dev/markets/{region_id}/orders/&#x60;   ---  This route is cached for up to 300 seconds
 *
 * @param regionId Return orders in this region
 * @param orderType Filter buy/sell orders, return all orders by default. If you query without type_id, we always return both buy and sell orders.
 * @param typeId(int64) Return orders only for this type
 * @param page(int32) Which page to query, only used for querying without type_id. Starting at 1
 * @param datasource(string) The server name you would like data from
 * @return []GetMarketsRegionIdOrders200Ok
 */
func (a MarketApiService) GetMarketsRegionIdOrders(regionId int64, orderType string, typeId interface{}, page interface{}, datasource interface{}) ([]GetMarketsRegionIdOrders200Ok, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/markets/{region_id}/orders/"
	localVarPath = strings.Replace(localVarPath, "{"+"region_id"+"}", fmt.Sprintf("%v", regionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(typeId, "int64", "typeId"); err != nil {
		return nil, err
	}
	if err := a.client.typeCheckParameter(page, "int32", "page"); err != nil {
		return nil, err
	}
	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, err
	}
	if typeId != nil {
		localVarQueryParams.Add("type_id", a.client.parameterToString(typeId, ""))
	}
	localVarQueryParams.Add("order_type", a.client.parameterToString(orderType, ""))
	if page != nil {
		localVarQueryParams.Add("page", a.client.parameterToString(page, ""))
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
	var successPayload = new([]GetMarketsRegionIdOrders200Ok)

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
