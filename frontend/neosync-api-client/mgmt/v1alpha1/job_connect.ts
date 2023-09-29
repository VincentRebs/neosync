// @generated by protoc-gen-connect-es v1.0.0 with parameter "target=ts,import_extension=none"
// @generated from file mgmt/v1alpha1/job.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CancelJobRunRequest, CancelJobRunResponse, CreateJobDestinationConnectionsRequest, CreateJobDestinationConnectionsResponse, CreateJobRequest, CreateJobResponse, CreateJobRunRequest, CreateJobRunResponse, DeleteJobDestinationConnectionRequest, DeleteJobDestinationConnectionResponse, DeleteJobRequest, DeleteJobResponse, DeleteJobRunRequest, DeleteJobRunResponse, GetJobRequest, GetJobResponse, GetJobRunEventsRequest, GetJobRunEventsResponse, GetJobRunRequest, GetJobRunResponse, GetJobRunsRequest, GetJobRunsResponse, GetJobsRequest, GetJobsResponse, GetTransformersRequest, GetTransformersResponse, IsJobNameAvailableRequest, IsJobNameAvailableResponse, PauseJobRequest, PauseJobResponse, UpdateJobDestinationConnectionRequest, UpdateJobDestinationConnectionResponse, UpdateJobScheduleRequest, UpdateJobScheduleResponse, UpdateJobSourceConnectionRequest, UpdateJobSourceConnectionResponse } from "./job_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service mgmt.v1alpha1.JobService
 */
export const JobService = {
  typeName: "mgmt.v1alpha1.JobService",
  methods: {
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.GetJobs
     */
    getJobs: {
      name: "GetJobs",
      I: GetJobsRequest,
      O: GetJobsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.GetJob
     */
    getJob: {
      name: "GetJob",
      I: GetJobRequest,
      O: GetJobResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.CreateJob
     */
    createJob: {
      name: "CreateJob",
      I: CreateJobRequest,
      O: CreateJobResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.DeleteJob
     */
    deleteJob: {
      name: "DeleteJob",
      I: DeleteJobRequest,
      O: DeleteJobResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.IsJobNameAvailable
     */
    isJobNameAvailable: {
      name: "IsJobNameAvailable",
      I: IsJobNameAvailableRequest,
      O: IsJobNameAvailableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.UpdateJobSchedule
     */
    updateJobSchedule: {
      name: "UpdateJobSchedule",
      I: UpdateJobScheduleRequest,
      O: UpdateJobScheduleResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.UpdateJobSourceConnection
     */
    updateJobSourceConnection: {
      name: "UpdateJobSourceConnection",
      I: UpdateJobSourceConnectionRequest,
      O: UpdateJobSourceConnectionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.UpdateJobDestinationConnection
     */
    updateJobDestinationConnection: {
      name: "UpdateJobDestinationConnection",
      I: UpdateJobDestinationConnectionRequest,
      O: UpdateJobDestinationConnectionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.DeleteJobDestinationConnection
     */
    deleteJobDestinationConnection: {
      name: "DeleteJobDestinationConnection",
      I: DeleteJobDestinationConnectionRequest,
      O: DeleteJobDestinationConnectionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.CreateJobDestinationConnections
     */
    createJobDestinationConnections: {
      name: "CreateJobDestinationConnections",
      I: CreateJobDestinationConnectionsRequest,
      O: CreateJobDestinationConnectionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.PauseJob
     */
    pauseJob: {
      name: "PauseJob",
      I: PauseJobRequest,
      O: PauseJobResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.GetJobRuns
     */
    getJobRuns: {
      name: "GetJobRuns",
      I: GetJobRunsRequest,
      O: GetJobRunsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.GetJobRunEvents
     */
    getJobRunEvents: {
      name: "GetJobRunEvents",
      I: GetJobRunEventsRequest,
      O: GetJobRunEventsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.GetJobRun
     */
    getJobRun: {
      name: "GetJobRun",
      I: GetJobRunRequest,
      O: GetJobRunResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.DeleteJobRun
     */
    deleteJobRun: {
      name: "DeleteJobRun",
      I: DeleteJobRunRequest,
      O: DeleteJobRunResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.CreateJobRun
     */
    createJobRun: {
      name: "CreateJobRun",
      I: CreateJobRunRequest,
      O: CreateJobRunResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.CancelJobRun
     */
    cancelJobRun: {
      name: "CancelJobRun",
      I: CancelJobRunRequest,
      O: CancelJobRunResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc mgmt.v1alpha1.JobService.GetTransformers
     */
    getTransformers: {
      name: "GetTransformers",
      I: GetTransformersRequest,
      O: GetTransformersResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

