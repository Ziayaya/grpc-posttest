# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: calculator.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10\x63\x61lculator.proto\"$\n\nAddRequest\x12\n\n\x02n1\x18\x01 \x01(\x05\x12\n\n\x02n2\x18\x02 \x01(\x05\"\x16\n\x08\x41\x64\x64Reply\x12\n\n\x02n1\x18\x01 \x01(\x05\")\n\x0fSubtractRequest\x12\n\n\x02n1\x18\x01 \x01(\x05\x12\n\n\x02n2\x18\x02 \x01(\x05\"\x1b\n\rSubtractReply\x12\n\n\x02n1\x18\x01 \x01(\x05\")\n\x0fMultiplyRequest\x12\n\n\x02n1\x18\x01 \x01(\x05\x12\n\n\x02n2\x18\x02 \x01(\x05\"\x1b\n\rMultiplyReply\x12\n\n\x02n1\x18\x01 \x01(\x05\"\'\n\rDivideRequest\x12\n\n\x02n1\x18\x01 \x01(\x05\x12\n\n\x02n2\x18\x02 \x01(\x05\"\x19\n\x0b\x44ivideReply\x12\n\n\x02n1\x18\x01 \x01(\x02\x32\xb7\x01\n\nCalculator\x12\x1f\n\x03\x41\x64\x64\x12\x0b.AddRequest\x1a\t.AddReply\"\x00\x12.\n\x08Subtract\x12\x10.SubtractRequest\x1a\x0e.SubtractReply\"\x00\x12.\n\x08Multiply\x12\x10.MultiplyRequest\x1a\x0e.MultiplyReply\"\x00\x12(\n\x06\x44ivide\x12\x0e.DivideRequest\x1a\x0c.DivideReply\"\x00\x42\x0fZ\r/calculatorpbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'calculator_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\r/calculatorpb'
  _globals['_ADDREQUEST']._serialized_start=20
  _globals['_ADDREQUEST']._serialized_end=56
  _globals['_ADDREPLY']._serialized_start=58
  _globals['_ADDREPLY']._serialized_end=80
  _globals['_SUBTRACTREQUEST']._serialized_start=82
  _globals['_SUBTRACTREQUEST']._serialized_end=123
  _globals['_SUBTRACTREPLY']._serialized_start=125
  _globals['_SUBTRACTREPLY']._serialized_end=152
  _globals['_MULTIPLYREQUEST']._serialized_start=154
  _globals['_MULTIPLYREQUEST']._serialized_end=195
  _globals['_MULTIPLYREPLY']._serialized_start=197
  _globals['_MULTIPLYREPLY']._serialized_end=224
  _globals['_DIVIDEREQUEST']._serialized_start=226
  _globals['_DIVIDEREQUEST']._serialized_end=265
  _globals['_DIVIDEREPLY']._serialized_start=267
  _globals['_DIVIDEREPLY']._serialized_end=292
  _globals['_CALCULATOR']._serialized_start=295
  _globals['_CALCULATOR']._serialized_end=478
# @@protoc_insertion_point(module_scope)
