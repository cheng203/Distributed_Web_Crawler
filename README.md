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



In order to support 