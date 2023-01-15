///
//  Generated code. Do not modify.
//  source: store.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'store.pb.dart' as $0;
export 'store.pb.dart';

class StoreServiceClient extends $grpc.Client {
  static final _$getStores =
      $grpc.ClientMethod<$0.StoreRequest, $0.StoreReponse>(
          '/StoreService/getStores',
          ($0.StoreRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.StoreReponse.fromBuffer(value));
  static final _$getStoreById =
      $grpc.ClientMethod<$0.StoreByIdRequest, $0.Store>(
          '/StoreService/getStoreById',
          ($0.StoreByIdRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Store.fromBuffer(value));
  static final _$getStoreByName =
      $grpc.ClientMethod<$0.StoreByNameRequest, $0.Store>(
          '/StoreService/getStoreByName',
          ($0.StoreByNameRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Store.fromBuffer(value));

  StoreServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.StoreReponse> getStores($0.StoreRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getStores, request, options: options);
  }

  $grpc.ResponseFuture<$0.Store> getStoreById($0.StoreByIdRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getStoreById, request, options: options);
  }

  $grpc.ResponseFuture<$0.Store> getStoreByName($0.StoreByNameRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getStoreByName, request, options: options);
  }
}

abstract class StoreServiceBase extends $grpc.Service {
  $core.String get $name => 'StoreService';

  StoreServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.StoreRequest, $0.StoreReponse>(
        'getStores',
        getStores_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.StoreRequest.fromBuffer(value),
        ($0.StoreReponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.StoreByIdRequest, $0.Store>(
        'getStoreById',
        getStoreById_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.StoreByIdRequest.fromBuffer(value),
        ($0.Store value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.StoreByNameRequest, $0.Store>(
        'getStoreByName',
        getStoreByName_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.StoreByNameRequest.fromBuffer(value),
        ($0.Store value) => value.writeToBuffer()));
  }

  $async.Future<$0.StoreReponse> getStores_Pre(
      $grpc.ServiceCall call, $async.Future<$0.StoreRequest> request) async {
    return getStores(call, await request);
  }

  $async.Future<$0.Store> getStoreById_Pre($grpc.ServiceCall call,
      $async.Future<$0.StoreByIdRequest> request) async {
    return getStoreById(call, await request);
  }

  $async.Future<$0.Store> getStoreByName_Pre($grpc.ServiceCall call,
      $async.Future<$0.StoreByNameRequest> request) async {
    return getStoreByName(call, await request);
  }

  $async.Future<$0.StoreReponse> getStores(
      $grpc.ServiceCall call, $0.StoreRequest request);
  $async.Future<$0.Store> getStoreById(
      $grpc.ServiceCall call, $0.StoreByIdRequest request);
  $async.Future<$0.Store> getStoreByName(
      $grpc.ServiceCall call, $0.StoreByNameRequest request);
}
