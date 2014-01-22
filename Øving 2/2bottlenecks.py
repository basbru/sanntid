# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread
from threading import Lock

i = 0
lock = Lock()

def adder():
    global i
    for x in range(0, 1000000):
	lock.acquire(True)     
        i += 1
	lock.release()
        
def subtractor():
    global i
    for x in range(0,1000000):
	lock.acquire(True)
        i -= 1
	lock.release()
    

def main():
	lock = Lock()
	adder_thr = Thread(target = adder)
	adder_thr.start()
	sub_thr = Thread(target = subtractor)
	sub_thr.start()
	adder_thr.join()
	sub_thr.join()
	print("Done: " + str(i))

main()
