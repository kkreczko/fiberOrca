cmake_minimum_required(VERSION 3.14)
project(fiberOrca)

set(CMAKE_CXX_STANDARD 23)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

include(FetchContent)
FetchContent_Declare(
    googletest
    URL https://github.com/google/googletest/archive/refs/tags/v1.14.0.zip
)
FetchContent_MakeAvailable(googletest)

enable_testing()

add_library(net net-cpp/device.cpp net-cpp/listen.cpp net-cpp/printer.cpp net-cpp/packet.cpp)
target_include_directories(net PUBLIC net-cpp)

add_library(helper net-cpp/helper.cpp)
target_include_directories(helper PUBLIC net-cpp)

add_executable(fiberorca net-cpp/main.cpp)
target_link_libraries(main PRIVATE helper net)

add_executable(net_tests tests/test_device.cpp tests/test_listen.cpp tests/test_packet.cpp tests/test_printer.cpp)
target_link_libraries(net_tests PRIVATE net helper GTest::gtest_main)

include(GoogleTest)
gtest_discover_tests(net_tests)

