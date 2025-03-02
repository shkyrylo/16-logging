FROM docker.elastic.co/logstash/logstash:7.10.2

RUN bin/logstash-plugin install logstash-output-gelf

COPY logstash.conf /usr/share/logstash/pipeline/logstash.conf

CMD ["logstash"]