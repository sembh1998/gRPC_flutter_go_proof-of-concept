// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: service.proto

package pb

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

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x73, 0x74, 0x61, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc4,
	0x04, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2b, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x0d, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x0c, 0x67, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x13, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x67, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x0e, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x07, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x19, 0x67, 0x65, 0x74, 0x45, 0x73, 0x74, 0x61, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x1f, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x14, 0x67, 0x65,
	0x74, 0x45, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x19, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x45, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x67, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6d, 0x62, 0x68, 0x31, 0x39, 0x39, 0x38, 0x2f, 0x67, 0x52,
	0x43, 0x50, 0x5f, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x2d, 0x6f, 0x66, 0x2d, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*StoreRequest)(nil),                   // 0: StoreRequest
	(*StoreByIdRequest)(nil),               // 1: StoreByIdRequest
	(*StoreByNameRequest)(nil),             // 2: StoreByNameRequest
	(*ProductRequest)(nil),                 // 3: ProductRequest
	(*ProductByIdRequest)(nil),             // 4: ProductByIdRequest
	(*MemberRequest)(nil),                  // 5: MemberRequest
	(*MemberByIdRequest)(nil),              // 6: MemberByIdRequest
	(*Member)(nil),                         // 7: Member
	(*EstablishmentsByStoreIdRequest)(nil), // 8: EstablishmentsByStoreIdRequest
	(*EstablishmentByIdRequest)(nil),       // 9: EstablishmentByIdRequest
	(*StoresResponse)(nil),                 // 10: StoresResponse
	(*StoreResponse)(nil),                  // 11: StoreResponse
	(*ProductsResponse)(nil),               // 12: ProductsResponse
	(*ProductResponse)(nil),                // 13: ProductResponse
	(*MembersResponse)(nil),                // 14: MembersResponse
	(*MemberResponse)(nil),                 // 15: MemberResponse
	(*EstablishmentsResponse)(nil),         // 16: EstablishmentsResponse
	(*EstablishmentResponse)(nil),          // 17: EstablishmentResponse
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: ModuleService.getStores:input_type -> StoreRequest
	1,  // 1: ModuleService.getStoreById:input_type -> StoreByIdRequest
	2,  // 2: ModuleService.getStoreByName:input_type -> StoreByNameRequest
	3,  // 3: ModuleService.getProducts:input_type -> ProductRequest
	4,  // 4: ModuleService.getProductById:input_type -> ProductByIdRequest
	5,  // 5: ModuleService.getMembers:input_type -> MemberRequest
	6,  // 6: ModuleService.getMemberById:input_type -> MemberByIdRequest
	7,  // 7: ModuleService.createMember:input_type -> Member
	8,  // 8: ModuleService.getEstablishmentsPerStore:input_type -> EstablishmentsByStoreIdRequest
	9,  // 9: ModuleService.getEstablishmentById:input_type -> EstablishmentByIdRequest
	10, // 10: ModuleService.getStores:output_type -> StoresResponse
	11, // 11: ModuleService.getStoreById:output_type -> StoreResponse
	11, // 12: ModuleService.getStoreByName:output_type -> StoreResponse
	12, // 13: ModuleService.getProducts:output_type -> ProductsResponse
	13, // 14: ModuleService.getProductById:output_type -> ProductResponse
	14, // 15: ModuleService.getMembers:output_type -> MembersResponse
	15, // 16: ModuleService.getMemberById:output_type -> MemberResponse
	15, // 17: ModuleService.createMember:output_type -> MemberResponse
	16, // 18: ModuleService.getEstablishmentsPerStore:output_type -> EstablishmentsResponse
	17, // 19: ModuleService.getEstablishmentById:output_type -> EstablishmentResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_store_proto_init()
	file_product_proto_init()
	file_member_proto_init()
	file_establishment_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
