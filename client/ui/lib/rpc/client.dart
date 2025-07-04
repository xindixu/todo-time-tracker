import 'package:ui/rpc/proto/ttt_service.pbgrpc.dart';
import 'package:grpc/grpc.dart';

class TTTService {

  String baseUrl = 'http://localhost';
  int port = 50051;

  TTTService._internal();

  static final TTTService _instance = TTTService._internal();

  factory TTTService() => _instance;


  static TTTService get instance => _instance;

  late TTTServiceClient _tttServiceClient;

  Future<void> init() async {
    _createChannel();
  }

  TTTServiceClient get tttServiceClient {
    return _tttServiceClient;
  }

  void _createChannel() {
    final channel = ClientChannel(
      baseUrl,
      port: port,

      // use credentials: ChannelCredentials.insecure() if you want to connect without TLS
      //options: const ChannelOptions(credentials: ChannelCredentials.insecure()),

      // use this if you are connecting with TLS
      options: const ChannelOptions(credentials: ChannelCredentials.insecure()),
    );
    _tttServiceClient = TTTServiceClient(channel);
  }
}