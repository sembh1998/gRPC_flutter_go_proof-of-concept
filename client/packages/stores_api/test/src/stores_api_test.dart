// ignore_for_file: prefer_const_constructors
import 'package:stores_api/stores_api.dart';
import 'package:test/test.dart';

void main() {
  group('StoresApi', () {
    test('can be instantiated', () {
      expect(Store(), isNotNull);
    });
  });
}
