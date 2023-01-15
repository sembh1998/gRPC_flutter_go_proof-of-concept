// ignore_for_file: public_member_api_docs

import 'package:stores_api/stores_api.dart';

/// {@template api_client}
/// An http API client
/// {@endtemplate}
class StoresApiClient {
  /// {@macro api_client}
  StoresApiClient({
    required String baseUrl,
    required int port,
  }) {
    _channel = ClientChannel(
      baseUrl,
      port: port,
      options: const ChannelOptions(
        credentials: ChannelCredentials.insecure(),
      ),
    );
    stub = StoreServiceClient(_channel);
  }

  late final ClientChannel _channel;
  late final StoreServiceClient stub;

  Future<List<Store>> getStores() async {
    final response = await stub.getStores(StoreRequest());

    return response.stores;
  }

  Future<Store> getStoreById({
    required String id,
  }) async {
    final store = await stub.getStoreById(StoreByIdRequest(id: id));

    return store;
  }

  Future<Store> getStoreByName({
    required String name,
  }) async {
    final store = await stub.getStoreByName(StoreByNameRequest(name: name));

    return store;
  }
}
