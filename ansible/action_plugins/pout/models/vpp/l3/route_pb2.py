# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: models/vpp/l3/route.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.gogo.protobuf.gogoproto import gogo_pb2 as github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='models/vpp/l3/route.proto',
  package='vpp.l3',
  syntax='proto3',
  serialized_options=_b('Z4github.com/ligato/vpp-agent/api/models/vpp/l3;vpp_l3\310\343\036\001'),
  serialized_pb=_b('\n\x19models/vpp/l3/route.proto\x12\x06vpp.l3\x1a-github.com/gogo/protobuf/gogoproto/gogo.proto\"\xf3\x01\n\x05Route\x12%\n\x04type\x18\n \x01(\x0e\x32\x17.vpp.l3.Route.RouteType\x12\x0e\n\x06vrf_id\x18\x01 \x01(\r\x12\x13\n\x0b\x64st_network\x18\x03 \x01(\t\x12\x15\n\rnext_hop_addr\x18\x04 \x01(\t\x12\x1a\n\x12outgoing_interface\x18\x05 \x01(\t\x12\x0e\n\x06weight\x18\x06 \x01(\r\x12\x12\n\npreference\x18\x07 \x01(\r\x12\x12\n\nvia_vrf_id\x18\x08 \x01(\r\"3\n\tRouteType\x12\r\n\tINTRA_VRF\x10\x00\x12\r\n\tINTER_VRF\x10\x01\x12\x08\n\x04\x44ROP\x10\x02\x42:Z4github.com/ligato/vpp-agent/api/models/vpp/l3;vpp_l3\xc8\xe3\x1e\x01\x62\x06proto3')
  ,
  dependencies=[github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2.DESCRIPTOR,])



_ROUTE_ROUTETYPE = _descriptor.EnumDescriptor(
  name='RouteType',
  full_name='vpp.l3.Route.RouteType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='INTRA_VRF', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INTER_VRF', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DROP', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=277,
  serialized_end=328,
)
_sym_db.RegisterEnumDescriptor(_ROUTE_ROUTETYPE)


_ROUTE = _descriptor.Descriptor(
  name='Route',
  full_name='vpp.l3.Route',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='vpp.l3.Route.type', index=0,
      number=10, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vrf_id', full_name='vpp.l3.Route.vrf_id', index=1,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dst_network', full_name='vpp.l3.Route.dst_network', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next_hop_addr', full_name='vpp.l3.Route.next_hop_addr', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outgoing_interface', full_name='vpp.l3.Route.outgoing_interface', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='weight', full_name='vpp.l3.Route.weight', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='preference', full_name='vpp.l3.Route.preference', index=6,
      number=7, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='via_vrf_id', full_name='vpp.l3.Route.via_vrf_id', index=7,
      number=8, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _ROUTE_ROUTETYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=85,
  serialized_end=328,
)

_ROUTE.fields_by_name['type'].enum_type = _ROUTE_ROUTETYPE
_ROUTE_ROUTETYPE.containing_type = _ROUTE
DESCRIPTOR.message_types_by_name['Route'] = _ROUTE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Route = _reflection.GeneratedProtocolMessageType('Route', (_message.Message,), dict(
  DESCRIPTOR = _ROUTE,
  __module__ = 'models.vpp.l3.route_pb2'
  # @@protoc_insertion_point(class_scope:vpp.l3.Route)
  ))
_sym_db.RegisterMessage(Route)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
