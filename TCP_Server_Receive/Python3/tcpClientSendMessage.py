import socket

s = socket.socket()
s.connect(('localhost', 12345))
s.sendall('Send information from here'.encode())
s.close()