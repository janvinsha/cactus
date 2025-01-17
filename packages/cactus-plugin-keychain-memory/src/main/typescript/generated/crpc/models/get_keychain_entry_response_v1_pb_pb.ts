//
//Hyperledger Cactus Plugin - Keychain Memory 
//
//Contains/describes the Hyperledger Cacti Keychain Memory plugin.
//
//The version of the OpenAPI document: 2.1.0
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.8.0 with parameter "target=ts"
// @generated from file models/get_keychain_entry_response_v1_pb.proto (package org.hyperledger.cacti.plugin.keychain.memory, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message org.hyperledger.cacti.plugin.keychain.memory.GetKeychainEntryResponseV1PB
 */
export class GetKeychainEntryResponseV1PB extends Message<GetKeychainEntryResponseV1PB> {
  /**
   * The key that was used to retrieve the value from the keychain.
   *
   * @generated from field: string key = 106079;
   */
  key = "";

  /**
   * The value associated with the requested key on the keychain.
   *
   * @generated from field: string value = 111972721;
   */
  value = "";

  constructor(data?: PartialMessage<GetKeychainEntryResponseV1PB>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "org.hyperledger.cacti.plugin.keychain.memory.GetKeychainEntryResponseV1PB";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 106079, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 111972721, name: "value", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetKeychainEntryResponseV1PB {
    return new GetKeychainEntryResponseV1PB().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetKeychainEntryResponseV1PB {
    return new GetKeychainEntryResponseV1PB().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetKeychainEntryResponseV1PB {
    return new GetKeychainEntryResponseV1PB().fromJsonString(jsonString, options);
  }

  static equals(a: GetKeychainEntryResponseV1PB | PlainMessage<GetKeychainEntryResponseV1PB> | undefined, b: GetKeychainEntryResponseV1PB | PlainMessage<GetKeychainEntryResponseV1PB> | undefined): boolean {
    return proto3.util.equals(GetKeychainEntryResponseV1PB, a, b);
  }
}

