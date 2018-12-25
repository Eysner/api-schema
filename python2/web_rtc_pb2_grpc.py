# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import miscellaneous_pb2 as miscellaneous__pb2
import web_rtc_pb2 as web__rtc__pb2


class WebRTCStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetCallInfo = channel.unary_unary(
        '/dialog.WebRTC/GetCallInfo',
        request_serializer=web__rtc__pb2.RequestGetCallInfo.SerializeToString,
        response_deserializer=web__rtc__pb2.ResponseGetCallInfo.FromString,
        )
    self.LoadCalls = channel.unary_unary(
        '/dialog.WebRTC/LoadCalls',
        request_serializer=web__rtc__pb2.RequestLoadCalls.SerializeToString,
        response_deserializer=web__rtc__pb2.ResponseLoadCalls.FromString,
        )
    self.DoCall = channel.unary_unary(
        '/dialog.WebRTC/DoCall',
        request_serializer=web__rtc__pb2.RequestDoCall.SerializeToString,
        response_deserializer=web__rtc__pb2.ResponseDoCall.FromString,
        )
    self.JoinCall = channel.unary_unary(
        '/dialog.WebRTC/JoinCall',
        request_serializer=web__rtc__pb2.RequestJoinCall.SerializeToString,
        response_deserializer=miscellaneous__pb2.ResponseVoid.FromString,
        )
    self.RejectCall = channel.unary_unary(
        '/dialog.WebRTC/RejectCall',
        request_serializer=web__rtc__pb2.RequestRejectCall.SerializeToString,
        response_deserializer=miscellaneous__pb2.ResponseVoid.FromString,
        )
    self.ChangeCallDisplayName = channel.unary_unary(
        '/dialog.WebRTC/ChangeCallDisplayName',
        request_serializer=web__rtc__pb2.RequestChangeCallDisplayName.SerializeToString,
        response_deserializer=miscellaneous__pb2.ResponseVoid.FromString,
        )
    self.DeleteCall = channel.unary_unary(
        '/dialog.WebRTC/DeleteCall',
        request_serializer=web__rtc__pb2.RequestDeleteCall.SerializeToString,
        response_deserializer=miscellaneous__pb2.ResponseVoid.FromString,
        )


class WebRTCServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetCallInfo(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def LoadCalls(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DoCall(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def JoinCall(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def RejectCall(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ChangeCallDisplayName(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteCall(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_WebRTCServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetCallInfo': grpc.unary_unary_rpc_method_handler(
          servicer.GetCallInfo,
          request_deserializer=web__rtc__pb2.RequestGetCallInfo.FromString,
          response_serializer=web__rtc__pb2.ResponseGetCallInfo.SerializeToString,
      ),
      'LoadCalls': grpc.unary_unary_rpc_method_handler(
          servicer.LoadCalls,
          request_deserializer=web__rtc__pb2.RequestLoadCalls.FromString,
          response_serializer=web__rtc__pb2.ResponseLoadCalls.SerializeToString,
      ),
      'DoCall': grpc.unary_unary_rpc_method_handler(
          servicer.DoCall,
          request_deserializer=web__rtc__pb2.RequestDoCall.FromString,
          response_serializer=web__rtc__pb2.ResponseDoCall.SerializeToString,
      ),
      'JoinCall': grpc.unary_unary_rpc_method_handler(
          servicer.JoinCall,
          request_deserializer=web__rtc__pb2.RequestJoinCall.FromString,
          response_serializer=miscellaneous__pb2.ResponseVoid.SerializeToString,
      ),
      'RejectCall': grpc.unary_unary_rpc_method_handler(
          servicer.RejectCall,
          request_deserializer=web__rtc__pb2.RequestRejectCall.FromString,
          response_serializer=miscellaneous__pb2.ResponseVoid.SerializeToString,
      ),
      'ChangeCallDisplayName': grpc.unary_unary_rpc_method_handler(
          servicer.ChangeCallDisplayName,
          request_deserializer=web__rtc__pb2.RequestChangeCallDisplayName.FromString,
          response_serializer=miscellaneous__pb2.ResponseVoid.SerializeToString,
      ),
      'DeleteCall': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteCall,
          request_deserializer=web__rtc__pb2.RequestDeleteCall.FromString,
          response_serializer=miscellaneous__pb2.ResponseVoid.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'dialog.WebRTC', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
