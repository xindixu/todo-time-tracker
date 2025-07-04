import 'package:fixnum/fixnum.dart';
import 'package:ui/rpc/google/protobuf/duration.pb.dart' as pb;

class DurationConverter {
  /// Convert Dart's [Duration] to Protobuf [Duration].
  static pb.Duration toProtobuf(Duration dartDuration) {
    return pb.Duration()
      ..seconds = Int64(dartDuration.inSeconds)
      ..nanos = (dartDuration.inMicroseconds % Duration.microsecondsPerSecond) * 1000;
  }

  /// Convert Protobuf [Duration] to Dart's [Duration].
  static Duration fromProtobuf(pb.Duration protoDuration) {
    return Duration(
      seconds: protoDuration.seconds.toInt(),
      microseconds: protoDuration.nanos ~/ 1000,
    );
  }
}