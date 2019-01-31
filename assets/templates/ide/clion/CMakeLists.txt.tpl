set(TARGET_NAME {{TARGET_NAME}})
set(WIO_CMD {{WIO_PATH}})

include("${CMAKE_CURRENT_LIST_DIR}/.wio/targets/${TARGET_NAME}/CMakeLists.txt")
project(${PROJECT_NAME} C CXX ASM)
include(args.cmake)

# CMAKE VERSION
cmake_minimum_required(VERSION ${CMAKE_VERSION})

add_custom_target(
    WIO_BUILD ALL
    COMMAND ${WIO_CMD} build ${TARGET_NAME} ${WIO_BUILD_ARGS}
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
)

add_custom_target(
    WIO_RUN ALL
    COMMAND ${WIO_CMD} run ${TARGET_NAME} ${WIO_RUN_ARGS}
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
)

add_custom_target(
    WIO_CLEAN ALL
    COMMAND ${WIO_CMD} clean ${TARGET_NAME} ${WIO_CLEAN_ARGS}
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
)

add_custom_target(
    WIO_CLEAN_HARD ALL
    COMMAND ${WIO_CMD} clean ${TARGET_NAME} --hard ${WIO_CLEAN_ARGS}
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
)

add_custom_target(
    WIO_INSTALL ALL
    COMMAND ${WIO_CMD} install ${WIO_INSTALL_ARGS}
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
)

add_custom_target(
    WIO_UPDATE ALL
    COMMAND ${WIO_CMD} update --ide clion --full ${WIO_UPDATE_ARGS}
    WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
)