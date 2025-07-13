import 'package:ui/rpc/proto/ttt_service.pbgrpc.dart';
import 'package:grpc/grpc.dart';

class TTTService {
  // For simulator/emulator use 'localhost', for physical device use your computer's IP
  String baseUrl = '127.0.0.1'; // Changed from localhost to 127.0.0.1 for macOS
  int port = 50051;

  TTTService._internal();

  static final TTTService _instance = TTTService._internal();

  factory TTTService() => _instance;

  static TTTService get instance => _instance;

  TTTServiceClient? _tttServiceClient;
  bool _isInitialized = false;

  Future<void> init() async {
    if (!_isInitialized) {
      print('TTTService: Initializing gRPC client...');
      print('TTTService: baseUrl = "$baseUrl", port = $port');
      _createChannel();
      _isInitialized = true;
      print('TTTService: gRPC client initialized successfully');
    }
  }

  TTTServiceClient get tttServiceClient {
    if (!_isInitialized || _tttServiceClient == null) {
      throw StateError('TTTService not initialized. Call init() first.');
    }
    return _tttServiceClient!;
  }

  void _createChannel() {
    print('TTTService: Creating channel to $baseUrl:$port');
    try {
      final channel = ClientChannel(
        baseUrl,
        port: port,
        options: const ChannelOptions(credentials: ChannelCredentials.insecure()),
      );
      _tttServiceClient = TTTServiceClient(channel);
      print('TTTService: Channel created and client assigned successfully');
    } catch (e) {
      print('TTTService: Error creating channel: $e');
      rethrow;
    }
  }
}