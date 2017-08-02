rpc.exe -i ../schema/protocol.rpc -o ../Client/Assets/Scripts/Common/Rpc/ -g cs
rpc.exe -i ../schema/protocol.rpc -o ../server/src/logic/ -g go
rpc.exe -i ../schema/protocol.rpc -o ../server/prpc_gen/ -g cc
pause