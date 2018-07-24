PROJECT(binaryenmerge C CXX)
CMAKE_MINIMUM_REQUIRED(VERSION 2.8.7)
INCLUDE(GNUInstallDirs)
#link_directories(CMAKE_SOURCE_DIR)

INCLUDE_DIRECTORIES(../binaryen/src)

# Support functionality.

FUNCTION(ADD_COMPILE_FLAG value)
  MESSAGE(STATUS "Building with ${value}")
  FOREACH(variable CMAKE_C_FLAGS CMAKE_CXX_FLAGS)
    SET(${variable} "${${variable}} ${value}" PARENT_SCOPE)
  ENDFOREACH(variable)
ENDFUNCTION()

FUNCTION(ADD_CXX_FLAG value)
  MESSAGE(STATUS "Building with ${value}")
  SET(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} ${value}" PARENT_SCOPE)
ENDFUNCTION()

FUNCTION(ADD_DEBUG_COMPILE_FLAG value)
  IF("${CMAKE_BUILD_TYPE}" MATCHES "Debug")
    MESSAGE(STATUS "Building with ${value}")
  ENDIF()
  FOREACH(variable CMAKE_C_FLAGS_DEBUG CMAKE_CXX_FLAGS_DEBUG)
    SET(${variable} "${${variable}} ${value}" PARENT_SCOPE)
  ENDFOREACH(variable)
ENDFUNCTION()

FUNCTION(ADD_NONDEBUG_COMPILE_FLAG value)
  IF(NOT "${CMAKE_BUILD_TYPE}" MATCHES "Debug")
    MESSAGE(STATUS "Building with ${value}")
  ENDIF()
  FOREACH(variable CMAKE_C_FLAGS_RELEASE CMAKE_CXX_FLAGS_RELEASE CMAKE_C_FLAGS_RELWITHDEBINFO CMAKE_CXX_FLAGS_RELWITHDEBINFO CMAKE_C_FLAGS_MINSIZEREL CMAKE_CXX_FLAGS_MINSIZEREL)
    SET(${variable} "${${variable}} ${value}" PARENT_SCOPE)
  ENDFOREACH(variable)
ENDFUNCTION()

FUNCTION(ADD_LINK_FLAG value)
  MESSAGE(STATUS "Linking with ${value}")
  FOREACH(variable CMAKE_EXE_LINKER_FLAGS)
    SET(${variable} "${${variable}} ${value}" PARENT_SCOPE)
  ENDFOREACH(variable)
ENDFUNCTION()

# Sources.
SET(binaryen_SOURCES
  bridge.cpp
)
ADD_COMPILE_FLAG("-fPIC")
ADD_LIBRARY(binaryenmerge STATIC ${binaryen_SOURCES})
#TARGET_LINK_LIBRARIES(binaryenmerge ${CMAKE_CURRENT_SOURCE_DIR}/lib/libwasm.a ${CMAKE_CURRENT_SOURCE_DIR}/lib/libasmjs.a ${CMAKE_CURRENT_SOURCE_DIR}/lib/libemscripten-optimizer.a ${CMAKE_CURRENT_SOURCE_DIR}/lib/libpasses.a ${CMAKE_CURRENT_SOURCE_DIR}/lib/libir.a ${CMAKE_CURRENT_SOURCE_DIR}/lib/libcfg.a ${CMAKE_CURRENT_SOURCE_DIR}/lib/libsupport.a ${CMAKE_CURRENT_SOURCE_DIR}/lib/libwasm.a)
