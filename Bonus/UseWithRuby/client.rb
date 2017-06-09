require 'ffi'

module SomeMath
  extend FFI::Library

  ffi_lib './SomeMath.so'

  attach_function :CircleSpace, [:double], :double
end

puts SomeMath.CircleSpace(10)