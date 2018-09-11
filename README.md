# ticket_system

Clone repo: <br />
** git clone https://github.com/IordanisPaschalidis/ticket_system/new/master?readme=1** <br />
Software require: <br />
  install dep: <br />
    brew install dep <br />
Install dependencies using command: <br />
  dep init <br />
Build main.go using command: <br />
  go build main.go <br />
Run it <br />
  ./main <br />
Should create a log file with name queue_size.log which will print every second the size of the queue 
and will have a db file with name ticket.db which is the database of the allocated tickets. When the 
script will finish it will print: <br />
  Number of request received <br />
  Number of requests severed <br />
  Number of requested rejected <br />
