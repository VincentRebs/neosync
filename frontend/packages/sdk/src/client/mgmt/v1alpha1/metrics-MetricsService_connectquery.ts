// @generated by protoc-gen-connect-query v1.4.2 with parameter "target=ts,import_extension=.js"
// @generated from file mgmt/v1alpha1/metrics.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { GetDailyMetricCountRequest, GetDailyMetricCountResponse, GetMetricCountRequest, GetMetricCountResponse } from "./metrics_pb.js";

/**
 * Retrieve a timed range of records
 *
 * @generated from rpc mgmt.v1alpha1.MetricsService.GetDailyMetricCount
 */
export const getDailyMetricCount = {
  localName: "getDailyMetricCount",
  name: "GetDailyMetricCount",
  kind: MethodKind.Unary,
  I: GetDailyMetricCountRequest,
  O: GetDailyMetricCountResponse,
  service: {
    typeName: "mgmt.v1alpha1.MetricsService"
  }
} as const;

/**
 * For the given metric and time range, returns the total count found
 *
 * @generated from rpc mgmt.v1alpha1.MetricsService.GetMetricCount
 */
export const getMetricCount = {
  localName: "getMetricCount",
  name: "GetMetricCount",
  kind: MethodKind.Unary,
  I: GetMetricCountRequest,
  O: GetMetricCountResponse,
  service: {
    typeName: "mgmt.v1alpha1.MetricsService"
  }
} as const;
