///
//  Generated code. Do not modify.
//  source: store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use storeRequestDescriptor instead')
const StoreRequest$json = const {
  '1': 'StoreRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

/// Descriptor for `StoreRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List storeRequestDescriptor = $convert.base64Decode('CgxTdG9yZVJlcXVlc3QSDgoCaWQYASABKAlSAmlk');
@$core.Deprecated('Use storeByIdRequestDescriptor instead')
const StoreByIdRequest$json = const {
  '1': 'StoreByIdRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

/// Descriptor for `StoreByIdRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List storeByIdRequestDescriptor = $convert.base64Decode('ChBTdG9yZUJ5SWRSZXF1ZXN0Eg4KAmlkGAEgASgJUgJpZA==');
@$core.Deprecated('Use storeByNameRequestDescriptor instead')
const StoreByNameRequest$json = const {
  '1': 'StoreByNameRequest',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `StoreByNameRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List storeByNameRequestDescriptor = $convert.base64Decode('ChJTdG9yZUJ5TmFtZVJlcXVlc3QSEgoEbmFtZRgBIAEoCVIEbmFtZQ==');
@$core.Deprecated('Use storeReponseDescriptor instead')
const StoreReponse$json = const {
  '1': 'StoreReponse',
  '2': const [
    const {'1': 'Stores', '3': 1, '4': 3, '5': 11, '6': '.Store', '10': 'Stores'},
  ],
};

/// Descriptor for `StoreReponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List storeReponseDescriptor = $convert.base64Decode('CgxTdG9yZVJlcG9uc2USHgoGU3RvcmVzGAEgAygLMgYuU3RvcmVSBlN0b3Jlcw==');
@$core.Deprecated('Use storeDescriptor instead')
const Store$json = const {
  '1': 'Store',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `Store`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List storeDescriptor = $convert.base64Decode('CgVTdG9yZRIOCgJpZBgBIAEoCVICaWQSEgoEbmFtZRgCIAEoCVIEbmFtZQ==');
