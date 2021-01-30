# LINE Assignment
*지원자 박현호*

1. Project 설명
    <br> 요구 사항에 맞도록 어플리케이션 구성
    <br> APP -> Go Routine -> Host별 log file 생성

2. 테스트 방법
   1. Ping Job 생성
      - curl -i -X POST -F "server=google.com" -F "count=100" localhost:1323/ping
      - curl -i -X POST -F "server=naver.com" -F "count=200" localhost:1323/ping

   2. Ping cat log
      - curl -i -X GET localhost:1323/google.com
      - curl -i -X GET localhost:1323/naver.com

   3. Ping tail -f log
      - curl -i -X GET localhost:1323/google.com\?wait\=true
      - curl -i -X GET localhost:1323/naver.com\?wait\=true

   4. Ping Job List
      - curl -i -X GET localhost:1323

   5. Delete Ping Job
      - curl -i -X DELETE localhost:1323/google.com
      - curl -i -X DELETE localhost:1323/naver.com