package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

import javax.annotation.PostConstruct;
import java.util.Random;
import java.util.logging.Logger;

@SpringBootApplication
public class DemoApplication {

    private static Logger logger = Logger.getLogger("Sender");

    public static void main(String[] args) {
        SpringApplication.run(DemoApplication.class, args);
    }

    @PostConstruct
    public void send() {
        for (int i = 0; i < 100000; i++) {
            int id = new Random().nextInt(100000);
            //template.convertAndSend("demo -- > %v", id);
        }
        logger.info("Sending completed.");
    }
}
