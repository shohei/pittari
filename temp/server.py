import threading
import os

def start_server():
    os.system("python -m SimpleHTTPServer 8000")

def scadviewer():
    os.system("./scadview.sh hoge.scad")

if __name__=="__main__":
    t1 = threading.Thread(target=start_server)
    t2 = threading.Thread(target=scadviewer)
    t2.start()
    t1.start()


