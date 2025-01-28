cmake_minimum_required(VERSION 3.10)
project(network_analyzer C)

set(CMAKE_C_STANDARD 11)
set(CMAKE_C_STANDARD_REQUIRED ON)

if(NOT CMAKE_BUILD_TYPE)
    set(CMAKE_BUILD_TYPE Debug)
endif()

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -Wextra -lpcap")
set(CMAKE_C_FLAGS_DEBUG "${CMAKE_C_FLAGS_DEBUG} -g -O0")
set(CMAKE_C_FLAGS_RELEASE "${CMAKE_C_FLAGS_RELEASE} -O3")

set(SOURCES
        src/main.c
        src/capture.c
        src/packet_handler.c
        src/globals.c
        src/globals.c
        include/connect_and_send.h
        src/connect_and_send.c
)

include_directories(
        ${PROJECT_SOURCE_DIR}/include
)

file(GLOB_RECURSE SOURCES
        "${PROJECT_SOURCE_DIR}/src/*.c"
)

add_executable(${PROJECT_NAME} ${SOURCES})

target_link_libraries(${PROJECT_NAME} ${PCAP_LIBRARIES})

enable_testing()
add_subdirectory(tests)