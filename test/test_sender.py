import socket

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('127.0.0.2', 1200))
s.send('Hello'.encode())
msg = s.recv(1024)
print(msg)
s.send(msg)
msg = s.recv(1024)
print(msg)
s.send(msg)
msg = s.recv(1024)
print(msg)
s.send(msg)
msg = s.recv(1024)
print(msg)
s.send(msg)
msg = s.recv(1024)
print(msg)
s.send(msg)
msg = s.recv(1024)
print(msg)
s.send(msg)
msg = s.recv(1024)
print(msg)
s.send(msg)
msg = s.recv(1024)
print(msg)
s.send(msg)
msg = s.recv(1024)
print(msg)
s.send(msg)
msg = s.recv(1024)
print(msg)
s.send(msg)

# shut_down = False

# while not shut_down:
#     try:
#         msg = s.recv(1024)
#         if not len(msg):
#             print(msg)
#             s.send(msg)
#     except KeyboardInterrupt:
#         shut_down = True
#         continue
#     except Exception as e:
#         print(e)
#         break