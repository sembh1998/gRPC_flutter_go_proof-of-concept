// ignore_for_file: public_member_api_docs

import 'package:stores_api/stores_api.dart';

abstract class StoresRepositoryException implements Exception {
  /// {@macro store_repository_exception}
  const StoresRepositoryException(this.error, this.stackTrace);

  /// The underlying error that occurred.
  final Object error;

  /// The relevant stack trace.
  final StackTrace stackTrace;
}

/// {@template get_stores}
/// Thrown during the get stores if a failure occurs.
/// {@endtemplate}
class GetStoresFailure extends StoresRepositoryException {
  /// {@macro get_stores}
  const GetStoresFailure(super.error, super.stackTrace);
}

class StoresRepository {
  const StoresRepository({
    required StoresApiClient apiClient,
  }) : _apiClient = apiClient;

  final StoresApiClient _apiClient;

  Future<List<Store>> getStores() async {
    try {
      final stores = await _apiClient.getStores();

      return stores;
    } catch (error, stackTrace) {
      throw GetStoresFailure(error, stackTrace);
    }
  }

  Future<Store> getStoreById(String id) async {
    try {
      final store = await _apiClient.getStoreById(id: id);

      return store;
    } catch (error, stackTrace) {
      throw GetStoresFailure(error, stackTrace);
    }
  }

  Future<Store> getStoreByName(String name) async {
    try {
      final store = await _apiClient.getStoreByName(name: name);

      return store;
    } catch (error, stackTrace) {
      throw GetStoresFailure(error, stackTrace);
    }
  }
}
