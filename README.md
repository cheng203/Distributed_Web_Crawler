# Distributed Web Crawler
### Background

This project is a distributed web crawler collecting users' non-sensitive information from a dating website [ZhenAi](http://www.zhenai.com/) and used for customerized search over different features (age, gender, marriage, occupation, income....). It is guaranteed that no sensitive information is ever reached and crawled.

The project is a **Golang** project that supports **high-concurrency** and **distributed tasks**.

Technique Stacks

- Golang
- Goroutine & Channel
- Regular Expression
- OOD
- Docker
- ElasticSearch
- HTML/CSS/Javascript

### Project Structure

Below is a overview of project structure

<img src="https://github.com/cheng203/Distributed_Web_Crawler/blob/master/images/distributed.jpg">

​										Figure 1: Distributed Web Crawler Architecture



Below is the algorithm web crawler follows for fetching data on ZhenAi website

<img src="https://github.com/cheng203/Distributed_Web_Crawler/blob/master/images/algorithm.jpg">

​														Figure 2: Web Crawler Algorithm



In order to support high concurrency, I initially created 100 goroutine for workers and only one goroutine for scheduler. However, it terms out that this case will be very inefficient as workers need to compete for just one request at a time. Then, I modified the design and implemented queue for both requests and workers so that both scheduler and worker goroutines now have great support for high concurrency. Demo is shown below

<img src="https://github.com/cheng203/Distributed_Web_Crawler/blob/master/images/scheduler.jpg">

​						Figure 3: High Concurrency Design Details in Scheduler and Worker



<img src="https://github.com/cheng203/Distributed_Web_Crawler/blob/master/images/highConcurrency.jpg">

​										Figure 4: Architecture without distributed servers



To better use of data collected, Elastic-search is implemented through docker container and instead of having data stored in the redis or any database, data is directly stored in elastic search. A simple front end web is created to allow user search based upon their needs. Front end design is followed by MVC and is shown below.

<img src="https://github.com/cheng203/Distributed_Web_Crawler/blob/master/images/frontEnd.jpg">

​																Figure 5: Front-End MVC 



<img src="https://github.com/cheng203/Distributed_Web_Crawler/blob/master/images/display.jpeg">

​														Figure 6: Demo of Elastic-Search