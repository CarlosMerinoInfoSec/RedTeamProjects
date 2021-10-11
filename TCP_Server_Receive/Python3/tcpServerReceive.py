import socket

s = socket.socket()
s.bind(('0.0.0.0', 12345))

s.listen(5)
while True:
	c, addr = s.accept()
	print ('Got connection from', addr)
	print (c.recv(1024))
	c.close()