# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread

i = 0

def adder():
    global i
    for x in range(0, 1000000):     
        i += 1
        
def subtractor():
    global i
    for x in range(0,1000000):
        i -= 1
    

def main():
    adder_thr = Thread(target = adder)
    adder_thr.start()
    sub_thr = Thread(target = subtractor)
    sub_thr.start()
    adder_thr.join()
    sub_thr.join()
    print("Done: " + str(i))

main()

