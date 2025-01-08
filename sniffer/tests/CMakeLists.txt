file(GLOB_RECURSE TEST_SOURCES
        "*.c"
)

add_executable(unit_tests ${TEST_SOURCES})
target_link_libraries(unit_tests ${PCAP_LIBRARIES})

add_test(NAME unit_tests COMMAND unit_tests)