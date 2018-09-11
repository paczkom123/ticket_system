# ticket_system

Clone repo:
  git clone https://github.com/IordanisPaschalidis/ticket_system/new/master?readme=1
Software require:
  install dep:
    brew install dep
Install dependencies using command:
  dep init
Build main.go using command:
  go build main.go
Run it
  ./main
Should create a log file with name queue_size.log which will print every second the size of the queue 
and will have a db file with name ticket.db which is the database of the allocated tickets. When the 
script will finish it will print:
  Number of request received;
  Number of requests severed; and
  Number of requested rejected;
