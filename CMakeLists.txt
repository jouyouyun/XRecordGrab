cmake_minimum_required(VERSION 2.8)

project("xrecord-grab")

#set(NAME "xrecord-grab")

# 查找当前目录下的所有源文件
# 并将名称保存到 DIR_SRCS 变量
AUX_SOURCE_DIRECTORY(src DIR_SRCS)

ADD_EXECUTABLE(${PROJECT_NAME} ${DIR_SRCS})

find_package(PkgConfig)
pkg_check_modules(X11 REQUIRED x11)
pkg_check_modules(XTST REQUIRED xtst)

target_link_libraries(${PROJECT_NAME}
	pthread
	${X11_LIBRARIES}
	${XTST_LIBRARIES}
)

include_directories(
	${X11_INCLUDE_DIRS}
	${XTST_INCLUDE_DIRS}
)
