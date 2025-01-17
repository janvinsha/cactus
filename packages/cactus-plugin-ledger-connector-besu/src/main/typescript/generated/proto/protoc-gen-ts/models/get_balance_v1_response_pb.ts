/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.19.1
 * source: models/get_balance_v1_response_pb.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../google/protobuf/any";
import * as pb_1 from "google-protobuf";
export namespace org.hyperledger.cacti.plugin.ledger.connector.besu {
    export class GetBalanceV1ResponsePB extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            balance?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("balance" in data && data.balance != undefined) {
                    this.balance = data.balance;
                }
            }
        }
        get balance() {
            return pb_1.Message.getFieldWithDefault(this, 339185956, "") as string;
        }
        set balance(value: string) {
            pb_1.Message.setField(this, 339185956, value);
        }
        static fromObject(data: {
            balance?: string;
        }): GetBalanceV1ResponsePB {
            const message = new GetBalanceV1ResponsePB({});
            if (data.balance != null) {
                message.balance = data.balance;
            }
            return message;
        }
        toObject() {
            const data: {
                balance?: string;
            } = {};
            if (this.balance != null) {
                data.balance = this.balance;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.balance.length)
                writer.writeString(339185956, this.balance);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetBalanceV1ResponsePB {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetBalanceV1ResponsePB();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 339185956:
                        message.balance = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetBalanceV1ResponsePB {
            return GetBalanceV1ResponsePB.deserialize(bytes);
        }
    }
}
