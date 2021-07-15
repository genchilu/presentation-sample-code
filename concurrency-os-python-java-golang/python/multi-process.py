import multiprocessing
import random

def doSomething():
    while True:
        random.randrange(100000000)

p_list = []
for i in range(10):
    p = multiprocessing.Process(target=doSomething)
    p_list.append(p)

for p in p_list:
    p.start()

for p in p_list:
    p.join()