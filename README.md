# Go-scrapper

Funcionality:
    It is a console script program built in Go. The scrapper funcionality consists in on reading the web page https://rock.com.ar/. In the web page you can see the alphabet. the scrapper make all the alphabet requests collecting information about the artists. After that, throught all artists we collect information about their song. Finally the information is persisted into a database

How to run:
    1) by console, go to docker-golang, run: docker-compose --build
        - Set on the web navigator the http://localhost:5003/ url. This is the access to the database. The credentials are user: root and pass: root
    2) Having Go cli installed on your machine you set the location on /scripts. You will see many different folders in wich you have a main.go file inside of each ones. Those main.go files are different solutions that I have developed in order to fullfit with the funcionality of the scrapper. 
        -one_main_rutine: This is a solution with one single go rutine. (inside the folder run: run *.go)
        -wait_group_mutex: This is a program solution built by using multiple go runtines (inside the folder run: go run -race *.go)
        -channel_send_recieve: This is a program solution that offer the use of two channels (inside the folder run: go run -race *.go)

After run the program:
    - Set on the web navigator the http://localhost:5003/ url. You will see one databse called "scrapper" with two tables "artirts" and "songs". You we see all the data of artits and their song stored in these tables. 