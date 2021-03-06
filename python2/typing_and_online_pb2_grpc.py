# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import miscellaneous_pb2 as miscellaneous__pb2
import typing_and_online_pb2 as typing__and__online__pb2


class TypingAndOnlineStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Typing = channel.unary_unary(
        '/dialog.TypingAndOnline/Typing',
        request_serializer=typing__and__online__pb2.RequestTyping.SerializeToString,
        response_deserializer=miscellaneous__pb2.ResponseVoid.FromString,
        )
    self.StopTyping = channel.unary_unary(
        '/dialog.TypingAndOnline/StopTyping',
        request_serializer=typing__and__online__pb2.RequestStopTyping.SerializeToString,
        response_deserializer=miscellaneous__pb2.ResponseVoid.FromString,
        )
    self.SetOnline = channel.unary_unary(
        '/dialog.TypingAndOnline/SetOnline',
        request_serializer=typing__and__online__pb2.RequestSetOnline.SerializeToString,
        response_deserializer=miscellaneous__pb2.ResponseVoid.FromString,
        )
    self.PauseNotifications = channel.unary_unary(
        '/dialog.TypingAndOnline/PauseNotifications',
        request_serializer=typing__and__online__pb2.RequestPauseNotifications.SerializeToString,
        response_deserializer=miscellaneous__pb2.ResponseVoid.FromString,
        )
    self.RestoreNotifications = channel.unary_unary(
        '/dialog.TypingAndOnline/RestoreNotifications',
        request_serializer=typing__and__online__pb2.RequestRestoreNotifications.SerializeToString,
        response_deserializer=miscellaneous__pb2.ResponseVoid.FromString,
        )


class TypingAndOnlineServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Typing(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def StopTyping(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetOnline(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PauseNotifications(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def RestoreNotifications(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_TypingAndOnlineServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Typing': grpc.unary_unary_rpc_method_handler(
          servicer.Typing,
          request_deserializer=typing__and__online__pb2.RequestTyping.FromString,
          response_serializer=miscellaneous__pb2.ResponseVoid.SerializeToString,
      ),
      'StopTyping': grpc.unary_unary_rpc_method_handler(
          servicer.StopTyping,
          request_deserializer=typing__and__online__pb2.RequestStopTyping.FromString,
          response_serializer=miscellaneous__pb2.ResponseVoid.SerializeToString,
      ),
      'SetOnline': grpc.unary_unary_rpc_method_handler(
          servicer.SetOnline,
          request_deserializer=typing__and__online__pb2.RequestSetOnline.FromString,
          response_serializer=miscellaneous__pb2.ResponseVoid.SerializeToString,
      ),
      'PauseNotifications': grpc.unary_unary_rpc_method_handler(
          servicer.PauseNotifications,
          request_deserializer=typing__and__online__pb2.RequestPauseNotifications.FromString,
          response_serializer=miscellaneous__pb2.ResponseVoid.SerializeToString,
      ),
      'RestoreNotifications': grpc.unary_unary_rpc_method_handler(
          servicer.RestoreNotifications,
          request_deserializer=typing__and__online__pb2.RequestRestoreNotifications.FromString,
          response_serializer=miscellaneous__pb2.ResponseVoid.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'dialog.TypingAndOnline', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
