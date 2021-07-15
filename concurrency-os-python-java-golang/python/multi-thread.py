import threading
import random

def doSomething():
    while True:
        random.randrange(100000000)

t_list = []
for i in range(10):
    t = threading.Thread(target=doSomething)
    t_list.append(t)

for t in t_list:
    t.start()

for t in t_list:
    t.join()