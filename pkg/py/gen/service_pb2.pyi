from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class HelloRequest(_message.Message):
    __slots__ = ("Name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    Name: str
    def __init__(self, Name: _Optional[str] = ...) -> None: ...

class HelloResponse(_message.Message):
    __slots__ = ("Greeting",)
    GREETING_FIELD_NUMBER: _ClassVar[int]
    Greeting: str
    def __init__(self, Greeting: _Optional[str] = ...) -> None: ...

class CountRequest(_message.Message):
    __slots__ = ("Name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    Name: str
    def __init__(self, Name: _Optional[str] = ...) -> None: ...

class CountResponse(_message.Message):
    __slots__ = ("Greeting", "Count")
    GREETING_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    Greeting: str
    Count: int
    def __init__(self, Greeting: _Optional[str] = ..., Count: _Optional[int] = ...) -> None: ...
