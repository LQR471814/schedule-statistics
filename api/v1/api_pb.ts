// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file v1/api.proto (syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file v1/api.proto.
 */
export const file_v1_api: GenFile = /*@__PURE__*/
  fileDesc("Cgx2MS9hcGkucHJvdG8iiQEKBUV2ZW50EgwKBG5hbWUYASABKAkSDAoEdGFncxgCIAMoDRIpCgVzdGFydBgDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXASJwoDZW5kGAQgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcBIQCghkdXJhdGlvbhgFIAEoDSIPCg1FdmVudHNSZXF1ZXN0IjYKDkV2ZW50c1Jlc3BvbnNlEgwKBHRhZ3MYASADKAkSFgoGZXZlbnRzGAIgAygLMgYuRXZlbnQyPAoPQ2FsZW5kYXJTZXJ2aWNlEikKBkV2ZW50cxIOLkV2ZW50c1JlcXVlc3QaDy5FdmVudHNSZXNwb25zZWIGcHJvdG8z", [file_google_protobuf_timestamp]);

/**
 * @generated from message Event
 */
export type Event = Message<"Event"> & {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * each element is an index for a tag
   *
   * @generated from field: repeated uint32 tags = 2;
   */
  tags: number[];

  /**
   * @generated from field: google.protobuf.Timestamp start = 3;
   */
  start?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp end = 4;
   */
  end?: Timestamp;

  /**
   * duration is in minutes
   *
   * @generated from field: uint32 duration = 5;
   */
  duration: number;
};

/**
 * Describes the message Event.
 * Use `create(EventSchema)` to create a new message.
 */
export const EventSchema: GenMessage<Event> = /*@__PURE__*/
  messageDesc(file_v1_api, 0);

/**
 * Events
 *
 * @generated from message EventsRequest
 */
export type EventsRequest = Message<"EventsRequest"> & {
};

/**
 * Describes the message EventsRequest.
 * Use `create(EventsRequestSchema)` to create a new message.
 */
export const EventsRequestSchema: GenMessage<EventsRequest> = /*@__PURE__*/
  messageDesc(file_v1_api, 1);

/**
 * @generated from message EventsResponse
 */
export type EventsResponse = Message<"EventsResponse"> & {
  /**
   * @generated from field: repeated string tags = 1;
   */
  tags: string[];

  /**
   * @generated from field: repeated Event events = 2;
   */
  events: Event[];
};

/**
 * Describes the message EventsResponse.
 * Use `create(EventsResponseSchema)` to create a new message.
 */
export const EventsResponseSchema: GenMessage<EventsResponse> = /*@__PURE__*/
  messageDesc(file_v1_api, 2);

/**
 * @generated from service CalendarService
 */
export const CalendarService: GenService<{
  /**
   * @generated from rpc CalendarService.Events
   */
  events: {
    methodKind: "unary";
    input: typeof EventsRequestSchema;
    output: typeof EventsResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_v1_api, 0);

