INCLUDE_DIRECTORIES(${CMAKE_CURRENT_BINARY_DIR})
INCLUDE_DIRECTORIES(${CMAKE_CURRENT_SOURCE_DIR})
INCLUDE_DIRECTORIES(${CMAKE_CURRENT_SOURCE_DIR}/../common)
INCLUDE_DIRECTORIES(${CMAKE_CURRENT_SOURCE_DIR}/../prpc_gen)

SET(INC
	logic.h

	unit_table.h
	skill_table.h
	
	skill.h
	session.h
	entity.h
	player.h
	battle.h
)
	
SET(SRC
	logic.cc
	
	unit_table.cc
	skill_table.cc
	
	skill.cc
	session.cc
	entity.cc
	player.cc
	battle.cc
)

SOURCE_GROUP(Header FILES ${INC})
SOURCE_GROUP(Source FILES ${SRC})
 
ADD_EXECUTABLE(logic ${YACC} ${INC} ${SRC})
TARGET_LINK_LIBRARIES(logic common)
TARGET_LINK_LIBRARIES(logic prpc_gen)
TARGET_LINK_LIBRARIES(logic ${Boost_LIBRARIES})

INSTALL(TARGETS logic DESTINATION bin)