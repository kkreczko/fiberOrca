cmake_minimum_required(VERSION 3.10)
project(network_analyzer C)

set(CMAKE_C_STANDARD 11)
set(CMAKE_C_STANDARD_REQUIRED ON)

if(NOT CMAKE_BUILD_TYPE)
    set(CMAKE_BUILD_TYPE Debug)
endif()

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -Wextra")
set(CMAKE_C_FLAGS_DEBUG "${CMAKE_C_FLAGS_DEBUG} -g -O0")
set(CMAKE_C_FLAGS_RELEASE "${CMAKE_C_FLAGS_RELEASE} -O3")

find_library(PCAP_LIBRARY pcap)
if(NOT PCAP_LIBRARY)
    message(FATAL_ERROR "libpcap not found")
endif()

find_path(PCAP_INCLUDE_DIR pcap.h
        PATHS /usr/include /usr/local/include)
if(NOT PCAP_INCLUDE_DIR)
    message(FATAL_ERROR "pcap.h not found")
endif()

set(SOURCES
        src/main.c
        src/capture.c
        src/packet_handler.c
        src/globals.c
        src/connect_and_send.c
        src/parse_packet.c
)

include_directories(
        ${PROJECT_SOURCE_DIR}/include
        ${PCAP_INCLUDE_DIR}
)

add_executable(${PROJECT_NAME} ${SOURCES})

target_link_libraries(${PROJECT_NAME} ${PCAP_LIBRARY})

enable_testing()
add_subdirectory(tests)