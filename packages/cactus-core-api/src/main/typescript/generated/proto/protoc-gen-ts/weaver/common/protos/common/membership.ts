/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.19.1
 * source: common/membership.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as pb_1 from "google-protobuf";
export namespace common.membership {
    export class Membership extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            securityDomain?: string;
            members?: Map<string, Member>;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("securityDomain" in data && data.securityDomain != undefined) {
                    this.securityDomain = data.securityDomain;
                }
                if ("members" in data && data.members != undefined) {
                    this.members = data.members;
                }
            }
            if (!this.members)
                this.members = new Map();
        }
        get securityDomain() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set securityDomain(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get members() {
            return pb_1.Message.getField(this, 2) as any as Map<string, Member>;
        }
        set members(value: Map<string, Member>) {
            pb_1.Message.setField(this, 2, value as any);
        }
        static fromObject(data: {
            securityDomain?: string;
            members?: {
                [key: string]: ReturnType<typeof Member.prototype.toObject>;
            };
        }): Membership {
            const message = new Membership({});
            if (data.securityDomain != null) {
                message.securityDomain = data.securityDomain;
            }
            if (typeof data.members == "object") {
                message.members = new Map(Object.entries(data.members).map(([key, value]) => [key, Member.fromObject(value)]));
            }
            return message;
        }
        toObject() {
            const data: {
                securityDomain?: string;
                members?: {
                    [key: string]: ReturnType<typeof Member.prototype.toObject>;
                };
            } = {};
            if (this.securityDomain != null) {
                data.securityDomain = this.securityDomain;
            }
            if (this.members != null) {
                data.members = (Object.fromEntries)((Array.from)(this.members).map(([key, value]) => [key, value.toObject()]));
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.securityDomain.length)
                writer.writeString(1, this.securityDomain);
            for (const [key, value] of this.members) {
                writer.writeMessage(2, this.members, () => {
                    writer.writeString(1, key);
                    writer.writeMessage(2, value, () => value.serialize(writer));
                });
            }
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Membership {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new Membership();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.securityDomain = reader.readString();
                        break;
                    case 2:
                        reader.readMessage(message, () => pb_1.Map.deserializeBinary(message.members as any, reader, reader.readString, () => {
                            let value;
                            reader.readMessage(message, () => value = Member.deserialize(reader));
                            return value;
                        }));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): Membership {
            return Membership.deserialize(bytes);
        }
    }
    export class Member extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            value?: string;
            type?: string;
            chain?: string[];
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [3], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("value" in data && data.value != undefined) {
                    this.value = data.value;
                }
                if ("type" in data && data.type != undefined) {
                    this.type = data.type;
                }
                if ("chain" in data && data.chain != undefined) {
                    this.chain = data.chain;
                }
            }
        }
        get value() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set value(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get type() {
            return pb_1.Message.getFieldWithDefault(this, 2, "") as string;
        }
        set type(value: string) {
            pb_1.Message.setField(this, 2, value);
        }
        get chain() {
            return pb_1.Message.getFieldWithDefault(this, 3, []) as string[];
        }
        set chain(value: string[]) {
            pb_1.Message.setField(this, 3, value);
        }
        static fromObject(data: {
            value?: string;
            type?: string;
            chain?: string[];
        }): Member {
            const message = new Member({});
            if (data.value != null) {
                message.value = data.value;
            }
            if (data.type != null) {
                message.type = data.type;
            }
            if (data.chain != null) {
                message.chain = data.chain;
            }
            return message;
        }
        toObject() {
            const data: {
                value?: string;
                type?: string;
                chain?: string[];
            } = {};
            if (this.value != null) {
                data.value = this.value;
            }
            if (this.type != null) {
                data.type = this.type;
            }
            if (this.chain != null) {
                data.chain = this.chain;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.value.length)
                writer.writeString(1, this.value);
            if (this.type.length)
                writer.writeString(2, this.type);
            if (this.chain.length)
                writer.writeRepeatedString(3, this.chain);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Member {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new Member();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.value = reader.readString();
                        break;
                    case 2:
                        message.type = reader.readString();
                        break;
                    case 3:
                        pb_1.Message.addToRepeatedField(message, 3, reader.readString());
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): Member {
            return Member.deserialize(bytes);
        }
    }
}
