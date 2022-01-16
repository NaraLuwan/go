// Copyright 2020 miraclerliu@tencent.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	ebpb "git.code.oa.com/tencentcloud-serverless/construction/mds_v2/eb"
	mdssdkschema "git.code.oa.com/tencentcloud-serverless/mds-sdk-go/modelschema"
	mdssdkstore "git.code.oa.com/tencentcloud-serverless/mds-sdk-go/store"
	apis "git.code.oa.com/tencentcloud-serverless/metadata-server/pkg/schema/core/pkg/apis"
	json "github.com/json-iterator/go"
)

func main() {
	parseSchema, _ := genParseSchema()
	fmt.Printf("%+v", parseSchema)
}

func genParseSchema() (map[string]string, error) {
	parseSchemas := map[string]*apis.Schema{
		mdssdkstore.ResourceEventBus:       apis.GenSchema(&ebpb.EventBusResponse{}),
		mdssdkstore.ResourceEventBusRule:   apis.GenSchema(&ebpb.EventBusRule{}),
		mdssdkstore.ResourceEventBusTarget: apis.GenSchema(&ebpb.EventBusTarget{}),
	}
	errList := make([]error, 0)
	resMap := make(map[string]string, len(parseSchemas))
	for resource, schema := range parseSchemas {
		s, sErr := schema.VisitRootSchema()
		if sErr != nil && sErr.HasError() {
			errList = append(errList, fmt.Errorf("failed to prepare schema,err:%v", sErr))
			continue
		}
		sd, err := json.Marshal(s)
		if err != nil {
			errList = append(errList, fmt.Errorf("failed to marshal schema,err:%v", err))
			continue
		}
		resMap[resource] = string(sd)
	}
	if len(errList) != 0 {
		return nil, fmt.Errorf("+%v", errList)
	}
	return resMap, nil
}

func genValidateSchema() (map[string]string, error) {
	validateSchemas := map[string]*apis.Schema{
		mdssdkstore.ResourceEventBus:       apis.GenSchema(&mdssdkschema.EventBusSchema{}),
		mdssdkstore.ResourceEventBusRule:   apis.GenSchema(&mdssdkschema.EventBusRuleSchema{}),
		mdssdkstore.ResourceEventBusTarget: apis.GenSchema(&mdssdkschema.TargetSchema{}),
	}
	errList := make([]error, 0)
	resMap := make(map[string]string, len(validateSchemas))
	for resource, schema := range validateSchemas {
		s, sErr := schema.VisitRootSchema()
		if sErr != nil && sErr.HasError() {
			errList = append(errList, fmt.Errorf("failed to prepare schema,err:%v", sErr))
			continue
		}
		sd, err := json.Marshal(s)
		if err != nil {
			errList = append(errList, fmt.Errorf("failed to marshal schema,err:%v", err))
			continue
		}
		resMap[resource] = string(sd)
	}
	if len(errList) != 0 {
		return nil, fmt.Errorf("+%v", errList)
	}
	return resMap, nil
}
