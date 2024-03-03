// Copyright (C) 2024  mieru authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: rpc.proto

package appctlpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x70,
	0x63, 0x74, 0x6c, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x63, 0x66, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xed, 0x03, 0x0a,
	0x16, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x41, 0x70, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x24, 0x0a, 0x04, 0x45, 0x78, 0x69,
	0x74, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x2c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x0d, 0x2e,
	0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x61,
	0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x34, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13,
	0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x44, 0x75, 0x6d, 0x70, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x39, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x43, 0x50, 0x55, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x70,
	0x63, 0x74, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x61, 0x76, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x2e, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x70, 0x43, 0x50, 0x55, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x70, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x0d, 0x2e,
	0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x32, 0xe2, 0x04, 0x0a,
	0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x41, 0x70, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x25, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x24, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x06, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x24,
	0x0a, 0x04, 0x45, 0x78, 0x69, 0x74, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63,
	0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74,
	0x6c, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x39, 0x0a, 0x0f,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x50, 0x55, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x17, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x70, 0x43,
	0x50, 0x55, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63,
	0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x65,
	0x61, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x63,
	0x74, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x3e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c,
	0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x32, 0x80, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x09, 0x53,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74,
	0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x14,
	0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x66, 0x65, 0x69, 0x6e, 0x2f, 0x6d, 0x69, 0x65, 0x72, 0x75, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x74,
	0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_rpc_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: appctl.Empty
	(*ProfileSavePath)(nil),  // 1: appctl.ProfileSavePath
	(*ServerConfig)(nil),     // 2: appctl.ServerConfig
	(*AppStatusMsg)(nil),     // 3: appctl.AppStatusMsg
	(*Metrics)(nil),          // 4: appctl.Metrics
	(*SessionInfo)(nil),      // 5: appctl.SessionInfo
	(*ThreadDump)(nil),       // 6: appctl.ThreadDump
	(*MemoryStatistics)(nil), // 7: appctl.MemoryStatistics
}
var file_rpc_proto_depIdxs = []int32{
	0,  // 0: appctl.ClientLifecycleService.GetStatus:input_type -> appctl.Empty
	0,  // 1: appctl.ClientLifecycleService.Exit:input_type -> appctl.Empty
	0,  // 2: appctl.ClientLifecycleService.GetMetrics:input_type -> appctl.Empty
	0,  // 3: appctl.ClientLifecycleService.GetSessionInfo:input_type -> appctl.Empty
	0,  // 4: appctl.ClientLifecycleService.GetThreadDump:input_type -> appctl.Empty
	1,  // 5: appctl.ClientLifecycleService.StartCPUProfile:input_type -> appctl.ProfileSavePath
	0,  // 6: appctl.ClientLifecycleService.StopCPUProfile:input_type -> appctl.Empty
	1,  // 7: appctl.ClientLifecycleService.GetHeapProfile:input_type -> appctl.ProfileSavePath
	0,  // 8: appctl.ClientLifecycleService.GetMemoryStatistics:input_type -> appctl.Empty
	0,  // 9: appctl.ServerLifecycleService.GetStatus:input_type -> appctl.Empty
	0,  // 10: appctl.ServerLifecycleService.Start:input_type -> appctl.Empty
	0,  // 11: appctl.ServerLifecycleService.Stop:input_type -> appctl.Empty
	0,  // 12: appctl.ServerLifecycleService.Reload:input_type -> appctl.Empty
	0,  // 13: appctl.ServerLifecycleService.Exit:input_type -> appctl.Empty
	0,  // 14: appctl.ServerLifecycleService.GetMetrics:input_type -> appctl.Empty
	0,  // 15: appctl.ServerLifecycleService.GetSessionInfo:input_type -> appctl.Empty
	0,  // 16: appctl.ServerLifecycleService.GetThreadDump:input_type -> appctl.Empty
	1,  // 17: appctl.ServerLifecycleService.StartCPUProfile:input_type -> appctl.ProfileSavePath
	0,  // 18: appctl.ServerLifecycleService.StopCPUProfile:input_type -> appctl.Empty
	1,  // 19: appctl.ServerLifecycleService.GetHeapProfile:input_type -> appctl.ProfileSavePath
	0,  // 20: appctl.ServerLifecycleService.GetMemoryStatistics:input_type -> appctl.Empty
	0,  // 21: appctl.ServerConfigService.GetConfig:input_type -> appctl.Empty
	2,  // 22: appctl.ServerConfigService.SetConfig:input_type -> appctl.ServerConfig
	3,  // 23: appctl.ClientLifecycleService.GetStatus:output_type -> appctl.AppStatusMsg
	0,  // 24: appctl.ClientLifecycleService.Exit:output_type -> appctl.Empty
	4,  // 25: appctl.ClientLifecycleService.GetMetrics:output_type -> appctl.Metrics
	5,  // 26: appctl.ClientLifecycleService.GetSessionInfo:output_type -> appctl.SessionInfo
	6,  // 27: appctl.ClientLifecycleService.GetThreadDump:output_type -> appctl.ThreadDump
	0,  // 28: appctl.ClientLifecycleService.StartCPUProfile:output_type -> appctl.Empty
	0,  // 29: appctl.ClientLifecycleService.StopCPUProfile:output_type -> appctl.Empty
	0,  // 30: appctl.ClientLifecycleService.GetHeapProfile:output_type -> appctl.Empty
	7,  // 31: appctl.ClientLifecycleService.GetMemoryStatistics:output_type -> appctl.MemoryStatistics
	3,  // 32: appctl.ServerLifecycleService.GetStatus:output_type -> appctl.AppStatusMsg
	0,  // 33: appctl.ServerLifecycleService.Start:output_type -> appctl.Empty
	0,  // 34: appctl.ServerLifecycleService.Stop:output_type -> appctl.Empty
	0,  // 35: appctl.ServerLifecycleService.Reload:output_type -> appctl.Empty
	0,  // 36: appctl.ServerLifecycleService.Exit:output_type -> appctl.Empty
	4,  // 37: appctl.ServerLifecycleService.GetMetrics:output_type -> appctl.Metrics
	5,  // 38: appctl.ServerLifecycleService.GetSessionInfo:output_type -> appctl.SessionInfo
	6,  // 39: appctl.ServerLifecycleService.GetThreadDump:output_type -> appctl.ThreadDump
	0,  // 40: appctl.ServerLifecycleService.StartCPUProfile:output_type -> appctl.Empty
	0,  // 41: appctl.ServerLifecycleService.StopCPUProfile:output_type -> appctl.Empty
	0,  // 42: appctl.ServerLifecycleService.GetHeapProfile:output_type -> appctl.Empty
	7,  // 43: appctl.ServerLifecycleService.GetMemoryStatistics:output_type -> appctl.MemoryStatistics
	2,  // 44: appctl.ServerConfigService.GetConfig:output_type -> appctl.ServerConfig
	2,  // 45: appctl.ServerConfigService.SetConfig:output_type -> appctl.ServerConfig
	23, // [23:46] is the sub-list for method output_type
	0,  // [0:23] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	file_base_proto_init()
	file_misc_proto_init()
	file_servercfg_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
